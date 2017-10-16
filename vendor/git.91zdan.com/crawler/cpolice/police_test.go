package cpolice

import (
	"testing"
	"fmt"
	"time"
	"strconv"
)

func TestPolice_PressIn(t *testing.T) {
	m := Police{
		AlertCount:   10,
		MailUserName: "notice@91zdan.com",
		MailPassword: "cEqaXZ6u",
		MailHost:     "smtp.exmail.qq.com:25",
		MailTo:       "zhangrenzhan@91zdan.com",
	}
	count := 0
	for i := 0; i < 20; i++ {
		go func() {
			count++
			if count%2 == 0 {
				m.PressIn("abc", fmt.Errorf("error message %d", count), strconv.Itoa(count))
			} else {
				m.PressIn("abc", nil, strconv.Itoa(count))
			}
		}()
		//go fmt.Printf("index %d \n",i)
	}
	time.Sleep(time.Second * 2)
}
