package logs_revlove_error

import (
	"testing"
	"log"
)

// 初始化私钥
func init() {
	if err := RSA.SetPrivateKey(Pirvatekey); err != nil {
		log.Fatalln(`set private key :`, err)
	}
}

// 时间点之后的所有日志
func TestConnect(t *testing.T) {
	dateStr := "2017-11-05 4:00:00"
	getAllErrorLog(dateStr)
}


