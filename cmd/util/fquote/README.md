# fquote

用CTP接收期货tick行情，推送到datasource中生成相应的k线

- 调用交易接口查询合约
- 订阅全部合约数据
- 用shell脚本控制每日定时重启