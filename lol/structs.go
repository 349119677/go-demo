package lol

import "net/http/cookiejar"

type LOL struct {
	AppId             string         // 项目标识
	UserAgent         string         // 浏览器
	Jar               *cookiejar.Jar // cookie
	QQ                string         // QQ号码
	Password          string         // 未加密的密码
	EncPassword       string         // 加密后的密码
	LoginSig          string         // 获取验证码和登录时 需要
	PtVcodeV1         string         // 是否需要验证码的标志， 提交登录时的参数 1需要 2不需要
	PtVerifySessionV1 string         // 获取验证码和登录时 需要
	CapCd             string         // 获取验证码时 需要
	Code              string         // 验证码
}
