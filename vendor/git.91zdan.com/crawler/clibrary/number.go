package clibrary

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	regFloat64 = regexp.MustCompile("[1-9]\\d*.\\d*|0.\\d*[1-9]\\d*")
	regInt     = regexp.MustCompile("[1-9]\\d*")
	regDate    = regexp.MustCompile("\\d{4}-\\d{2}-\\d{2}")
)

// 提取字符串中的正浮点数
func FindFloat64(str string) float64 {
	money := regFloat64.FindString(str)
	if money == "" {
		return 0
	}
	money = strings.Replace(money, ",", "", -1) // 去掉金额中的逗号
	float, err := strconv.ParseFloat(money, 64)
	if err != nil {
		log.Warning("Find float64 [%s] error %s", str, err.Error())
	}
	return float
}

// 提取字符串中的正整数
func FindUInt8(str string) uint8 {
	integer, err := strconv.ParseInt(regInt.FindString(str), 10, 16)
	if err != nil {
		log.Warning("Find uint8 [%s] error: %s", str, err.Error())
	}
	return uint8(integer)
}

// 单位转换，分 -> 元
func FenToYuan(money int) float64 {
	return float64(money / 100)
}
