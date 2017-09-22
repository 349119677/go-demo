package mail

import "io"

// 邮件源码包装体
type Literal interface {
	io.Reader
	Len() int
}

// 邮件解析后的实体
type ResponseMessage struct {
	From       string
	Subject    string
	Content    string
	IsOriginal bool   // 是否是原始邮件 true代表银行是发件人 false代表是转发的邮件
	Ip         string // X-Originating-IP
}
