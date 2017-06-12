package db

import (
	"gopkg.in/mgo.v2"
	"time"
	"tradeInfo/conf"
	"tradeInfo/log"
	"tradeInfo/util"
)

var session *mgo.Session

func init() {

}

func Dial() {
	s, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{conf.Conf.DBConf.Address},
		Username: conf.Conf.DBConf.Username,
		Password: conf.Conf.DBConf.Password,
		Database: conf.Conf.DBConf.ConnectDBName,
		Timeout:  time.Second * 10,
	})
	if err != nil {
		log.Fatalf("err=%v", err)
	} else {
		log.Info("连接mongodb success")
	}
	s.SetMode(mgo.Monotonic, true)
	session = s
}

func InsertTradeInfo(d interface{}) error {
	coll := session.DB(conf.Conf.DBConf.DataDBName).C(util.COLLECTION_TRADE_INFO)
	if err := coll.Insert(d); err != nil {
		return err
	} else {
		return nil
	}
}

func Close() {
	session.Close()
}
