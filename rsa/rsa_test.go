package rsa

import (
	"log"
	"fmt"
	"encoding/base64"
	"testing"
)

var Pubkey = `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA03CEqnRzYWtI5PHgIaMJ
nMyZSTaDbBoxbrPJt0qpjncB9J+bJlC3xvOOOaEdEpQWOrgPttWxyps60scHrjht
WckIo9Re/FGOMrL7f+cLLxPhGzMc/o7fPd6NhPXxUKK/Iu07ISJsu+D+1Fl6dHY7
SFT7UQX/iySYRWep6Vcrk0zQYgO1lAeUPDi60HeR/UlIFamSyrksvFlZu4p2hjY5
un7aiaC5yMQma6Pmnt1B8kuWEyzOtxIVH7Ryb5ahdEn2IYbG+rgrhKn7JSmY08kf
IYHdA63TMobOEcGRU3ftdIeBlIximtm2uvIcgmJI5peOihq/fctovs6J0EvbVho7
iQIDAQAB
-----END PUBLIC KEY-----
`

var Pirvatekey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEpgIBAAKCAQEArYkPlffsyC0G1wZPhF0pcQFudgBltGfUiW4xlR3IWUbpxZaL
lWDdmVWCJW6NRs4Qqif5I0+zNGbW3ml0a0NCUxtuHlVhGuo80T6D4FqWXEyJh4Qv
ey6k2J9suVahlcZ8SfkXtcojCIwotQ/OlDjvpkfBvCzGEcj1UjiWlso7msSFJ7eQ
BJ8vsue/nO6APpN9ZoacQIF2QyPPrfHOlLN2tNjnNR4YCyDdZdekHALJTAIhN459
Cp7bL60QvwrEx1jXFeb6NhETMEvmG3I+ypQUac4dIrSxQs6pIkgYGU/AFbTcJuo6
XMmigmWuuCfrNEEqTK+uMezo4uK7stF6jArI7wIDAQABAoIBAQCLYqvCKYFmx8PW
sprsFmhS/HNdFLScU0nDmV76BxIFo4/hxSoYsdVMdAI1TrbrSFjaU4Epe7rVPEUa
IFoCTePYHRA2DR4SIFL5Pt1uN1TOjitpTiNVLgH6fRM3Sv4+706lnA4PVm0NUIbh
5/Bl3dWgGcLjApOVdXSWth0+wPFfPWcZlykFCXir4JiC5Q18LHkTopcIRf6pj8hf
9+EUSXFcSUlpIr+tVLgFwZ39eE8XdF/RLaySWeaRUQV7NH5lrpnUQ7Qo6h+hF56p
sXpwEFdQxvc9CdWOZh4eULRHaSRInhF/RuFSOiW9CXDbSzQOVYUEFbk/XvKX4qr6
dblEUFQBAoGBAN7EvLsJklOfwhYCIh6HzBGDjJYv5arQeLwxKU2xlW9MGj/zK06v
rkJu3jLChMIR37HjVebg+4pQaKB54wxop6OoRXRMdJf7aIrTsyQwI7EOChDHtffS
+JlaTnMkBtkBvbew/bdEdCH0MWzK6AcwSxKnaOJtVuUy3dki4IX2jwNfAoGBAMds
K4d/tPGaQr4hRMEa38A728frHixTUVmYy0yxu82lJwx6iTsx8+bviRCOc4pXOBvj
VCSOwvEcAFIFEVdx7R4/48cOqGw2VM6XMP/79lf+Fsz/U+pXHPcD+T4hLttR5Fc0
6RXhk6KGMYtbmwLcOxhMmNKDQyKzfrlWC8hCpTRxAoGBAJwWfacQLdcqOXPySrXZ
sMMUyYzyQppNOEukpqwpoes6B6qHV/DqZ89kn4wxf5P7siGSiL4dJiLj9Gt8+/nV
2qBxZzj12TwkjDu52euhT2Df1W2sgXQD/ux0ukMzGxkyfxwWh6/B5x1jXL497hYb
YH7RjkwZEKnH2moqZed3uv1BAoGBALP2k9m1j46wTSbTUv6Ix/7shcwGC7aogbM/
RceyZO6R6YlDU3BOBfMTmzeLtgzT5NiV76sGAjofsNJTrHCjaP5xXZaXTbpJ1YKR
wjBgKtEBhYX9rLa/Ym4v8c71lxlo740bPHT4moiCXVeLlXmhUNr+1LK0rfM01Flg
o0VXj0qhAoGBAJznWoEleWIcViBnRp8Y76chUJ1wj5dqRPZeIBXNaJMiNfGL/dCA
cKMjiN69g5Tiuxxu/7/Cao4axnqIrG+0vp/St6gBs2XPw5lfp+502voK6zBW7dFS
35ABLg+hy8ICHR8xGcQhgnDqZOc/0T2btHjPz1datOrn9MGdHAEDmFQ4
-----END RSA PRIVATE KEY-----
`
var RSA = &RSASecurity{}
// 初始化设置公钥和私钥
func init() {

	//if err := RSA.SetPublicKey(Pubkey); err != nil {
	//	log.Fatalln(`set public key :`, err)
	//}
	if err := RSA.SetPrivateKey(Pirvatekey); err != nil {
		log.Fatalln(`set private key :`, err)
	}
}

// 公钥加密私钥解密
func Test_加密解密(t *testing.T) {
	sourceStr := "你好"
	pubenctypt, err := RSA.RsaEncrypt([]byte(sourceStr))
	if err != nil {
		fmt.Println(err)
	}
	pridecrypt, err := RSA.RsaDecrypt(pubenctypt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("源字符串为：", sourceStr)
	fmt.Println("解析后字符串为：", string(pridecrypt))
	fmt.Println("加密后字符串为：", base64.StdEncoding.EncodeToString(pubenctypt))
	fmt.Println(string(pridecrypt))

}

// 公钥加密私钥解密
func Test_解密(t *testing.T) {
	// 源base64字符串
	str := "RQjGdVC/6QzyBCN1yUE3nhLYGemo1aKzRtDj3RYiTEs4FyS/xrOchIN8jCSShywd+tRHZhgQ+EWa/uepy0Quas+3R9yQwuq5c6ssiPJB7qAXYMdRi/gqBnP2AaNcOWknbxnA89wGtRH0hM3/vWGjwurZ2DCVWJPHszKucLFOIrOo9QNsz5gssjKiHgYddAYHWuWsJI8BL1LRQ1KMBcu3+KkJbwmPRlsxo0Q1Hv3y0ZWg1gBowBOrUR4CCtWeMRqqRJiT+aiew5+db06UNgNsWapSWiZ03jhG6wMKmVoK8PjLLgWb5uZ/pFBMKSJHG4T//K/WcvmGPUKu3xQxB27J+A=="
	// base解码后的字符数组
	decodeBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Fatalln(err)
	}
	pridecrypt, err := RSA.RsaDecrypt(decodeBytes)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("解析后字符串为：", string(pridecrypt))

}
