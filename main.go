package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"tradeInfo/db"
	"tradeInfo/log"
	"tradeInfo/msgMgr"
	"tradeInfo/net"
)

func main() {
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
	fmt.Println("系统结束时间:%v", time.Now())
}
