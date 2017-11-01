package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://japi.wolaidai.com/jrocket2/api/v3/sessions/18016333986?timestamp=1509515438&sign=7a8384f713addf6697d881ddea9cb2e9"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("host", "japi.wolaidai.com")
	req.Header.Add("accept", "application/json")
	req.Header.Add("x-user-identity", "0")
	req.Header.Add("x-origin", "AppStore")
	req.Header.Add("accept-language", "zh-cn")
	req.Header.Add("x-source-id", "3")
	req.Header.Add("x-product-code", "WLD")
	req.Header.Add("accept-encoding", "gzip, deflate")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("user-agent", "%E6%88%91%E6%9D%A5%E8%B4%B7/75 CFNetwork/887 Darwin/17.0.0")
	req.Header.Add("connection", "keep-alive")
	req.Header.Add("x-app-version", "4.7.3")
	req.Header.Add("cache-control", "no-cache")


	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}