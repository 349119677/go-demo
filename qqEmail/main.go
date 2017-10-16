package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://w.mail.qq.com/cgi-bin/mail_list?fromsidebar=1&sid=L2LY_jfWCq5NJXApUcfT5vNw%2C8%2CqVk8qUWpqeU1od3lsUXk1dkd4VjRVWHVsdDVMRGhFV1NDY2k5VWFQc1pyY18."

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cookie", "pgv_pvi=1444466688; pgv_si=s9319421952; ptisp=ctc; RK=GuGHBSXaF0; ptcz=9a2068468ed2f44bc0fcd05f601db7cf2d058e08d098454b1b33e1259a3b44cb; pt2gguin=o1436863821; uin=o1436863821; skey=@zYsEyBNXs; p_uin=o1436863821; p_skey=VO*QjjyMhwylQy5vGxV4UXult5LDhEWSCci9UaPsZrc_; pt4_token=0InC20hYSwbfq463UhJfGasMAREFp1ClsfZFC4exQeQ_; wimrefreshrun=0&; qm_domain=https://mail.qq.com; qm_ptsk=1436863821&@zYsEyBNXs; foxacc=1436863821&0; qm_loginfrom=1436863821&wsk; webp=1; CCSHOW=000001; mcookie=0&y; qm_flag=0; qqmail_alias=1436863821@qq.com; msid=L2LY_jfWCq5NJXApXcLT5vNw,4,qVk8qUWpqeU1od3lsUXk1dkd4VjRVWHVsdDVMRGhFV1NDY2k5VWFQc1pyY18.; sid=1436863821&e2f320108c673f7f2c650c4d217ee6d5,qVk8qUWpqeU1od3lsUXk1dkd4VjRVWHVsdDVMRGhFV1NDY2k5VWFQc1pyY18.; qm_username=1436863821; ssl_edition=sail.qq.com; edition=mail.qq.com; username=1436863821&1436863821; device=; qm_sk=1436863821&UcfT5vNw; qm_ssum=1436863821&f4ac78a218fe4e3e9b71e2c4eb78c38b; new_mail_num=1436863821&0")
	req.Header.Add("host", "w.mail.qq.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}