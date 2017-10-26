package logs_revlove_error

import (
	"time"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"encoding/base64"
	"go-demo/rsa"
	"encoding/json"
)

var Pirvatekey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpgIBAAKCAQEArYkPlffsyC0G1wZPhF0pcQFudgBltGfUiW4xlR3IWUbpxZaL
lWDdmVWCJW6NRs4Qqif5I0+zNGbW3ml0a0NCUxtuHlVhGuo80T6D4FqWXEyJh4Qv
ey6k2J9suVahlcZ8SfkXtcojCIwotQ/OlDjvpkfBvCzGEcj1UjiWlso7msSFJ7eQ
BJ8vsue/nO6APpN9ZoacQIF2QyPPrfHOlLN2tNjnNR4YCyDdZdekHALJTAIhN459
Cp7bL60QvwrEx1jXFeb6NhETMEvmG3I+ypQUac4dIrSxQs6pIkgYGU/AFbTcJuo6
XMmigmWuuCfrNEEqTK+uMezo4uK7stF6jArI7wIDAQABAoIBAQCLYqvCKYFmx8PW
sprsFmhS/HNdFLScU0nDmV76BxIFo4/hxSoYsdVMdAI1TrbrSFjaU4Epe7rVPEUa
IFoCTePYHRA2DR4SIFL5Pt1uN1TOjitpTiNVLgH6fRM3Sv4+706lnA4PVm0NUIbh
5/Bl3dWgGcLjApOVdXSWth0+wPFfPWcZlykFCXir4JiC5Q18LHkTopcIRf6pj8hf
9+EUSXFcSUlpIr+tVLgFwZ39eE8XdF/RLaySWeaRUQV7NH5lrpnUQ7Qo6h+hF56p
sXpwEFdQxvc9CdWOZh4eULRHaSRInhF/RuFSOiW9CXDbSzQOVYUEFbk/XvKX4qr6
dblEUFQBAoGBAN7EvLsJklOfwhYCIh6HzBGDjJYv5arQeLwxKU2xlW9MGj/zK06v
rkJu3jLChMIR37HjVebg+4pQaKB54wxop6OoRXRMdJf7aIrTsyQwI7EOChDHtffS
+JlaTnMkBtkBvbew/bdEdCH0MWzK6AcwSxKnaOJtVuUy3dki4IX2jwNfAoGBAMds
K4d/tPGaQr4hRMEa38A728frHixTUVmYy0yxu82lJwx6iTsx8+bviRCOc4pXOBvj
VCSOwvEcAFIFEVdx7R4/48cOqGw2VM6XMP/79lf+Fsz/U+pXHPcD+T4hLttR5Fc0
6RXhk6KGMYtbmwLcOxhMmNKDQyKzfrlWC8hCpTRxAoGBAJwWfacQLdcqOXPySrXZ
sMMUyYzyQppNOEukpqwpoes6B6qHV/DqZ89kn4wxf5P7siGSiL4dJiLj9Gt8+/nV
2qBxZzj12TwkjDu52euhT2Df1W2sgXQD/ux0ukMzGxkyfxwWh6/B5x1jXL497hYb
YH7RjkwZEKnH2moqZed3uv1BAoGBALP2k9m1j46wTSbTUv6Ix/7shcwGC7aogbM/
RceyZO6R6YlDU3BOBfMTmzeLtgzT5NiV76sGAjofsNJTrHCjaP5xXZaXTbpJ1YKR
wjBgKtEBhYX9rLa/Ym4v8c71lxlo740bPHT4moiCXVeLlXmhUNr+1LK0rfM01Flg
o0VXj0qhAoGBAJznWoEleWIcViBnRp8Y76chUJ1wj5dqRPZeIBXNaJMiNfGL/dCA
cKMjiN69g5Tiuxxu/7/Cao4axnqIrG+0vp/St6gBs2XPw5lfp+502voK6zBW7dFS
35ABLg+hy8ICHR8xGcQhgnDqZOc/0T2btHjPz1datOrn9MGdHAEDmFQ4
-----END RSA PRIVATE KEY-----
`
// 错误日志实体
type LogsResolveError struct {
	CrawlId    string    `json:"crawlId"`                      // 此次爬取的唯一 ID
	Platform   string    `bson:"platform" json:"platform"`     // 平台唯一标识
	ErrMessage string    `bson:"errMessage" json:"errMessage"` // 错误信息
	CreateTime time.Time `bson:"createTime" json:"createTime"` // 创建时间
}

type Traffic struct {
	CrawlId     string `bson:"crawlId" json:"crawlId"`         // 爬取任务唯一 ID
	RequestBody string `bson:"requestBody" json:"requestBody"` // 请求体
}

type MyJsonName struct {
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Platform string `json:"platform"`
}

var RSA = rsa.RSASecurity{}
// 获取所有的错误日志
func getAllErrorLog() {
	Connect()
	defer SafeClose()
	// 查询多条数据
	var results []LogsResolveError
	loc, _ := time.LoadLocation("Local")
	endTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2017-10-26 15:58:44", loc)
	if err != nil {
		fmt.Println("时间发生了错误", err)
	}
	err = Logs_resolve_error_collection.Find(bson.M{"createTime": bson.M{"$gte": endTime}}).Sort("platform").All(&results)
	if err != nil {
		fmt.Println("查询发生了错误", err)
	}

	myMap := make(map[string]string)

	for i := 0; i < len(results); i++ {
		// 获取账号密码
		var trafficItem Traffic
		err = Logs_traffic_collection.Find(bson.M{"crawlId": results[i].CrawlId}).One(&trafficItem)
		if err != nil {
			fmt.Println("查询发生了错误", err)
		}
		mapKey := results[i].Platform + results[i].ErrMessage
		decodeBytes, err := base64.StdEncoding.DecodeString(trafficItem.RequestBody)
		if err != nil {
			fmt.Println("解密发生了错误", err)
		}
		pridecrypt, err := RSA.RsaDecrypt(decodeBytes)
		if err != nil {
			fmt.Println(err)
		}
		mapValue := string(pridecrypt)

		var jsonName MyJsonName
		err = json.Unmarshal([]byte(mapValue), &jsonName)
		if err != nil {
			fmt.Println("转json出错了", err)
		}
		jsonStr, err := json.Marshal(jsonName)
		if err != nil {
			fmt.Println("转string出错了", err)
		}
		mapValue = string(jsonStr)
		myMap[mapKey] = mapValue

	}
	for key, value := range myMap {
		fmt.Println(value, "----------", key)
	}

}
