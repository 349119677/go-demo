package mail

import (
	"testing"
	"fmt"
)

func Test(t *testing.T) {
	response := GetList("imap.qq.com:993", "1436863821@qq.com", "mrosvtjgojhdgddj")
	// 测试输出
	for _, value := range response {
		fmt.Println("发件人：", value.From)
		fmt.Println("主题：", value.Subject)
		fmt.Println("IP：", value.Ip)
		fmt.Println("IsOriginal：", value.IsOriginal)
		fmt.Println("发件内容：", value.Content)
	}
}
