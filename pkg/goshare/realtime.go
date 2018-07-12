package goshare

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"

	"github.com/mineralres/goshare/pkg/base"

	"github.com/mineralres/goshare/pkg/pb"
)

// GetLastTick 取最新行情
func (p *Service) GetLastTick(symbol *pb.Symbol) (*pb.MarketDataSnapshot, error) {
	if symbol.Exchange == pb.ExchangeType_SSE || symbol.Exchange == pb.ExchangeType_SZE {
		return getStockLastTick(symbol)
	}

	if symbol.Exchange == pb.ExchangeType_INDEX {
		return getIndexLastTick(symbol)
	}

	if symbol.Exchange == pb.ExchangeType_OPTION_SSE {
		return getOptionSSETick(symbol)
	}

	return nil, base.ErrUnsupported
}

// getStockLastTick 取股票最新报价
func getStockLastTick(symbol *pb.Symbol) (*pb.MarketDataSnapshot, error) {
	ret := &pb.MarketDataSnapshot{}
	exstr := "sh"
	if symbol.Exchange == pb.ExchangeType_SZE {
		exstr = "sz"
	}
	tickArr := getRawTickString(exstr, symbol.Code)
	if tickArr == nil || len(tickArr) < 38 {
		return nil, errors.New("ErrGetStockTick")
	}
	if tickArr != nil && len(tickArr) >= 38 {
		timeStr := tickArr[30]
		tx, err := time.Parse("20060102150405", timeStr)
		if err == nil {
			ret.Time = tx.Unix()
		}
		ret.Symbol = *symbol
		ret.Price = base.ParseFloat(tickArr[3])
		ret.Close = ret.Price
		ret.PreClose = base.ParseFloat(tickArr[4])
		ret.Open = base.ParseFloat(tickArr[5])
		ret.High = base.ParseFloat(tickArr[33])
		ret.Low = base.ParseFloat(tickArr[34])
		ret.Volume = (base.ParseFloat(tickArr[6]))
		ret.Amount = float64(base.ParseInt(tickArr[37]) * 10000)
		ret.UpperLimitPrice = base.ParseFloat(tickArr[47])
		ret.LowerLimitPrice = base.ParseFloat(tickArr[48])
		var ob5 pb.OrderBook
		ob5.BidVolume = base.ParseFloat(tickArr[18])
		ob5.Bid = base.ParseFloat(tickArr[17])
		ob5.AskVolume = base.ParseFloat(tickArr[28])
		ob5.Ask = base.ParseFloat(tickArr[27])
		var ob4 pb.OrderBook
		ob4.BidVolume = base.ParseFloat(tickArr[16])
		ob4.Bid = base.ParseFloat(tickArr[15])
		ob4.AskVolume = base.ParseFloat(tickArr[26])
		ob4.Ask = base.ParseFloat(tickArr[25])
		var ob3 pb.OrderBook
		ob3.BidVolume = base.ParseFloat(tickArr[14])
		ob3.Bid = base.ParseFloat(tickArr[13])
		ob3.AskVolume = base.ParseFloat(tickArr[24])
		ob3.Ask = base.ParseFloat(tickArr[23])
		var ob2 pb.OrderBook
		ob2.BidVolume = base.ParseFloat(tickArr[12])
		ob2.Bid = base.ParseFloat(tickArr[11])
		ob2.AskVolume = base.ParseFloat(tickArr[22])
		ob2.Ask = base.ParseFloat(tickArr[21])
		var ob1 pb.OrderBook
		ob1.BidVolume = base.ParseFloat(tickArr[10])
		ob1.Bid = base.ParseFloat(tickArr[9])
		ob1.AskVolume = base.ParseFloat(tickArr[20])
		ob1.Ask = base.ParseFloat(tickArr[19])
		ret.OrderBookList = []pb.OrderBook{ob1, ob2, ob3, ob4, ob5}
	}
	return ret, nil
}

func getRawTickString(exstr string, symbol string) []string {
	resp, err := http.Get("http://web.sqt.gtimg.cn/q=" + exstr + symbol)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			tickArr := strings.Split(string(body), "~")
			if len(tickArr) > 1 {
				data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(tickArr[1])), simplifiedchinese.GBK.NewDecoder()))
				if err == nil {
					tickArr[1] = string(data)
				}
			}
			return tickArr
		}
	}
	return nil
}

func getIndexLastTick(symbol *pb.Symbol) (*pb.MarketDataSnapshot, error) {
	ret := &pb.MarketDataSnapshot{}
	resp, err := http.Get("http://hq.sinajs.cn/list=" + symbol.Code)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			tickArr := strings.Split(string(body), ",")
			//sym := strings.Split(string(tickArr[0]), "=")
			ret.Symbol = *symbol
			ret.Price = base.ParseFloat(tickArr[1])
			ret.Close = ret.Price
			return ret, nil
		}
	}
	return nil, errors.New("ErrGetIndex")
}

// parse sina tick string
func parseSinaOptionTick(body string) (*pb.MarketDataSnapshot, string, error) {
	ret := &pb.MarketDataSnapshot{}
	tickArr := strings.Split(string(body), ",")
	if len(tickArr) >= 42 {
		var ss string
		tickSym2 := strings.Split(strings.Split(tickArr[0], "=")[0], "_")
		ss = tickSym2[4]
		symbol := pb.Symbol{Exchange: pb.ExchangeType_OPTION_SSE, Code: ss}
		ret.Symbol = symbol
		ret.Price = base.ParseFloat(tickArr[2])
		ret.Close = ret.Price
		ret.PreClose = base.ParseFloat(tickArr[8])
		ret.Open = base.ParseFloat(tickArr[9])
		ret.High = base.ParseFloat(tickArr[39])
		ret.Low = base.ParseFloat(tickArr[40])
		ret.Volume = (base.ParseFloat(tickArr[41]))
		ret.Amount = float64(base.ParseInt(tickArr[42]))
		ret.UpperLimitPrice = base.ParseFloat(tickArr[10])
		ret.LowerLimitPrice = base.ParseFloat(tickArr[11])
		var ob5 pb.OrderBook
		ob5.BidVolume = base.ParseFloat(tickArr[12])
		ob5.Bid = base.ParseFloat(tickArr[13])
		ob5.AskVolume = base.ParseFloat(tickArr[30])
		ob5.Ask = base.ParseFloat(tickArr[31])
		var ob4 pb.OrderBook
		ob4.BidVolume = base.ParseFloat(tickArr[14])
		ob4.Bid = base.ParseFloat(tickArr[15])
		ob4.AskVolume = base.ParseFloat(tickArr[28])
		ob4.Ask = base.ParseFloat(tickArr[29])
		var ob3 pb.OrderBook
		ob3.BidVolume = base.ParseFloat(tickArr[16])
		ob3.Bid = base.ParseFloat(tickArr[17])
		ob3.AskVolume = base.ParseFloat(tickArr[26])
		ob3.Ask = base.ParseFloat(tickArr[27])
		var ob2 pb.OrderBook
		ob2.BidVolume = base.ParseFloat(tickArr[18])
		ob2.Bid = base.ParseFloat(tickArr[19])
		ob2.AskVolume = base.ParseFloat(tickArr[24])
		ob2.Ask = base.ParseFloat(tickArr[25])
		var ob1 pb.OrderBook
		ob1.BidVolume = base.ParseFloat(tickArr[20])
		ob1.Bid = base.ParseFloat(tickArr[21])
		ob1.AskVolume = base.ParseFloat(tickArr[22])
		ob1.Ask = base.ParseFloat(tickArr[23])
		return ret, tickArr[37], nil
	}
	return nil, "", errors.New("error")
}

// 根据合约获取单个期权合约的tick数据
func getOptionSSETick(symbol *pb.Symbol) (*pb.MarketDataSnapshot, error) {
	//ret := &pb.MarketDataSnapshot{}
	resp, err := http.Get("http://hq.sinajs.cn/list=" + symbol.Code)
	if err == nil {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		ret, _, err1 := parseSinaOptionTick(string(body))
		if err1 == nil {
			return ret, nil
		}
	}
	return nil, errors.New("ErrGetIndex")
}

// 批量获取50etf tick数据
func getOptionSSETickT(symbol string) ([]pb.MarketDataSnapshot, []string, error) {
	rets := []pb.MarketDataSnapshot{}
	retsName := []string{}
	syms := GetSina50EtfSym(symbol)
	all := "http://hq.sinajs.cn/list="
	for _, value := range syms {
		//log.Printf(" sina 期权合约代码为: %s\n", value)
		all = all + value + ","
	}
	resp, err := http.Get(all)
	if err == nil {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		tickArr1 := strings.Split(string(body), ";")
		for _, v := range tickArr1 {
			ret, name, err1 := parseSinaOptionTick(string(v))
			if err1 == nil {
				rets = append(rets, *ret)
				retsName = append(retsName, name)
			}
		}
		return rets, retsName, nil
	}
	return nil, nil, errors.New("ErrGetIndex")
}

// GetOptionSinaTick 根据交割月获取t型报价表数据
/* date 如1808 为8月到期的
 */
func (p *Service) GetOptionSinaTick(date string) ([]pb.MarketDataSnapshot, error) {
	rets := []pb.MarketDataSnapshot{}

	all := "OP_DOWN_510050" + date
	allTick, _, _ := getOptionSSETickT(all)
	rets = append(rets, allTick...)

	all = "OP_UP_510050" + date
	allTick, _, _ = getOptionSSETickT(all)
	rets = append(rets, allTick...)

	return rets, errors.New("ErrGetIndex")
}

// GetOptionSinaTick 根据交割月获取t型报价表数据
/* date 如1808 为8月到期的
 */
func (p *Service) GetOptionSinaTickMarket(date string) ([]pb.OptionTMarket, error) {
	rets := []pb.OptionTMarket{}

	all := "OP_DOWN_510050" + date
	allTick, allName, _ := getOptionSSETickT(all)

	all = "OP_UP_510050" + date
	allTick1, _, _ := getOptionSSETickT(all)

	for kk, val := range allName {
		msg := pb.OptionTMarket{}
		msg.Name = val
		msg.CallTk = allTick1[kk]
		msg.PutTk = allTick[kk]
		rets = append(rets, msg)
		//log.Printf("执行价为%s,call 为%s,put 为%s", val, msg.CallTk.Symbol.Code, msg.PutTk.Symbol.Code)
	}
	return rets, errors.New("ErrGetIndex")
}

// GetSina50EtfSym 获取50ETF期权合约列表，sina代码
func GetSina50EtfSym(sym string) (slice []string) {
	//说明：
	//OP_DOWN_5100501807:OP 期权、DOWN 看跌、UP 看涨、510050 50etf标的代码、1807 到期月份
	//根据到期月的期权从接口获取t型的合约表： CON_OP_10001394
	// 参数解释：CON_OP_ 为固定title，10001394这个是交易所的合约代码，在任何一个行情软件都可以查到，也可以通过GetSina50EtfSym接口获取
	// GetLastTick 根据CON_OP_10001394可以获取最新的报价
	// GetKData 根据CON_OP_10001394可以获取日k线
	resp, err := http.Get("http://hq.sinajs.cn/list=" + sym)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		tickArr := strings.Split(string(body), ",")
		//log.Println(string(body), tickArr)
		i := len(tickArr)
		if err == nil {
			slice = make([]string, i-2)
			for j := 1; j < i-1; j++ {
				slice[j-1] = tickArr[j]
			}
			return slice
		}
	}
	return nil
}

// GetMainFutureLastTick 取主力合约
func (p *Service) GetMainFutureLastTick(et pb.ExchangeType) ([]pb.MarketDataSnapshot, error) {
	var ret []pb.MarketDataSnapshot
	var etStr string
	switch et {
	case pb.ExchangeType_SHFE:
		etStr = "SHFE"
	case pb.ExchangeType_CZCE:
		etStr = "CZCE"
	case pb.ExchangeType_DCE:
		etStr = "DCE"
	case pb.ExchangeType_CFFEX:
		etStr = "_168"
	default:
		return ret, fmt.Errorf("error ExchangeType %s", et)
	}

	address := fmt.Sprintf("http://nufm.dfcfw.com/EM_Finance2014NumericApplication/JS.aspx?type=CT&cmd=C.%s", etStr) + "&sty=FCFL4O&sortType=(ChangePercent)&sortRule=-1&page=1&pageSize=200&js={rank:[(x)],pages:(pc),total:(tot)}&token=7bc05d0d4c3c22ef9fca8c2a912d779c&jsName=quote_123&_g=0.628606915911589&_=1521620666159"

	resp, err := http.Get(address)
	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		if err == nil {
			tickArr := strings.Split(string(body), "\"")
			// fmt.Println(string(body))
			i := 0
			for i < len(tickArr) {

				mktStrArr := strings.Split(string(tickArr[i]), ",")
				i++
				if len(mktStrArr) < 15 {
					continue
				}
				if len(mktStrArr[1]) > 3 {
					continue
				}

				mkt := pb.MarketDataSnapshot{}
				mkt.Symbol = pb.Symbol{Exchange: et, Code: mktStrArr[1]}
				mkt.Open = base.ParseFloat(mktStrArr[11])
				mkt.High = base.ParseFloat(mktStrArr[13])
				mkt.Low = base.ParseFloat(mktStrArr[14])
				mkt.Price = base.ParseFloat(mktStrArr[3])
				mkt.Close = mkt.Price
				mkt.Volume = base.ParseFloat(mktStrArr[10])
				mkt.Amount = base.ParseFloat(mktStrArr[15])
				mkt.Position = base.ParseFloat(mktStrArr[9])
				mkt.PreSettlementPrice = base.ParseFloat(mktStrArr[8])
				ret = append(ret, mkt)
				// fmt.Println(mkt.Symbol, mkt.Open, mkt.High, mkt.Low)
			}
		}

	}
	return ret, nil
}
