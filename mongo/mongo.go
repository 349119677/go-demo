package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
	//"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/bson"
)

var (
	session *mgo.Session
	db      *mgo.Database
)

const (
	DATABASE_NAME = "myDB"
)

// 连接
func Connect() {
	var err error
	session, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	db = session.DB(DATABASE_NAME)
}

// 关闭
func SafeClose() {
	if session != nil {
		session.Close()
	}
}

// 向mongo插入数据
func InsertIntoMongo() {

	logsOriginalBill := new(LogsOriginalBill)
	logsOriginalBill.Id = 1473
	logsOriginalBill.Crawl_id = "e0d0a8fcd8174eb5b3fcae4ad4db8d47"
	logsOriginalBill.Platform = "xianjinxia"
	logsOriginalBill.Url = "http://super.xianjinxia.com/credit-user/login"
	logsOriginalBill.Action = "LOGIN"
	logsOriginalBill.Request_header = "{\"Cookie\":[\"ssid=7810863b52c63de8\"]}"
	logsOriginalBill.Request_body = "{\"appMarket\":[\"yingyongbao\"],\"appName\":[\"mld\"],\"appVersion\":[\"1.0.0\"],\"clientType\":[\"android\"],\"deviceId\":[\"862555023499809\"],\"deviceName\":[\"HUAWEI G750-T00\"],\"osVersion\":[\"4.4.2\"],\"password\":[\"heyang5206\"],\"username\":[\"18508313576\"]}"
	logsOriginalBill.Response_header = "{\"Connection\":[\"keep-alive\"],\"Content-Type\":[\"application/json;charset=UTF-8\"],\"Date\":[\"Thu, 27 Jul 2017 12:27:20 GMT\"],\"Set-Cookie\":[\"acw_tc=AQAAAKuueDLpRAUAYLIOaqdky1Q98dzx; Path=/; HttpOnly\",\"JSESSIONID=A31F8E6732AFE4F41F790FCE26FC4BB4; Path=/; HttpOnly\"]}"
	logsOriginalBill.Response_body = "{\"code\":\"-1\",\"message\":\"你输入的用户或密码不正确，请重新输入。\",\"data\":{\"item\":{}}}"
	logsOriginalBill.Create_time = "2017-07-27 20:27:20"

	Connect()
	startTime2 := time.Now().UnixNano()
	defer func() {
		SafeClose()
		fmt.Printf("Insert time spend: %d", time.Now().UnixNano()-startTime2)
	}()
	collection := db.C("test")

	for i := 1; i <= 10000000; i++ {
		err := collection.Insert(logsOriginalBill)
		if err != nil {
			panic(err)
		}
	}
}

// 根据主键查询
func SelectByPrimaryKey() {
	type AAAA struct {
		Id              int    "bson:`id`"
		Crawl_id        string "bson:`crawl_id`"
		Platform        string "bson:`platform`"
		Url             string "bson:`url`"
		Action          string "bson:`action`"
		Request_header  string "bson:`request_header`"
		Request_body    string "bson:`request_body`"
		Response_header string "bson:`response_header`"
		Response_body   string "bson:`response_body`"
		Create_time     string "bson:`create_time`"
	}

	Connect()
	startTime2 := time.Now().UnixNano()
	defer func() {
		SafeClose()
		fmt.Printf("Insert time spend: %d", time.Now().UnixNano()-startTime2)
	}()
	collection := db.C("test")
	var result AAAA
	//ms := []LogsOriginalBill{}
	id := "59c4b0fa09ae418aac652949"
	objectId := bson.ObjectIdHex(id)
	//collection.Find(bson.M{"_id": objectId}).One(&result)
	collection.UpdateAll(bson.M{"_id": objectId},bson.M{"$set": bson.M{"id": 79797988}})
	//for _, m := range ms {
	//	fmt.Printf(m.Create_time)
	//}
	fmt.Println(result.Create_time)

}
