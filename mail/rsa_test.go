package mail

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestRsaDecrypt(t *testing.T) {
	data, err := RsaEncrypt([]byte("{\"code\":0,\"message\":\"\",\"data\":{\"crawlId\":\"8faf253bfe364153877a56192f3c4e05\",\"business\":\"\",\"platform\":\"shanjiekuan\",\"phone\":\"13515023152\",\"bills\":[{\"id\":\"ea00c8a4282d055795ccec95b56c19a5\",\"borrowAmount\":6000,\"amount\":6000,\"hadRepayAmount\":0,\"hadRepayPeriod\":0,\"totalPeriod\":1,\"status\":false,\"borrowTime\":\"2017-09-28T00:00:00+08:00\",\"periodBills\":[{\"id\":\"ea00c8a4282d055795ccec95b56c19a5\",\"repayAmount\":6000,\"repayDay\":\"2017-12-29T00:00:00+08:00\",\"status\":false}]},{\"id\":\"355401793180c1f392f26dada1724960\",\"borrowAmount\":6000,\"amount\":6000,\"hadRepayAmount\":6000,\"hadRepayPeriod\":1,\"totalPeriod\":1,\"status\":true,\"borrowTime\":\"2017-08-22T00:00:00+08:00\",\"periodBills\":[{\"id\":\"355401793180c1f392f26dada1724960\",\"repayAmount\":6000,\"repayDay\":\"2017-12-01T00:00:00+08:00\",\"status\":true}]},{\"id\":\"c50e3b63e21d29229c27316723be93ca\",\"borrowAmount\":2000,\"amount\":2000,\"hadRepayAmount\":2000,\"hadRepayPeriod\":1,\"totalPeriod\":1,\"status\":true,\"borrowTime\":\"2017-06-11T00:00:00+08:00\",\"periodBills\":[{\"id\":\"c50e3b63e21d29229c27316723be93ca\",\"repayAmount\":2000,\"repayDay\":\"2017-09-13T00:00:00+08:00\",\"status\":true}]}]}}"))
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(data))
	origData, err := RsaDecrypt(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))

}
