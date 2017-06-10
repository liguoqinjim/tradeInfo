package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"tradeInfo/log"
)

type ConfStruct struct {
	DBConf *DBConf
}

type DBConf struct {
	Username      string
	Password      string
	Address       string
	ConnectDBName string
	DataDBName    string
}

var Conf *ConfStruct

func readConf() {
	file, err := os.Open("conf.json")
	if err != nil {
		log.Fatalf("err=%v", err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("err=%v", err)
	}

	Conf = new(ConfStruct)
	err = json.Unmarshal(data, Conf)
	if err != nil {
		log.Fatalf("err=%v", err)
	}
}

func init() {
	readConf()
}
