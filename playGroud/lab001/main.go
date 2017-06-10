package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
)

type Trade struct {
	TradeType string  `json:"trade_type"`
	Date      int     `json:"date"`
	Price     float64 `json:"price"`
	Amount    int     `json:"amount"`
	Tid       int     `json:"tid"`
}

type Trades struct {
	Ts []Trade `json`
}

func main() {
	//最原始的字符串 := `42["update:trades",[{"trade_type":"bid","date":1497080954,"price":0.1259,"amount":3030,"tid":30322061},{"trade_type":"bid","date":1497080948,"price":0.1258,"amount":26843,"tid":30322044},{"trade_type":"bid","date":1497080947,"price":0.1258,"amount":9458,"tid":30322042},{"trade_type":"bid","date":1497080943,"price":0.1257,"amount":1001,"tid":30322004}]]`
	text := `42["update:trades",[{"trade_type":"bid","date":1497080954,"price":0.1259,"amount":3030,"tid":30322061},{"trade_type":"bid","date":1497080948,"price":0.1258,"amount":26843,"tid":30322044},{"trade_type":"bid","date":1497080947,"price":0.1258,"amount":9458,"tid":30322042},{"trade_type":"bid","date":1497080943,"price":0.1257,"amount":1001,"tid":30322004}]]`
	text = strings.Replace(text, `42["update:trades",[`, `{"update:trades":[`, 1)
	fmt.Println(text)

	result := gjson.Get(text, "update:trades")
	fmt.Println(len(result.Array()))

	//2
	text = `[{"trade_type":"bid","date":1497080954,"price":0.1259,"amount":3030,"tid":30322061},{"trade_type":"bid","date":1497080948,"price":0.1258,"amount":26843,"tid":30322044},{"trade_type":"bid","date":1497080947,"price":0.1258,"amount":9458,"tid":30322042},{"trade_type":"bid","date":1497080943,"price":0.1257,"amount":1001,"tid":30322004}]`
	trades := make([]Trade, 0)
	gjson.Unmarshal([]byte(text), &trades)
	fmt.Println(trades)

	//3
	text = `42["update:trades",[{"trade_type":"bid","date":1497080954,"price":0.1259,"amount":3030,"tid":30322061},{"trade_type":"bid","date":1497080948,"price":0.1258,"amount":26843,"tid":30322044},{"trade_type":"bid","date":1497080947,"price":0.1258,"amount":9458,"tid":30322042},{"trade_type":"bid","date":1497080943,"price":0.1257,"amount":1001,"tid":30322004}]]`
	text = strings.Replace(text, `42["update:trades",`, ``, 1)
	text = text[:len(text)-1]
	fmt.Println(text)
	trades2 := make([]Trade, 0)
	gjson.Unmarshal([]byte(text), &trades2)
	fmt.Println(trades)
}
