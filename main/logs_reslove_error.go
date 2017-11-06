package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "http://api.ffrpbank.com/credit-user/login?clientType=ios&appName=ffrp&deviceId=962ba760f935486bafaefd321cb09da3&appVersion=1.0.0&mobilePhone=13752347055"

	payload := strings.NewReader("password=123456&username=18321782870")

	req, _ := http.NewRequest("POST", url, payload)



	req.Header.Add("content-type", "application/x-www-form-urlencoded")


	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}