package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
	"tradeInfo/conf"
	"tradeInfo/db"
	"tradeInfo/log"
	"tradeInfo/msgMgr"
	"tradeInfo/net"
)

func main() {
	//设置是否输出日志到文件
	if conf.Conf.LogConf.FileLog {
		log.Info("日志输出至文件")
		log.SetFilelog()
	}
	log.Infof("系统开始时间:%v", time.Now())

	//websocket模块
	net.Connect()
	net.InitSendMessage()

	//mongodb模块
	db.Dial()

	//消息处理模块
	msgMgr.Handle()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	log.Infof("系统结束时间:%v", time.Now())
}
