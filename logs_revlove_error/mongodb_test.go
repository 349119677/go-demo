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

func TestConnect(t *testing.T) {
	dateStr := "2017-10-26 22:00:00"
	getAllErrorLog(dateStr)
}
