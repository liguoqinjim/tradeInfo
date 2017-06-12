package msgMgr

import (
	"github.com/tidwall/gjson"
	"strings"
	"tradeInfo/db"
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

type Trade struct {
	TradeType string  `json:"trade_type"`
	Date      int     `json:"date"`
	Price     float64 `json:"price"`
	Amount    int     `json:"amount"`
	Tid       int     `json:"tid"`
}

func HandleMessage(message []byte) {
	s := string(message)

	//判断是否是sc交易记录
	if strings.Contains(s, "update:trades") {
		log.Info(s)
		s = strings.Replace(s, `42["update:trades",[`, `{"update:trades":[`, 1)
		s = s[:len(s)-1]
		s += "}"
		log.Info(s)
		result := gjson.Get(s, "update:trades")

		log.Infof("开始插入数据,%d", len(result.Array()))
		if len(result.Array()) > 0 {
			for _, v := range result.Array() {
				err := db.InsertTradeInfo(&v)
				if err != nil {
					log.Error("数据insert fail")
				} else {
					log.Info("插入数据成功")
				}
			}
		}
	} else {
		//log.Info(s)
	}
}
