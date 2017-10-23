package main

import (
	"fmt"
	"log"
	"encoding/base64"
	"github.com/axgle/mahonia"
)

func main() {
	inputStr := "bmloYW8="
	inputStr= "kYcOZF6itsDGsnPVPjge+yf6NyX3PPcVej+T4atvdomgxTo90zOds3GokX4pi5mACu/1jBSOXJl9pc69wN9/fR7gT9YsoFqV0cJfMQxJkZjbt8o39t9OecFm9l5D5OdzGK1GLMuAmlRddpTsw6m52ViMP94TRVqPfbc1qdKqz+HW2Wh39Op4HI19DHO2N6ZWNLyxCtI8BfZx/r1UBLctQy02UZVAm8WS77cwaooGbJx1aHj3fVGWtKjqnk08yVdsqEMmQPlxc/whlQHQmbxfHge6/9e798Z6QzstUb1aAbSWGByDhILKLRkT/Thca2BfHuMPaOIdn7ec+QCMa47BPQ=="
	// 对上面的编码结果进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(inputStr)
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println("解密后文本为", string(decodeBytes))
	fmt.Println(ConvertToString(string(decodeBytes),"gbk","utf-8"))
}

// 编码转换
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
