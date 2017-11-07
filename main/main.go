package main

import (
	"regexp"
	"strconv"
	"fmt"
	"errors"
)

var (
	regFloat64 = regexp.MustCompile("[1-9]\\d*.\\d*|0.\\d*[1-9]\\d*")
	regInt     = regexp.MustCompile("[1-9]\\d*")
	regDate    = regexp.MustCompile("\\d{4}-\\d{2}-\\d{2}")
	// 匹配数值
	regNumber = regexp.MustCompile("\\d+(\\.{0,1}\\d+){0,1}")
)

func main() {
	a := "10..05554dasd"
	resp, err := FindNumber(a, 1)
	fmt.Println(resp, err)
}

func FindInt(str string) (int, error) {
	integer := regInt.FindString(str)
	resp, err := strconv.Atoi(integer)
	if err != nil {
		return 0, nil
	}
	return resp, nil
}

// 匹配数值 0为第一个 str为源字符串 index为需要匹配的索引 index必须大于等于0
func FindNumber(str string, index uint8) (float64, error) {
	// 查找所有匹配的数值 返回[]string
	number := regNumber.FindAllString(str, -1)
	if int(index)+1 > len(number) {
		return 0, errors.New("not found matched number")
	}
	resp, err := strconv.ParseFloat(number[index], 64)
	if err != nil {
		return 0, err
	}
	return resp, nil
}
