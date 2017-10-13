package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://w.mail.qq.com/cgi-bin/mail_list?fromsidebar=1&sid=jMECLOHXiCzzWw5XUcfT5vNw%2C4%2CqYjFqSDlkSEwwTmx4ajlyTHNzWUJaN3Nsek9CNmdPeURrT2Z0UDVHTHluUV8.&folderid=1&page=0&pagesize=100&sorttype=time&t=mail_list&loc=today%2C%2C%2C151&version=html"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("host", "w.mail.qq.com")
	req.Header.Add("connection", "keep-alive")
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.101 Safari/537.36")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Add("dnt", "1")
	req.Header.Add("accept-encoding", "gzip, deflate, br")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.8")
	req.Header.Add("cookie", "mcookie=0&y; pgv_pvi=5357012992; pgv_si=s6904669184; ptisp=ctc; RK=GuGHBSXaF0; luin=o1436863821; lskey=000100003d6cefcf1114d71eb87e3f7e71608d66ef8b246d5448bacaf954e861cb630b260edf641621a54800; pt2gguin=o1436863821; uin=o1436863821; skey=@834Xth3zU; p_uin=o1436863821; p_skey=b1jH9dHL0Nlxj9rLssYBZ7slzOB6gOyDkOftP5GLynQ_; pt4_token=xEDDvLh6TGR0330YRxZjaJthj*zJ42p2Dw0be-hp3u0_; p_luin=o1436863821; p_lskey=00040000c3b1aaa0e80beebc93b7414126314913cd981e68b4ccac90aaab9ce89d003c5609e581501227cded; qm_flag=0; qqmail_alias=1436863821@qq.com; msid=jMECLOHXiCzzWw5XXcLT5vNw,4,qYjFqSDlkSEwwTmx4ajlyTHNzWUJaN3Nsek9CNmdPeURrT2Z0UDVHTHluUV8.; sid=1436863821&1bfc9b4fd279d0f8113d9c17ffa1eec1,qYjFqSDlkSEwwTmx4ajlyTHNzWUJaN3Nsek9CNmdPeURrT2Z0UDVHTHluUV8.; qm_username=1436863821; ssl_edition=sail.qq.com; edition=mail.qq.com; username=1436863821&1436863821; pcache=318cad394ccfc7fMTUxMDM5NDAwMg@1436863821@4; mpwd=E4F2DC9AF57CFE52D317FAD1D9851F105C640F1B29D1FCFCE13723CCB754F57D@1436863821@4; qm_sk=1436863821&UcfT5vNw; qm_ssum=1436863821&c825a795420ba11c527d00b0284f1688; new_mail_num=1436863821&10; device=")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}