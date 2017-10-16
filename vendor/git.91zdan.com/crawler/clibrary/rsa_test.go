package clibrary

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestRsaDecrypt(t *testing.T) {
	data, err := RsaEncrypt([]byte("cfmmail#sina.com"))
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(data))
	origData, err := RsaDecrypt(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))

}
