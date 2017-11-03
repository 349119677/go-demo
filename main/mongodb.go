package main

import (
	"gopkg.in/mgo.v2"
	"log"
)

var (
	Session                       *mgo.Session
	Mongo                         *mgo.Database
	Logs_resolve_error_collection *mgo.Collection
	Logs_traffic_collection       *mgo.Collection
)

func Connect() {
	var err error
	Session, err = mgo.Dial("dds-uf6eeab5a7b4e4541.mongodb.rds.aliyuncs.com:3717")
	if err != nil {
		log.Fatal("连接数据库发生错误 :", err.Error())
	}

	if err = Session.Login(&mgo.Credential{
		Username:  "readaccount",
		Password:  "readaccount",
		Mechanism: "SCRAM-SHA-1",
		Source:    "cloan",
	}); err != nil {
		log.Fatal("登录数据库发生错误 :", err.Error())
	}
	Session.SetMode(mgo.Monotonic, true)
	Mongo = Session.DB("cloan")
	// 返回需要的两张表的实例
	Logs_resolve_error_collection = Mongo.C("logs_resolve_error")
	Logs_traffic_collection = Mongo.C("logs_traffic")
}

func SafeClose() {
	if Session != nil {
		Session.Close()
	}
}
