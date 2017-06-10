# tradeInfo
采集交易信息并存储到mongodb

## todolist
1. 容错机制，链接断开或者其他情况可以切换到其他的服务
2. 数据要判断去重

## 系统架构
|模块名|简介|
|---|---|
|conf|读取配置文件|
|db|和mongodb交互|
|log|日志|
|net|和websocket交互|
|msgMgr|消息处理模块|