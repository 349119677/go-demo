package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "http://jybkapi.51meishidi.com/v1/notice/loan?sign=760D189FBDB864D25915422A1CBD1F03"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("host", "jybkapi.51meishidi.com")
	req.Header.Add("connection", "keep-alive")
	req.Header.Add("accept", "*/*")
	req.Header.Add("user-agent", "JYBK/1.1.0927 (iPhone; iOS 11.0.3; Scale/2.00)")
	req.Header.Add("accept-language", "zh-Hans-CN;q=1")
	req.Header.Add("authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOi8vanlia2FwaS41MW1laXNoaWRpLmNvbS92MS9hdXRob3JpemF0aW9ucyIsImlhdCI6MTUwODQ2NDU2OSwiZXhwIjo3NTA4NDY0NTA5LCJuYmYiOjE1MDg0NjQ1NjksImp0aSI6IllLZ1pKVDk2WU9YbHFzR0QiLCJzdWIiOjUzNTk3ODZ9.vxEw4slmOMmtj6QNuMGagIAtayq0CpKS5ZWTGqQZz8c")
	req.Header.Add("accept-encoding", "gzip, deflate")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}