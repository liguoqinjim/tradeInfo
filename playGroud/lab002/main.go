package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"log"
	"os"
	"time"
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

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Username: conf.Username,
		Password: conf.Password,
		Addrs:    []string{conf.Address},
		Database: conf.ConnectDBName,
		Timeout:  time.Second * 5,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	//insert
	trade := Trade{TradeType: "bid", Date: int(time.Now().Unix()), Price: 0.2222, Amount: 4567, Tid: 2011221}
	c := session.DB(conf.DataDBName).C("test1")
	if err := c.Insert(trade); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("insert 成功")
	}
}
