package lol

import (
	"testing"
	"net/http/cookiejar"
)

func TestLol(t *testing.T) {
	jar, _ := cookiejar.New(nil)
	lol := LOL{
		AppId:     "21000501",
		QQ:        "305665638",
		Password:  "qqqq1111....",
		Jar:       jar,
		UserAgent: "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	}
	lol.run()
}
