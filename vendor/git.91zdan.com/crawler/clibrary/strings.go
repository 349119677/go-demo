package clibrary

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
)

//字符串截取 从doc中找出第一个 start字符开始 到end结束( 不包含start)
func SplitDoc(doc string, start string, end string) string {
	n := strings.Index(doc, start)
	if n == -1 {
		n = 0
	} else {
		n += len(start)
	}
	doc = string([]byte(doc)[n:])
	m := strings.Index(doc, end)
	if m == -1 {
		m = len(doc)
	}
	doc = string([]byte(doc)[:m])
	return doc
}

//字符串去空格和换行
func TrimDoc(doc string) string {
	// 去除空格
	doc = strings.Replace(doc, " ", "", -1)
	// 去除换行符
	doc = strings.Replace(doc, "\n", "", -1)
	doc = strings.Replace(doc, "\t", "", -1)
	return doc
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
