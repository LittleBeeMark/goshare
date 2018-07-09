## goshare

 * goshare是一个开源的golang量化数据工具集合
 * 集成A股行情数据,基本面数据， 中国期货 等数据的抓取功能
 * 封装相关市场的交易接口
 * 集成量化交易框架
 * 量化交易webui控制台
 * 组建测试集群，充分利用多机多核

<!-- [START getstarted] -->
## Getting Started

### Installation
 ```
 go get -u github.com/mineralres/goshare
```
### Usage
```
import (
  "github.com/mineralres/goshare"
)

func test(){
  var s goshare.Service
  symbol := pb.Symbol{Exchange: pb.ExchangeType_SHFE, Code: "rb1805"}
  // 获取历史数据
  data, err := s.GetKData(&symbol, pb.PeriodType_M5, 19990101, 20180307, 1)
  if err != nil {
    panic(err)
  }
}
```

## API Documentation
### 数据接口
 * GetKData 请求历史K线数据
 * GetLastTick 请求指定symbol的最新行情
 * GetSina50EtfSym 获取50ETF期权合约列表(sina代码)
 * GetIndexMember 获取指数成份股
 * GetRealtimeMoneyTrendList 取实时资金流向
 ### 交易接口(待实现)
 * CTP
 * LTS
 

# FAQ
* QQ群 249379339 
