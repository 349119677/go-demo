package mail

import (
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"log"
)

// 调用的主函数  GetSourceMailList("imap.qq.com:993", "1436863821@qq.com", "mrosvtjgojhdgddj")
func GetList(addr, user, pass string) []*ResponseMessage {
	request, err := FetchMails(addr, user, pass)
	if err != nil {
		log.Fatal(err)
	}
	response := GetResolvedMailList(request)
	return response
}

// 获取邮件源码集合
func FetchMails(addr, user, pass string) ([]Literal, error) {
	client, err := client.DialTLS(addr, nil)
	if err != nil {
		return nil, err
	}
	defer client.Logout()
	if err := client.Login(user, pass); err != nil {
		return nil, err
	}
	// 收件箱
	mbox, err := client.Select("INBOX", true)
	if err != nil {
		return nil, err
	}

	if mbox.Messages == 0 {
		return nil, fmt.Errorf("未获取到邮件")
	}
	// 所有邮件队列
	seqset := new(imap.SeqSet)
	seqset.AddRange(uint32(1), mbox.Messages)

	messages := make(chan *imap.Message, 10)
	done := make(chan error, 1)
	go func() {
		done <- client.Fetch(seqset, []string{"BODY[]"}, messages)
	}()
	// 返回体
	response := []Literal{}
	// 收件箱的所有邮件
	for msg := range messages {
		r := msg.GetBody("BODY[]")
		if r == nil {
			return nil, fmt.Errorf("未获取到邮件内容")
		}
		response = append(response, r)
	}
	return response, nil
}
