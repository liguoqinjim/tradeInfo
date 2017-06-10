package db

import (
	"gopkg.in/mgo.v2"
	"tradeInfo/conf"
	"tradeInfo/log"
)

var session mgo.Session

func init() {

}

func Dial() {
	s, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{conf.Conf.DBConf.Address},
		Username: conf.Conf.DBConf.Username,
		Password: conf.Conf.DBConf.Password,
		Database: conf.Conf.DBConf.DBName,
	})
	if err != nil {
		log.Fatal("err=%v", err)
	}
	session = s
}

func Close() {
	session.Close()
}
