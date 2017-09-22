package mail

import (
	"github.com/emersion/go-message/mail"
	"io"
	"io/ioutil"
)

// 解析一封邮件 只要有一步解析失败就返回nil
func CleanOne(oneLiteral Literal) *ResponseMessage {
	// 读取邮件流
	mailMessage, err := mail.CreateReader(oneLiteral)
	if err != nil {
		return nil
	}
	// 返回体
	var response = new(ResponseMessage)
	header := mailMessage.Header


	// 发件人
	from, err := header.AddressList("From")
	if err != nil {
		return nil
	}
	response.From = from[0].Address

	// 主题
	subject, err := header.Subject()
	if err != nil {
		return nil
	}
	response.Subject = subject

	// 原始ip
	ip := header.Get("X-Originating-IP")
	response.Ip = ip

	// 解析邮件body
	for {
		thisPart, err := mailMessage.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil
		}
		switch thisPart.Header.(type) {
		case mail.TextHeader:
			content, err := ioutil.ReadAll(thisPart.Body)
			if err != nil {
				return nil
			}
			response.Content = string(content)
		}
	}
	return response
}

// 解析得到邮件集合
func GetResolvedMailList(literalList []Literal) []*ResponseMessage {
	// 返回体
	response := []*ResponseMessage{}
	for _, value := range literalList {
		resolvedOneMail := CleanOne(value)
		// 这边可能要根据具体邮件跳过某些邮件，定位信用卡邮件
		response = append(response, resolvedOneMail)
	}
	return response
}
