package clibrary

import "testing"

func Test_SendMail(t *testing.T) {

	body := `
		<html>
		<body>
		<h3>
		"Test send to email"
		</h3>
		</body>
		</html>
		`
	err := SendEmail("notice@91zdan.com", "cEqaXZ6u", "smtp.exmail.qq.com:25", "", "测试邮件", body, "html")
	if err != nil {
		println(err.Error())
	} else {
		println("send success")
	}
}
