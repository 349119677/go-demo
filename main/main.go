package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {
	
	url := "https://zchlhd.com/api/accounts/login.json"
	
	payload := strings.NewReader("user%5Bpassword%5D=aa123456&user%5Blogin%5D=15926901283")
	
	req, _ := http.NewRequest("POST", url, payload)
	
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	
	res, _ := http.DefaultClient.Do(req)
	
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	
	fmt.Println(res)
	fmt.Println(string(body))
	
}