package main

import (
	"time"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"encoding/base64"
	"go-demo/rsa"
	"encoding/json"
	"log"
)

var RSA = rsa.RSASecurity{}
// 获取所有的错误日志
func getAllErrorLog(dateStr string) {
	Connect()
	defer SafeClose()
	// 查询多条数据
	var results []LogsResolveError
	loc, _ := time.LoadLocation("Local")
	endTime, err := time.ParseInLocation("2006-01-02 15:04:05", dateStr, loc)
	if err != nil {
		log.Fatal("时间格式不正确 :", err.Error())
	}
	err = Logs_resolve_error_collection.Find(bson.M{"createTime": bson.M{"$gte": endTime}}).Sort("platform").All(&results)
	if err != nil {
		log.Fatal("查询所有数据异常 :", err.Error())
	}

	myMap := make(map[string]string)
	for i := 0; i < len(results); i++ {
		// 获取账号密码
		var trafficItem Traffic
		err = Logs_traffic_collection.Find(bson.M{"crawlId": results[i].CrawlId}).One(&trafficItem)
		if err != nil {
			fmt.Println("查询详情发生了错误", err)
			break
		}
		decodeBytes, err := base64.StdEncoding.DecodeString(trafficItem.RequestBody)
		if err != nil {
			fmt.Println("解密发生了错误", err)
			break
		}
		pridecrypt, err := RSA.RsaDecrypt(decodeBytes)
		if err != nil {
			fmt.Println(err)
			break
		}

		var jsonName MyJsonName
		err = json.Unmarshal([]byte(string(pridecrypt)), &jsonName)
		if err != nil {
			fmt.Println("转json出错了", err)
			break
		}
		jsonStr, err := json.Marshal(jsonName)
		if err != nil {
			fmt.Println("转string出错了", err)
			break
		}
		mapValue := string(jsonStr)
		mapKey := results[i].Platform + "----" + results[i].ErrMessage
		myMap[mapKey] = results[i].Platform + "----" + mapValue + "----" + results[i].ErrMessage

	}
	err = WriteMaptoFile(myMap, "C:/Users/user/Desktop/错误日志的平台账号和密码.txt")
	if err != nil {
		fmt.Println("写文件出错", err)
	}
}
