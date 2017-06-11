package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type ConfStruct struct {
	Username      string
	Password      string
	Address       string
	ConnectDBName string
	DataDBName    string
}

var conf *ConfStruct

func readConf() {
	file, err := os.Open("conf.json")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	c := new(ConfStruct)
	err = json.Unmarshal(data, c)
	if err != nil {
		log.Fatal(err)
	}

	conf = c
}

type Trade struct {
	TradeType string  `json:"trade_type"`
	Date      int     `json:"date"`
	Price     float64 `json:"price"`
	Amount    int     `json:"amount"`
	Tid       int     `json:"tid"`
}

func main() {
	readConf()
	fmt.Println(conf)

}
