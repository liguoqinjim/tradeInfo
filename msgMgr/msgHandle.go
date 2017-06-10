package msgMgr

import (
	"tradeInfo/log"
)

var MessageChan chan []byte

func init() {
	MessageChan = make(chan []byte, 10)
}

func Handle() {
	go func() {
		for {
			select {
			case m := <-MessageChan:
				go HandleMessage(m)
			}
		}
	}()
}

func HandleMessage(message []byte) {
	log.Info(string(message))
}
