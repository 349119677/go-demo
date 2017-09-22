package proxy

import (
	"net/http"
	"net/url"
	"time"
)

type AbuyunProxy struct {
	AppID       string
	AppSecret   string
	ProxyServer string
}

func Client(proxy AbuyunProxy) http.Client {
	return proxy.proxyClient()
}

func (p AbuyunProxy) proxyClient() http.Client {
	proxyUrl, _ := url.Parse("http://" + p.AppID + ":" + p.AppSecret + "@" + p.ProxyServer)
	return http.Client{Timeout: time.Second * 30, Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
}
