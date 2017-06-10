package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"tradeInfo/db"
	"tradeInfo/msgMgr"
	"tradeInfo/net"
)

func main() {
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
	fmt.Println("exiting")
}
