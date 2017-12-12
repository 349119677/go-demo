package main

import (
	"github.com/sclevine/agouti"
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	driver := agouti.PhantomJS()
	err := driver.Start()
	if err != nil {
		log.Panic(err.Error())
	}
	defer driver.Stop()
	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Panic(err.Error())
	}
	err = page.Navigate("http://www.pubgtracker.com/")
	if err != nil {
		log.Panic(err.Error())
	}

	redirectURL := "https://pubgtracker.com/"
	// 判断是否请求成功
	for true {
		sleepTime := 0
		url, err := page.URL()
		if err != nil {
			log.Panic(err.Error())
		}
		if strings.Contains(url, redirectURL) || sleepTime >= 10000 {
			break
		} else {
			sleepTime += 300
			time.Sleep(time.Millisecond * 300)
		}
	}

	err = page.Navigate(redirectURL)
	if err != nil {
		log.Panic(err.Error())
	}
	page.Screenshot("2.png")
	cookies, err := page.GetCookies()
	if err != nil {
		log.Panic(err.Error())
	}
	fmt.Println(cookies)

}
