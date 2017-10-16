package cpolice

import (
	"sync"
	"git.91zdan.com/crawler/clibrary"
	"fmt"
	"time"
)

const (
	MAIL_TYPE  = "html"
	MAIL_TITLE = "%s 平台警报"
	MAIL_BODY  = `<html>
					 <body>
					 <h3>
						平台名称: %s
 					 </h3>
 					 <h3>
 						失败次数: %d
 					 </h3>
 					 <h3>
 					 	成功次数: %d
 					 </h3>
 					 <h3>
 					 	失败率:  %f
 					 </h3>
 					 <h3>
 					 	错误信息:
 					 </h3>
 					 <ul>
 					 %s
 					 </ul>
					 </body>
					 </html>
	`
)

var (
	recordMap = make(map[string]*Record)
	mutex     sync.Mutex
	police    Police
)

//添加错误记录 err为nil是成功
func (p Police) PressIn(tag string, err error, id string) {
	police = p

	//同步锁
	mutex.Lock()
	defer mutex.Unlock()

	record := getRecord(tag)

	if addRecordMessage(record, err, id) {
		postAlert(record)
	}
}

//从map中拿出记录  没有就创建一个新的
func getRecord(tag string) *Record {
	record, ok := recordMap[tag]
	if !ok {
		record = &Record{
			Tag:       tag,
			ShowAlert: false,
			Message:   make([]RecordMessage, 0),
		}
		recordMap[tag] = record
	}
	return record
}

//添加错误信息, 返回是否输出警告
func addRecordMessage(record *Record, err error, id string) bool {
	//如果已经发出过警告就不再收集错误和警告了
	if record.ShowAlert {
		return false
	}
	if err != nil {
		record.Message = append(record.Message, RecordMessage{
			Id:         id,
			RecordTime: time.Now(),
			Message:    err.Error(),
		})
		record.FailureCount++
	} else {
		//record.Message = append(record.Message, "success")
		record.SuccessCount++
	}

	record.FailureRate = float32(record.FailureCount) / float32(record.SuccessCount+record.FailureCount)
	//错误次数大于警报阈值
	record.ShowAlert = int(record.FailureCount) >= police.AlertCount

	return record.ShowAlert
}

//发出警告信息
func postAlert(record *Record) {
	title := fmt.Sprintf(MAIL_TITLE, record.Tag)
	body := fmt.Sprintf(MAIL_BODY, record.Tag, record.FailureCount, record.SuccessCount, record.FailureRate, getRecordMessage(record))
	postMail(title, body)
}

func getRecordMessage(record *Record) string {
	var message string
	for _, err := range record.Message {
		message +=
				"<li>" + err.Message +
				"<ul>" +
				"<li> ID: " + err.Id + "</li>" +
				"<li> Time: " + err.RecordTime.Format("2006-01-02 15:04:05") + "	</li>" +
				"</ul>" +
				"</li>"
	}
	return message
}

func postMail(title string, body string) {
	clibrary.SendEmail(police.MailUserName, police.MailPassword, police.MailHost, police.MailTo, title, body, MAIL_TYPE)
}

type Record struct {
	ShowAlert    bool            //是否发出过警告
	Tag          string          //平台标记
	SuccessCount uint32          //成功次数
	FailureCount uint32          //失败次数
	FailureRate  float32         //失败率
	Message      []RecordMessage //消息记录
}

type RecordMessage struct {
	Id         string
	RecordTime time.Time
	Message    string
}

type Police struct {
	AlertCount   int //警报阈值
	MailUserName string
	MailPassword string
	MailHost     string
	MailTo       string
}
