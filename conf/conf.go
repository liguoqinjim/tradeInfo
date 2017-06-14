package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"tradeInfo/log"
)

type ConfStruct struct {
	LogConf *LogConf
	DBConf  *DBConf
}

type LogConf struct {
	FileLog bool //是否在文件中输出
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
