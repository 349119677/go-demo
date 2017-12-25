package com_pubgtracker

import (
	"math/rand"
	"github.com/sclevine/agouti"
	"strings"
	"time"
)

// 浏览器伪装
type UserAgent struct {
	UserAgentItemList []*UserAgentItem
}
type UserAgentItem struct {
	UserAgent string // 浏览器标识
	InCookie  string // puBgTracker的cookie
}

var (
	userAgentList []string
	apiKeyList    []string
)

// 初始化认证对象
func Init() (*UserAgent, []string) {
	userAgentList = []string{
		// 浏览器标识对象集合
		"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",                    // safari 5.1 – MAC
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",                             // safari 5.1 – Windows
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.6; rv,2.0.1) Gecko/20100101 Firefox/4.0.1",                                                          // Firefox 4.0.1 – MAC
		"Mozilla/5.0 (Windows NT 6.1; rv,2.0.1) Gecko/20100101 Firefox/4.0.1",                                                                          // Firefox 4.0.1 – Windows
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.8.131 Version/11.11",                                                            // Opera 11.11 – MAC
		"Opera/9.80 (Windows NT 6.1; U; en) Presto/2.8.131 Version/11.11",                                                                              // Opera 11.11 – Windows
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",                       // Chrome 17.0 – MAC
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Maxthon 2.0)",                                                                              // 傲游（Maxthon）
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; TencentTraveler 4.0)",                                                                      // 腾讯TT
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1)",                                                                                           // 世界之窗（The World） 2.x
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; The World)",                                                                                // 世界之窗（The World） 3.x
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; SE 2.X MetaSr 1.0; SE 2.X MetaSr 1.0; .NET CLR 2.0.50727; SE 2.X MetaSr 1.0)", // 搜狗浏览器 1.x
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; 360SE)",                                                                                    // 360浏览器
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; 360SE)",                                                                                    // Avant
	}
	// 授权码对象集合
	apiKeyList = []string{
		"fb485316-47ea-4ca7-bcdc-5abfe1f35580",
	}
	var userAgentResponse = UserAgent{}.UserAgentItemList
	// 初始化userAgent
	for i := 0; i < len(userAgentList); i++ {
		// 随机获取一个userAgent
		userAgent := GetUserAgent()
		item := UserAgentItem{
			UserAgent: userAgent,
			//InCookie:  "", // GetInCookie(userAgent) 刚开始先不初始化
		}
		userAgentResponse = append(userAgentResponse, &item)
	}
	return &UserAgent{
		UserAgentItemList: userAgentResponse,
	}, apiKeyList
}

// 随机获取一个cookie及其对应的userAgent
func (this *UserAgent) Get() (*UserAgentItem) {
	if len(this.UserAgentItemList) <= 0 {
		return nil
	}
	item := this.UserAgentItemList[rand.Intn(len(this.UserAgentItemList))]
	return item
}

// 获取一个随机的userAgent
func GetUserAgent() string {
	return userAgentList[rand.Intn(len(userAgentList))]
}

// 通过UA获取到cookie
func GetInCookie(userAgent string) (string) {
	capabilities := agouti.NewCapabilities()
	capabilities["phantomjs.page.settings.userAgent"] = userAgent
	capabilitiesOption := agouti.Desired(capabilities)
	driver := agouti.PhantomJS(capabilitiesOption)
	err := driver.Start()
	if err != nil {
		return ""
	}
	defer driver.Stop()
	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		return ""
	}
	err = page.Navigate("http://www.pubgtracker.com/")
	if err != nil {
		return ""
	}
	redirectURL := "https://pubgtracker.com/"
	// 判断是否请求成功
	for true {
		sleepTime := 0
		url, err := page.URL()
		if err != nil {
			return ""
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
		return ""
	}
	cookies, err := page.GetCookies()
	if err != nil {
		return ""
	}
	for _, item := range cookies {
		if item.Name == "cf_clearance" {
			return "cf_clearance=" + item.Value
		}
	}
	return ""
}
