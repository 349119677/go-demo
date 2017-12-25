package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://sb.1eightyeightbet.com/zh-cn/Service/CentralService?GetData="

	payload := strings.NewReader("reqUrl=%2Fzh-cn%2Fsports%2Fe-sports%2Fcompetition%2Foutright%3Fcompetitionids")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}