package clibrary

import (
	"net/http"
	"net/url"
	"time"
)

var DefaultAbuyunProxy = AbuyunProxy{AppID: "HDPMO338J83I7QQD", AppSecret: "1120BDB2BCC4AB11", ProxyServer: "http-dyn.abuyun.com:9020"}

type AbuyunProxy struct {
	AppID       string
	AppSecret   string
	ProxyServer string
}

func (p AbuyunProxy) client() http.Client {
	proxyUrl, _ := url.Parse("http://" + p.AppID + ":" + p.AppSecret + "@" + p.ProxyServer)
	return http.Client{Timeout: time.Second * 30, Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
}
