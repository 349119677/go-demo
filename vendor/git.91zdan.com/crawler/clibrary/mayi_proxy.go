package clibrary

import (
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

var DefaultMayiProxy = MaYiProxy{AppKey: "276536", AppSecret: "adf73190d7add1015967b63016690a9a", ProxyServer: "http://s2.proxy.mayidaili.com:8123"}

type MaYiProxy struct {
	ProxyServer string
	AppKey      string
	AppSecret   string
}

func (m MaYiProxy) client() http.Client {
	proxyUrl, _ := url.Parse(m.ProxyServer)
	return http.Client{Timeout: time.Second * 30, Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
}

// 注意 authHeader 必须在每次请求时都重新计算，要不然会因为时间误差而认证失败
// Auth Header 文档
// http://www.mayidaili.com/dynamic#p3
// APPKEY认证,在请求头中加入"Proxy-Authorization"字段进行验证，适合需要对请求过程进行控制的用户
// 如：Proxy-Authorization:MYH-AUTH-MD5 sign=83EAA31C85F9DDE368F5266A90087488&app_key=12345678&timestamp=2015-12-02 15:32:58。
func (m MaYiProxy) authHeader() string {
	params := map[string]string{}
	params["app_key"] = m.AppKey
	params["timestamp"] = ToPrimaryTimeLayout(time.Now())

	// 将表格中除 sign 外的所有使用到的参数按照参数名的字母先后顺序排序
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	// 将排序后的参数表进行字符串连接
	str := m.AppSecret
	for _, key := range keys {
		str += key + params[key]
	}
	str += m.AppSecret
	// 对符串进行MD5计算，并转换为全大写形式后即获得签名串
	params["sign"] = strings.ToUpper(MD5(str))

	header := "MYH-AUTH-MD5 "
	for key := range params {
		header += key + "=" + params[key] + "&"
	}
	return header[0 : len(header)-1]
}

func (m MaYiProxy) proxyAuthorization(req *http.Request) {
	req.Header.Add("Proxy-Authorization", m.authHeader())
}
