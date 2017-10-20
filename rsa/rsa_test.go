package myRsa

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
-----BEGIN PRIVATE KEY-----
MIIEpAIBAAKCAQEA03CEqnRzYWtI5PHgIaMJnMyZSTaDbBoxbrPJt0qpjncB9J+b
JlC3xvOOOaEdEpQWOrgPttWxyps60scHrjhtWckIo9Re/FGOMrL7f+cLLxPhGzMc
/o7fPd6NhPXxUKK/Iu07ISJsu+D+1Fl6dHY7SFT7UQX/iySYRWep6Vcrk0zQYgO1
lAeUPDi60HeR/UlIFamSyrksvFlZu4p2hjY5un7aiaC5yMQma6Pmnt1B8kuWEyzO
txIVH7Ryb5ahdEn2IYbG+rgrhKn7JSmY08kfIYHdA63TMobOEcGRU3ftdIeBlIxi
mtm2uvIcgmJI5peOihq/fctovs6J0EvbVho7iQIDAQABAoIBAG/9cMc31sUKphld
Y3FtgXHjjG0Sypk/Zl9UKstCaHxk3ExNHUg2CKD/75zmkRd+CCghxXD5zqmZfpaV
hKPqj0C3EjR0D7tlFwQTeNJN9caBqQFXGUxbMDL85cg+3AnxqXs+W8s5CI6apV6j
5hA5bzzohhRsMOXqBBz+wygsW5Awb8Jy+tjIeGXCQeAA7E4RLIP93MRc/cstOh0G
i82B2nzPNcdlPjr2t/rXxKFfIg0mJKCi+Jsu0owWp0EaUX5ZGzY0kZLwmlSgClhh
3/M8WEU99gLhE1DkXXqRdxb0CoX/5DyEXxlZ11QAqLhcDrjIjd8TA+Kssxv2yQ5C
kzSET+ECgYEA08UHXDfHpU2Nbl1v0ah65JgepGOizC24oGd9W8e6QIQzHN8CVcT+
yU0qhXD3y2o8CpzQO02er8WluHkgTZLu0xxggoXdtN6F5jyPW/Ig4nC5fCNER4HD
MNzQYpRwvJAouaDaygKfZ42RCLS33pAfHaS3u/ooPm57+Pw3/+Qz3L8CgYEA/5nW
qwwahM4FzpYONaFY9cWQWohoyLDEM+nJqWApMmSKBoqeWWSry44agwloH0xZ/Ugq
8qayXYrXa0JD7DZ5NxXJ/k8yCg0E4N4g33J4Tau4x0tEgdKw+2/zHgvSnH4bF4KW
MYnxYWyRbB5XliOPG1lfsBsGxfNEEuK94jyRUbcCgYAUWpyhfW7/8VtFn09vDE17
iS9wx27PRxm2uRuwZZr+NnafMMQNsFh4yqTqnHgEohpRPs840/YLmOibuXCkZBn2
SoTwJqOvja+6+FkjEjuc6A674rveT5eOK7sProrDZOu1I8PDpMrjbhThjFUb3ChR
dhJ+y8VwcrgRr1RGkQ13CwKBgQDR+EFln/rE8C3LCG/B1Lqso1AzXu2dN8Dl2H7m
Ge2dzQOp3gO48b9C54iV7otPcrxWGgvV6+SIfX77SKNdj15CRy75L3uu0AUa6L7e
cD2tqIyRgx5S+46R7uQr4ZBxKBL/XDIfne7hlntb8w0GdE2iLOgzVfBZer6IOSW9
jP3fvwKBgQDBPvt9E3AuHv6a2/aZmIggo7muPbolZPqM4ZB1+w4wyb5UsZuMALdM
SEM0PUzcR6yYDoIkZSq3HQuI1BY6qS3a8COniKkT2mFEpSfri+cErb3h7gwYhUkx
QkuJJVGLNNucKccquH/+NXTWaD9Isd4kmvY33YChb3n6l8SLTtZWNA==
-----END PRIVATE KEY-----
`
// 初始化设置公钥和私钥
func init() {
	if err := RSA.SetPublicKey(Pubkey); err != nil {
		log.Fatalln(`set public key :`, err)
	}
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
