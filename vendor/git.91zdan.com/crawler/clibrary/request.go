package clibrary

import (
	"container/list"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const REQUEST_LIMIT_IN_SECOND = 5 // 每秒请求限制

var (
	queue    = list.New()              // 超出请求限制的队列
	reqCount = 0                       // 当前请求总数
	request  = make(chan *httpRequest) // httpclient request 通道
	_        = StartSendHttpRequestJob()
)

// httpclient 请求通道结构体
type httpRequest struct {
	req *http.Request
	jar http.CookieJar
	res chan *http.Response // httpclient response 通道
	err chan error          // httpclient error 通道
}

// httpClient.Do(req)
func Do(req *http.Request, jar http.CookieJar) (*http.Response, error) {
	userAgentNexus6p(req)

	// 建立通道以接收响应和异常
	res := make(chan *http.Response)
	err := make(chan error)

	// 请求放入通道中
	request <- &httpRequest{
		req: req,
		jar: jar,
		res: res,
		err: err,
	}
	return <-res, <-err
}

// httpClient.PostForm(url, data)
func PostForm(url string, data url.Values, jar http.CookieJar) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return Do(req, jar)
}

// httpClient.Get(url)
func Get(url string, jar http.CookieJar) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return Do(req, jar)
}

// 发送 httpclient request 工作
// 限制每秒请求数不能超过 REQUEST_LIMIT_IN_SECOND
// 多出的请求会被代理拒绝
func StartSendHttpRequestJob() int {
	// goroutine
	// 处理实时请求，并检查是否超出每秒请求限制
	// 如果超出限制就将请求放入队列中，下一秒优先执行
	go func() {
		for true {
			r := <-request
			if reqCount < REQUEST_LIMIT_IN_SECOND {
				go do(r)
				reqCount++
			} else {
				queue.PushFront(r)
			}
		}
	}()
	// goroutine
	// 处理超出每秒请求限制的请求
	// 重置 reqCount 变量
	go func() {
		for true {
			for i := 0; i < REQUEST_LIMIT_IN_SECOND; i++ {
				if e := queue.Back(); e != nil {
					queue.Remove(e)
					httpRequest := e.Value.(*httpRequest)
					go do(httpRequest)
				} else {
					reqCount = i // 已处理的请求等于 reqCount
					break
				}
			}
			// sleep 1 second
			time.Sleep(time.Second * 1)
		}
	}()

	return 0
}

// 发送请求，将响应放入结构体的通道
func do(request *httpRequest) {
	Mayi.Jar = request.jar
	DefaultMayiProxy.proxyAuthorization(request.req)
	res, err := Mayi.Do(request.req)
	request.res <- res
	request.err <- err
}
