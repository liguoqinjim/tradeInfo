package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"log"
	"strings"
)

type Trade struct {
	TradeType string  `json:"trade_type"`
	Date      int     `json:"date"`
	Price     float64 `json:"price"`
	Amount    int     `json:"amount"`
	Tid       int     `json:"tid"`
}

func main() {
	text := `42["update:trades",[{"trade_type":"bid","date":1497080954,"price":0.1259,"amount":3030,"tid":30322061},{"trade_type":"bid","date":1497080948,"price":0.1258,"amount":26843,"tid":30322044},{"trade_type":"bid","date":1497080947,"price":0.1258,"amount":9458,"tid":30322042},{"trade_type":"bid","date":1497080943,"price":0.1257,"amount":1001,"tid":30322004}]]`
	text = strings.Replace(text, `42["update:trades",[`, `{"update:trades":[`, 1)
	fmt.Println(text)

	result := gjson.Get(text, "update:trades")
	fmt.Println(len(result.Array()))

	for _, v := range result.Array() {
		trade := &Trade{}
		err := gjson.Unmarshal([]byte(v.Raw), trade)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(trade)
		}
	}
}
