package goshare

import (
	"testing"
	"log"
	//"net/http"
	//"io/ioutil"
	//"strings"
	"github.com/mineralres/goshare/aproto"
)

var s Service

func TestKData(t *testing.T) {
	symbol := aproto.Symbol{Exchange: aproto.ExchangeType_SHFE, Code: "rb1805"}
	k, err := s.GetKData(&symbol, aproto.PeriodType_M5, 19990101, 20180307, 1)
	if err != nil {
		t.Error(err)
	}
	if len(k.List) == 0 {
		t.Error("GetKData error")
	}
	// log.Println()
}

func TestGetLastTick(t *testing.T) {
	symbol := aproto.Symbol{Exchange: aproto.ExchangeType_SSE, Code: "600000"}
	md, err := s.GetLastTick(&symbol)
	if err != nil {
		t.Error(err)
	}
	if len(md.OrderBookList) == 0 {
		t.Error("获取行情盘口深度为空")
	}
	if md.Open == 0 {
		t.Error("md.Open == 0")
	}
	// log.Printf("Tick[%s], Open[%.2f], High[%.2f], Low[%.2f], Close[%.2f]", md.Symbol.Code, md.Open, md.High, md.Low, md.Close)
}

func TestIndexTick(t *testing.T) {
	//测试获取sina各种指数
	log.Printf("测试获取sina各种指数")
	m_index := map[string]string{
		"道琼斯指数" : "int_dji",
		"上证指数" : "sh000001",
		"纳斯达克" : "int_nasdaq",
		"恒生指数" : "int_hangseng",
		"日经指数" : "b_TWSE",
		"新加坡指数" : "b_FSSTI",
	}		
	for key, views := range m_index {		
		symbol := aproto.Symbol{Exchange: aproto.ExchangeType_INDEX, Code: views}
		md, err := s.GetLastTick(&symbol)
		if err != nil {
			t.Error(err)
		}
		if (md.Close) <= 0 {
			t.Error("获取行情为空")
		}
		md.Symbol.Code = key
		log.Printf("Tick[%s],Close[%.2f]", md.Symbol.Code,  md.Close)
	}
}


func TestOptionSSETick(t *testing.T) {
	// 获取sina50etf期权的合约列表：
	// 同一个月份看涨和看跌: OP_UP_5100501804   OP_DOWN_5100501804
	sym  :="OP_DOWN_5100501804"
	syms := s.GetSina50EtfSym(sym)
	for _, value := range syms {
		//log.Printf("Index: %d  Value: %s\n", index, value)
		symbol := aproto.Symbol{Exchange: aproto.ExchangeType_OPTION_SSE, Code: value}
		md, err := s.GetLastTick(&symbol)
		if err != nil {
			t.Error(err)
		}
		if (md.Close) <= 0 {
			t.Error("获取行情为空")
		}
		log.Printf("Tick[%s], Open[%.2f], High[%.2f], Low[%.2f], Close[%.2f]", md.Symbol.Code, md.Open, md.High, md.Low, md.Close)
	  }
}

