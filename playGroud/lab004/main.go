package main

import (
	"github.com/gorilla/websocket"
	"golang.org/x/text"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	dialer := &websocket.Dialer{}

	conn, _, err := dialer.Dial("wss://io.sosobtc.com/socket.io/?EIO=3&transport=websocket",
		map[string][]string{"Origin": []string{"https://k.sosobtc.com"},
			"Cookie":                []string{"OID=aEj%252BuelTgv0RAuNv%252FFJPBfslpUnvZ26EiWsVM7TiIgNZ%252FaJQtLXiwjyAsqbnRKA%252BpQ7UkYv1rrO92kq8%252BZ4ifZQex9e7Sbgj7BVy3DtSflfIJd4koi1JTx61ElPwSY8x%7C8dad2860013668cf3e1c4aa6c4e19154; _ga=GA1.2.241055403.1496717869; _gid=GA1.2.1622829335.1496717869; _gat=1; theme=dark"},
			"User-Agent":            []string{"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko"},
			"Sec-WebSocket-Version": []string{"13"}})
	if err != nil {
		log.Fatalf("err=%v", err)
	}

	go func() {
		time.Sleep(time.Millisecond * 500)
		//SendMessage(websocket.TextMessage, util.SEND_INFO_TRADE_SC)
		conn.WriteMessage(websocket.TextMessage, []byte(`4212["market.subscribe","sc:yunbi"]`))
	}()

	//收消息
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Fatal("read message error", err)
			} else {
				log.Println("收到消息", string(message))
			}
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	log.Println("程序结束", time.Now())
}
