package ctp

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
	"unsafe"

	"github.com/golang/protobuf/proto"
	"github.com/mineralres/goshare/pkg/util"
	"github.com/mineralres/goshare/pkg/pb"
	"github.com/mineralres/goshare/pkg/service/tahub"
)

// Trader api
type Trader struct {
	inBuffer  []byte
	outBuffer []byte
	session
	sync.RWMutex
}

func (trader *Trader) callAPI(callType pb.BindingMessageType, d proto.Message) (r1, r2 uintptr, lastErr error) {
	// 因为会修改inBuffer，所以需要锁
	trader.Lock()
	defer trader.Unlock()
	trader.inType = uint64(callType)
	if d != nil {
		out, _ := proto.Marshal(d)
		copy(trader.inBuffer[:], out[:])
		trader.inDataLen = uint64(len(out))
	}
	// log.Println("ctp call ", callType)
	return ctpTradeCall.Call(uintptr(unsafe.Pointer(&trader.session)))
}

func onTraderPopupMessage(s *session) {
	if s.gospi < 0 || s.gospi > 990 {
		panic("onTraderPopupMessage s.gospi < 0 or > 990 ")
	}
	data := (*[bufferSize]byte)(s.outData)[:s.outDataLen]
	// s.gospi 是index
	spi := HandlerList[s.gospi]
	// log.Println("pb.BindingMessageType(s.outType)", pb.BindingMessageType(s.outType))
	switch pb.BindingMessageType(s.outType) {
	case pb.BindingMessageType_CTP_ON_FRONT_CONNECTED:
		spi.OnFrontConnected()
	case pb.BindingMessageType_CTP_RSP_USER_LOGIN:
		var rsp pb.RspTradingAccountLogin
		if err := proto.Unmarshal(data, &rsp); err == nil {
			spi.OnRspUserLogin(&rsp)
		}
	case pb.BindingMessageType_CTP_ON_RTN_ORDER:
		var s pb.CTPOrderField
		if err := proto.Unmarshal(data, &s); err == nil {
			spi.OnRtnOrder(fromCTPOrderField(&s))
		}
	case pb.BindingMessageType_CTP_ON_RTN_TRADE:
		var s pb.CTPTradeField
		if err := proto.Unmarshal(data, &s); err == nil {
			spi.OnRtnTrade(fromCTPTradeField(&s))
		}
	case pb.BindingMessageType_CTP_ON_RSP_ORDER_INSERT:
		var rsp pb.RspOrderInsert
		if err := proto.Unmarshal(data, &rsp); err == nil {
			spi.OnRspOrderInsert(&rsp)
		}
	default:
		spi.OnOther(pb.BindingMessageType(s.outType), data)
	}
}

// MakeTrader 创建接口, 返回接口ID
func MakeTrader(r *pb.TradingRoute, spi tahub.Handler) *Trader {
	if spi == nil {
		panic("")
	}
	trader := &Trader{}
	trader.goapi = uintptr(unsafe.Pointer(trader))
	trader.gospi = insertHandler(spi)
	// log.Println("trader.gospi", trader.gospi, unsafe.Sizeof(trader.session))

	trader.inBuffer = make([]byte, bufferSize)
	trader.outBuffer = make([]byte, bufferSize)
	trader.inData = unsafe.Pointer(&trader.inBuffer[0])
	trader.outData = unsafe.Pointer(&trader.outBuffer[0])

	trader.callAPI(pb.BindingMessageType_CTP_CREATE_API, r)
	return trader
}

// Destory 销毁接口
func (trader *Trader) Destory() {
	trader.callAPI(pb.BindingMessageType_CTP_DELETE_API, nil)
}

// Login 登陆
func (trader *Trader) Login(ea *pb.TradingAccount) error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_USER_LOGIN, ea)
	return nil
}

func parseAPIError(r1, r2 uintptr) error {
	switch int32(r1) {
	case 0:
		return nil
	case -1:
		return errors.New("网络连接失败")
	case -2:
		return errors.New("未处理请求超过许可数")
	case -3:
		return errors.New("每秒发送请求超过许可数")
	}
	if r2 == 0 {
		return nil
	}
	return nil
}

// InsertOrder 发单
func (trader *Trader) InsertOrder(order *pb.Order) error {
	r1, r2, _ := trader.callAPI(pb.BindingMessageType_CTP_REQ_INSERT_ORDER, order)
	return parseAPIError(r1, r2)
}

// CancelOrder 撤单
func (trader *Trader) CancelOrder(req *pb.CancelOrderRequest) error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_CANCEL_ORDER, req)
	return nil
}

// QuerySummary 查询资金
func (trader *Trader) QuerySummary() error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_QRY_TRADING_ACCOUNT, nil)
	return nil
}

// QueryPositionDetail 查询持仓明细
func (trader *Trader) QueryPositionDetail() error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_QRY_POSITION_DETAIL, nil)
	return nil
}

// QueryCommissionRate 查询手续费率
func (trader *Trader) QueryCommissionRate(s *pb.Symbol) error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_QRY_COMMISSION_RATE, s)
	return nil
}

// QueryMarginRate 查询保证金率
func (trader *Trader) QueryMarginRate(s *pb.Symbol) error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_QRY_MARGIN_RATE, s)
	return nil
}

// Connect 连接
func (trader *Trader) Connect(req *pb.CtpReqConnect) error {
	var err error
	req.Directory, err = os.Getwd()
	trader.callAPI(pb.BindingMessageType_CTP_REQ_CONNECT, req)
	return err
}

// Authencate 认证
func (trader *Trader) Authencate(req *pb.CtpReqAuthencate) error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_AUTHENTICATE, req)
	return nil
}

// LoginOnly 登陆
func (trader *Trader) LoginOnly(req *pb.CtpReqUserLogin) error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_USER_LOGIN, req)
	return nil
}

// ConfirmSettlementInfo 确认结算单
func (trader *Trader) ConfirmSettlementInfo(req *pb.CtpReqSettlementInfoConfirm) error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_SETTLEMENT_INFO_CONFIRM, req)
	return nil
}

// QueryInvestor 查询投资者
func (trader *Trader) QueryInvestor(req *pb.CtpReqQryInvestor) error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_QRY_INVESTOR, req)
	return nil
}

// QueryTransferBank 查询银行
func (trader *Trader) QueryTransferBank(req *pb.CtpReqQryTransferBank) error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_QRY_TRANSFER_BANK, req)
	return nil
}

// Transfer 银行转期货
func (trader *Trader) Transfer(req *pb.CtpReqTransfer) error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_TRANSFER, req)
	return nil
}

// QueryAccountRegister 查询注册资金账号
func (trader *Trader) QueryAccountRegister(req *pb.CtpReqQryAccountRegister) error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_QRY_ACCOUNT_REGISTER, req)
	return nil
}

// Init 初始化
func (trader *Trader) Init() error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_CALL_INIT, nil)
	return nil
}

// Disconnect 断开连接
func (trader *Trader) Disconnect() error {
	trader.callAPI(pb.BindingMessageType_CTP_REQ_DISCONNECT, nil)
	return nil
}

func fromCTPOrderRef(s string) int32 {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return int32(i)
}

func fromCTPDirection(d int32) pb.OrderDirection {
	if d == '0' {
		return pb.OrderDirection_OD_LONG
	}
	return pb.OrderDirection_OD_SHORT
}

func fromCTPExchange(ex string) pb.ExchangeType {
	switch ex {
	case "SHFE":
		return pb.ExchangeType_SHFE
	case "DCE":
		return pb.ExchangeType_DCE
	case "CZCE":
		return pb.ExchangeType_CZCE
	case "CFFEX":
		return pb.ExchangeType_CFFEX
	case "INE":
		return pb.ExchangeType_INE
	}
	return pb.ExchangeType_INVALIDEX
}

func fromCTPPriceType(s int32) pb.OrderPriceType {
	switch s {
	case '1':
		return pb.OrderPriceType_OPT_MARKET_PRICE
	case '2':
		return pb.OrderPriceType_OPT_LIMIT_PRICE
	case '3':
		return pb.OrderPriceType_OPT_BEST_PRICE
	case '4':
		return pb.OrderPriceType_OPT_LAST_PRICE
	}
	return pb.OrderPriceType_OPT_LIMIT_PRICE
	///任意价
	// #define THOST_FTDC_OPT_AnyPrice '1'
	// ///限价
	// #define THOST_FTDC_OPT_LimitPrice '2'
	// ///最优价
	// #define THOST_FTDC_OPT_BestPrice '3'
	// ///最新价
	// #define THOST_FTDC_OPT_LastPrice '4'
}

func fromCTPOrderStatus(s int32) pb.OrderStatus {
	switch s {
	case '0':
		return pb.OrderStatus_OS_DONE
	case '1':
		return pb.OrderStatus_OS_PENDING_WITH_PARTIAL_DONE
	case '2', '4', '5':
		return pb.OrderStatus_OS_CANCELED
	case '3':
		return pb.OrderStatus_OS_PENDING
	}
	return pb.OrderStatus_OS_UNKOWN
	/*
		///全部成交
		#define THOST_FTDC_OST_AllTraded '0'
		///部分成交还在队列中
		#define THOST_FTDC_OST_PartTradedQueueing '1'
		///部分成交不在队列中
		#define THOST_FTDC_OST_PartTradedNotQueueing '2'
		///未成交还在队列中
		#define THOST_FTDC_OST_NoTradeQueueing '3'
		///未成交不在队列中
		#define THOST_FTDC_OST_NoTradeNotQueueing '4'
		///撤单
		#define THOST_FTDC_OST_Canceled '5'
		///未知
		#define THOST_FTDC_OST_Unknown 'a'
		///尚未触发
		#define THOST_FTDC_OST_NotTouched 'b'
		///已触发
		#define THOST_FTDC_OST_Touched 'c'
	*/

}

func fromCTPOrderField(s *pb.CTPOrderField) *pb.Order {
	var ret pb.Order
	ret.Id = &pb.OrderID{FrontId: s.FrontId, SessionId: s.SessionId, OrderRef: fromCTPOrderRef(s.OrderRef)}
	ret.ExchangeOrderId = s.OrderSysId
	ret.Direction = fromCTPDirection(s.Direction)
	ret.Symbol = &pb.Symbol{Exchange: fromCTPExchange(s.ExchangeId), Code: s.InstrumentId}
	ret.LimitPrice = s.LimitPrice
	ret.Volume = s.VolumeTotalOriginal
	ret.VolumeTraded = s.VolumeTraded
	ret.PriceType = fromCTPPriceType(s.OrderPriceType)
	ret.OffsetFlag = fromCTPOffsetFlag(s.CombOffsetFlag)
	ret.Status = fromCTPOrderStatus(s.OrderStatus)
	if ret.Status == pb.OrderStatus_OS_CANCELED {
		ret.VolumeCanceled = ret.Volume - ret.VolumeTraded
	}
	ret.Comment = util.StringFromGBK2(s.StatusMsg)
	ret.TradingDay = s.TradingDay
	ret.UserProductInfo = s.UserProductInfo
	return &ret
}

func fromCTPTradeField(s *pb.CTPTradeField) *pb.TradeReport {
	var ret pb.TradeReport
	ret.ExchangeOrderId = s.OrderSysId
	ret.Price = s.Price
	ret.Symbol = &pb.Symbol{Code: s.InstrumentId}
	ret.Symbol.Exchange = fromCTPExchange(s.ExchangeId)
	ret.TradeId = s.TradeId
	ret.TradedTradingDay = s.TradingDay
	t, err := time.Parse("20060102 15:04:05", fmt.Sprintf("%s %s", s.TradeDate, s.TradeTime))
	if err == nil {
		ret.TradedTime = t.Unix() - 28800 // 北京时间
	}
	ret.Volume = s.Volume
	if s.Direction == '0' {
		ret.Direction = pb.OrderDirection_OD_LONG
	} else {
		ret.Direction = pb.OrderDirection_OD_SHORT
	}
	ret.OffsetFlag = fromCTPOffsetFlag(s.OffsetFlag)
	ret.TradeType = pb.TradeType_TT_NORMAL
	return &ret
}

func fromCTPOffsetFlag(s int32) pb.OffsetFlag {
	switch s {
	case '0':
		return pb.OffsetFlag_OF_OPEN
	case '1':
		return pb.OffsetFlag_OF_CLOSE
	case '2':
		return pb.OffsetFlag_OF_FORCE_CLOSE
	case '3':
		return pb.OffsetFlag_OF_CLOSE_TODAY
	case '4':
		return pb.OffsetFlag_OF_CLOSE_YESTERDAY
	}
	return pb.OffsetFlag_OF_OPEN
}

func fromCTPHedgeFlag(s int32) pb.OrderHedgeType {
	return pb.OrderHedgeType_OHT_HEDGE
}

// ReqQryInstrument 查询合约
func (trader *Trader) ReqQryInstrument(s *pb.Symbol) error {
	r1, r2, _ := trader.callAPI(pb.BindingMessageType_CTP_REQ_QRY_INSTRUMENT, s)
	return parseAPIError(r1, r2)
}

func fromCTPDate(str string) int32 {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return int32(i)
}

// FromCTPInstrumentField FromCTPInstrumentField
func FromCTPInstrumentField(s *pb.CTPInstrumentField) *pb.TradingInstrument {
	d := &pb.TradingInstrument{}
	d.InstrumentInfo = &pb.InstrumentInfo{}
	d.ProductInfo = &pb.ProductInfo{ProductId: &pb.ProductID{}}
	d.InstrumentInfo.SymbolName = util.StringFromGBK2(s.InstrumentName)
	d.ProductInfo.ProductId.Code = s.ProductId
	d.Symbol = &pb.Symbol{Exchange: fromCTPExchange(s.ExchangeId), Code: s.InstrumentId}
	d.ProductInfo.ProductId.Exchange = d.Symbol.Exchange
	d.ProductInfo.PriceTick = s.PriceTick
	d.ProductInfo.VolumeMultiple = s.VolumeMultiple
	if d.ProductInfo.VolumeMultiple == 0 {
		d.ProductInfo.VolumeMultiple = 1
	}
	d.InstrumentInfo.PreClosePrice = 0
	d.ProductInfo.Type = pb.ProductType_PT_FUTURE
	d.InstrumentInfo.IsCloseTodayAllowed = 1
	d.InstrumentInfo.MaxLimitOrderVolume = s.MaxLimitOrderVolume
	d.InstrumentInfo.MinLimitOrderVolume = s.MinLimitOrderVolume
	d.InstrumentInfo.MaxMarketOrderVolume = s.MaxMarketOrderVolume
	d.InstrumentInfo.MinMarketOrderVolume = s.MinMarketOrderVolume
	d.InstrumentInfo.IsTrading = s.IsTrading
	d.InstrumentInfo.CreateDate = fromCTPDate(s.CreateDate)
	d.InstrumentInfo.OpenDate = fromCTPDate(s.OpenDate)
	d.InstrumentInfo.ExpireDate = fromCTPDate(s.ExpireDate)
	d.InstrumentInfo.StartDeliverDate = fromCTPDate(s.StartDelivDate)
	d.InstrumentInfo.EndDeliverDate = fromCTPDate(s.EndDelivDate)
	if d.InstrumentInfo.CreateDate == 0 {
		d.InstrumentInfo.CreateDate = 19700101
	}
	if d.InstrumentInfo.OpenDate == 0 {
		d.InstrumentInfo.OpenDate = 19700101
	}
	if d.InstrumentInfo.StartDeliverDate == 0 {
		d.InstrumentInfo.StartDeliverDate = 19700101
	}
	if d.InstrumentInfo.EndDeliverDate == 0 {
		d.InstrumentInfo.EndDeliverDate = 19700101
	}
	return d
}
