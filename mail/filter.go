package mail

// 是否是原始邮件 获取到就说明是转发的邮件
func IsOriginal(ip string) bool {
	if ip == "" {
		return true
	} else {
		return false
	}
}

// 是否需要解析邮件内容 是否是用户邮件
func IsBankMail() bool {
	return false
}
