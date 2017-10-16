package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
)

func main() {

	url := "https://w.mail.qq.com/cgi-bin/mail_list?fromsidebar=1&sid=wI1y7iOnJYFhAVQNUcfT5vNw%2C4%2Cc9bRKC-yz4MY.&folderid=1&page=0&pagesize=10&sorttype=time&t=mail_list&loc=today%2C%2C%2C151&version=html"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("host", "w.mail.qq.com")
	req.Header.Add("connection", "keep-alive")
	req.Header.Add("cache-control", "max-age=0")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.101 Safari/537.36")
	req.Header.Add("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Add("dnt", "1")
	req.Header.Add("referer", "https://w.mail.qq.com/cgi-bin/today?sid=wI1y7iOnJYFhAVQNXcLT5vNw,4,c9bRKC-yz4MY.&first=1")

	req.Header.Add("accept-language", "zh-CN,zh;q=0.8Cookie:pgv_pvi=5357012992; RK=GuGHBSXaF0; luin=o1436863821; lskey=000100003d6cefcf1114d71eb87e3f7e71608d66ef8b246d5448bacaf954e861cb630b260edf641621a54800; pt2gguin=o1436863821; p_luin=o1436863821; p_lskey=00040000c3b1aaa0e80beebc93b7414126314913cd981e68b4ccac90aaab9ce89d003c5609e581501227cded; pcache=318cad394ccfc7fMTUxMDM5NDAwMg@1436863821@4; mpwd=E4F2DC9AF57CFE52D317FAD1D9851F105C640F1B29D1FCFCE13723CCB754F57D@1436863821@4; mcookie=0&y; qm_flag=0; qqmail_alias=1436863821@qq.com; msid=wI1y7iOnJYFhAVQNXcLT5vNw,4,c9bRKC-yz4MY.; sid=1436863821&dc927c9a483fe4da8723f9d2076868bc,c9bRKC-yz4MY.; qm_username=1436863821; ssl_edition=sail.qq.com; edition=mail.qq.com; username=1436863821&1436863821; qm_sk=1436863821&UcfT5vNw; qm_ssum=1436863821&00625b7741651f776ba9fc774d51d50e; new_mail_num=1436863821&11; device=")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	time.Sleep(time.Second * 1000)
}
