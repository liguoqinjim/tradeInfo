package net

import (
	"github.com/gorilla/websocket"
	"time"
	"tradeInfo/log"
	"tradeInfo/msgMgr"
	"tradeInfo/util"
)

func init() {

}

var conn *websocket.Conn

func Connect() {
	dialer := &websocket.Dialer{}

	con, _, err := dialer.Dial("wss://io.sosobtc.com/socket.io/?EIO=3&transport=websocket",
		map[string][]string{"Origin": []string{"https://k.sosobtc.com"},
			"Cookie":                []string{"OID=aEj%252BuelTgv0RAuNv%252FFJPBfslpUnvZ26EiWsVM7TiIgNZ%252FaJQtLXiwjyAsqbnRKA%252BpQ7UkYv1rrO92kq8%252BZ4ifZQex9e7Sbgj7BVy3DtSflfIJd4koi1JTx61ElPwSY8x%7C8dad2860013668cf3e1c4aa6c4e19154; _ga=GA1.2.241055403.1496717869; _gid=GA1.2.1622829335.1496717869; _gat=1; theme=dark"},
			"User-Agent":            []string{"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko"},
			"Sec-WebSocket-Version": []string{"13"}})
	if err != nil {
		log.Fatalf("err=%v", err)
	}
	conn = con

	go ReadMessage()
}

func ReadMessage() {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Fatalf("read message err:%v", err)
		}
		msgMgr.MessageChan <- message
	}
}

func InitSendMessage() {
	go func() {
		time.Sleep(time.Millisecond * 500)
		SendMessage(websocket.TextMessage, util.SEND_INFO_TRADE_SC)
	}()
}

func SendMessage(messageType int, message string) error {
	err := conn.WriteMessage(messageType, []byte(message))
	return err
}

func Close() {
	conn.Close()
}
