// 拍拍贷 爬虫
//
// 拍拍贷只有多期账单
// 关键 Tag：DOM, PHP, Cookies, application/x-www-form-urlencoded
// author zhanghua 2017/07/12
package main

import (
	"cloan/disguise"
	"cloan/kafka"
	"cloan/platforms"
	"cloan/structs"
	"cloan/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/op/go-logging"
	"io/ioutil"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

var (
	log         = logging.MustGetLogger("ppdai")
	crawlParams *structs.CrawlParams
	jar, _      = cookiejar.New(nil)
)

const (
	BASE_URL                      = "http://loan.ppdai.com"                                                  // 拍拍贷服务器基础地址
	LOGIN_RESOURCE                = "https://ac.ppdai.com/User/Login"                                        // 登录接口地址
	REPAYMENT_LIST_RESOURCE       = "http://loan.ppdai.com/account/repaymentlist"                            // 正在还款的接口地址
	LOGIN_SUCCESS_CONTAINS_CHARS  = "<script>location.href='http://www.ppdai.com/firstuserdetail';</script>" // 登录成功后的返回
	LOGIN_SUCCESS_CONTAINS_CHARS2 = "<script>location.href='http://www.ppdai.com/account/borrow';</script>"  // 登录成功后的返回
	LOGIN_SUCCESS_CONTAINS_CHARS3 = "<script>location.href='http://wap.ppdai.com';</script>"
	USERNAME_KEY                  = "UserName"                  // 用户名参数
	PASSWORD_KEY                  = "Password"                  // 密码参数
	REMEMBER_ME_KEY               = "RememberMe"                // 记住我
	SASYNC_KEY                    = "sAsync"                    // sAsync
	TIME_LAYOUT                   = "2006-01-02"                // 时间格式
	STATUS_FINISHED               = "还款"                        // 账单状态已完成标识，有准时还款和逾期还款，所以应该判断状态是否包含还款
	SELECTOR_REPAY_LIST           = ".operations a:first-child" // 选择器：正在还款的列表
	SELECTOR_BILL_OVERVIEW        = ".huankuantab tr"           // 选择器：账单总览
	SELECTOR_BILL_ARRAY           = ".danbaotab tr"             // 选择器：分期账单列表
)

// define PPDai struct to implement Crawler interface
type PPDai struct {
}

// 爬取方法
//
// 主要流程：身份认证(Login), 拉取 HTML Document(Fetch), goquery 解析
// 身份验证使用 Cookies, 用 cookie jar 自动处理
func (p PPDai) Crawl(params structs.CrawlParams) ([]*structs.Bills, error) {
	crawlParams = &params
	err := login(params.Phone, params.Password)
	if err != nil {
		return nil, err
	}

	return bills(params.Phone)
}

// 身份认证 (Login)
//
// 构造 POST Form 请求，如果登录成功会返回 <script>location.href='http://www.ppdai.com/firstuserdetail';</script>
// 身份认证使用 cookies, 使用 cookie jar 会自动设置 cookies
func login(phone string, password string) error {
	// 构造 form 表单，提交登录请求
	params := url.Values{}
	params.Add(USERNAME_KEY, phone)
	params.Add(PASSWORD_KEY, password)
	params.Add(REMEMBER_ME_KEY, "false")
	params.Add(SASYNC_KEY, "true")
	// 使用自己构造的 httpClient 发送请求以自动处理 cookies
	res, err := disguise.PostForm(LOGIN_RESOURCE, params, jar)
	if err != nil {
		return err
	}

	defer res.Body.Close() // 方法运行结束后关闭 Body

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	bodyStr := string(body)

	kafka.SendOriginalBillsMessagep(crawlParams, LOGIN_RESOURCE, platforms.ACTION_LOGIN,
		nil, params, res.Header, string(body))

	// 判断是否登录成功，也就是是否包含成功登录返回的字符
	if !strings.Contains(bodyStr, LOGIN_SUCCESS_CONTAINS_CHARS) &&
		!strings.Contains(bodyStr, LOGIN_SUCCESS_CONTAINS_CHARS2) &&
		!strings.Contains(bodyStr, LOGIN_SUCCESS_CONTAINS_CHARS3) {
		return structs.NewCrawlAuthorizationError()
	}

	return nil
}

// 获得正在还款的账单列表
//
// 根据账单列表的 href 属性，循环调用账单详情页面，解析账单
func bills(phone string) ([]*structs.Bills, error) {
	res, err := disguise.Get(REPAYMENT_LIST_RESOURCE, jar)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return nil, err
	}

	html, _ := doc.Html()
	kafka.SendOriginalBillsMessagep(crawlParams, REPAYMENT_LIST_RESOURCE, platforms.ACTION_FETCH_BILLS,
		nil, nil, res.Header, html)

	billsArray := []*structs.Bills{}
	// 选择正在还款的列表
	doc.Find(SELECTOR_REPAY_LIST).Each(func(i int, a *goquery.Selection) {
		href, exist := a.Attr("href")
		if exist {
			bills, err := billDetail(href, phone)
			if err != nil {
				log.Error("", err)
			} else {
				billsArray = append(billsArray, bills)
			}
		}
	})
	return billsArray, nil
}

// 获取账单详情页
//
// 链接地址等于：BASE_URL + href
func billDetail(href string, phone string) (*structs.Bills, error) {
	myURL := BASE_URL + href
	res, err := disguise.Get(myURL, jar)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return nil, err
	}

	html, _ := doc.Html()
	kafka.SendOriginalBillsMessagep(crawlParams, myURL, platforms.ACTION_FETCH_BILL_DETAIL,
		nil, nil, res.Header, html)

	var bills structs.Bills
	var billArray []structs.PeriodBill
	// 账单总览 table，选择器：.huankuantab tr
	// <table cellpadding="0" cellspacing="0" class="huankuantab">
	// <tr>
	// <th colspan="3">
	// <a href="/list/51913904" target="_blank">手机app用户的第1次闪电借款</a>
	// </th>
	// </tr>
	// <tr>
	// <td>借款金额：&#165;500.00</td>
	// <td>待还本息：&#165;443.80</td>
	// <td>期限：6 个月</td>
	// </tr>
	// </table>
	doc.Find(SELECTOR_BILL_OVERVIEW).Each(func(i int, tr *goquery.Selection) {
		// 只取下标是 1 的 tr
		// <tr>
		//  <td>借款金额：&#165;500.00</td>
		//  <td>待还本息：&#165;443.80</td>
		//  <td>期限：6 个月</td>
		// </tr>
		if i == 1 {
			bills, err = cleanBillsOverview(tr)
		}
	})
	// 找到还款列表，选择器：.danbaotab tr
	// <tr>
	// <td>2017-07-09</td>
	// <td>&#165;0.00/ &#165;88.76</td>
	// <td>&#165;0.00</td>
	// <td>&#165;0.00 </td>
	// <td>
	//	 &#165;0.00 </td>
	// <td>
	// <span class="repaystatus repay-status1">准时还款</span>
	// </td>
	// </tr>
	doc.Find(SELECTOR_BILL_ARRAY).Each(func(i int, tr *goquery.Selection) {
		// 跳过第一个 tr, 第一个 tr 是头部
		// <tr>
		// <th>还款日</th>
		// <th>未还金额/已还金额</th>
		// <th>未还本金</th>
		// <th>未还利息</th>
		// <th>逾期利息</th>
		// <th>状态</th>
		// </tr>
		if i != 0 {
			bill, err := cleanBill(tr, phone)
			if err != nil {
				log.Error("", err)
			} else {
				billArray = append(billArray, bill)
			}
		}
	})
	// 如果账单列表 length 为 0，有可能选择器出问题了，打日志警告
	if len(billArray) == 0 {
		log.Warning("Not found selector %s", SELECTOR_BILL_ARRAY)
	} else {
		compute(&bills, billArray, phone)
	}

	bills.PeriodBills = billArray
	return &bills, nil
}

// 计算 已还期数，是否还款状态，借款时间
// 已还期数 = 分期账单列表中所有已还的数量之和
// 是否还款状态 = 已还期数 == 总期数
// 借款时间 = 第一期账单时间 - 1month
// 已还金额等于所有已还账单金额之和
// 贷款金额等于所有分期账单金额之和
func compute(bills *structs.Bills, billArray []structs.PeriodBill, phone string) {
	var hadRepayPeriod uint8
	var hadRepayAmount float64
	var amount float64
	for i := range billArray {
		bill := billArray[i]
		if i == 0 {
			bills.BorrowTime = bill.RepayDay.AddDate(0, -1, 0)
		}
		if bill.Status {
			hadRepayPeriod++
			hadRepayAmount += bill.RepayAmount
		}
		amount += bill.RepayAmount
	}

	bills.ID = utils.ComputeIdByPhoneAndDate(platforms.KEY_PPDAI, phone, bills.BorrowTime)
	bills.Amount = amount
	bills.HadRepayAmount = hadRepayAmount
	bills.Status = bills.TotalPeriod == hadRepayPeriod
	bills.HadRepayPeriod = hadRepayPeriod
}

// 清洗账单总览
// <tr>
//  <td>借款金额：&#165;500.00</td>
//  <td>待还本息：&#165;443.80</td>
//  <td>期限：6 个月</td>
// </tr>
func cleanBillsOverview(tr *goquery.Selection) (structs.Bills, error) {
	bill := structs.Bills{}
	tr.Find("td").Each(func(i int, td *goquery.Selection) {
		switch i {
		case 0:
			bill.BorrowAmount = utils.FindFloat64(td.Text())
			break
		case 1:
			break
		case 2:
			bill.TotalPeriod = utils.FindUInt8(td.Text())
			break
		default:
			log.Warning("Did not catch data %s", td.Text())
		}
	})
	return bill, nil
}

// 解析一条分期账单
//
// <tr>
// <th>还款日</th>
// <th>未还金额/已还金额</th>
// <th>未还本金</th>
// <th>未还利息</th>
// <th>逾期利息</th>
// <th>状态</th>
// </tr>
// 源代码：
// <tr>
// <td>2017-07-09</td>
// <td>&#165;0.00/ &#165;88.76</td>
// <td>&#165;0.00</td>
// <td>&#165;0.00 </td>
// <td>
//	 &#165;0.00 </td>
// <td>
// <span class="repaystatus repay-status1">准时还款</span>
// </td>
// </tr>
//
// 页面显示：
// 2017-07-09
// ¥0.00/ ¥88.76
// ¥0.00
// ¥0.00
//
// ¥0.00
// 准时还款 / 待还
func cleanBill(tr *goquery.Selection, phone string) (structs.PeriodBill, error) {
	bill := structs.PeriodBill{}
	// 获得所有 td 标签
	tr.Find("td").Each(func(i int, td *goquery.Selection) {
		switch i {
		case 0:
			bill.RepayDay, _ = utils.StringTimeToTime(TIME_LAYOUT, td.Text())
			bill.ID = utils.ComputeIdByPhoneAndDateStr(platforms.KEY_PPDAI, phone, td.Text())
			break
		case 1:
			bill.RepayAmount = cleanRepayAmount(td.Text())
			break
		case 2: // ignore
			break
		case 3: // ignore
			break
		case 4: // ignore
			break
		case 5:
			bill.Status = cleanStatus(td.Text())
			break
		default:
			log.Warning("Did not catch data %s", td.Text())
		}
	})
	return bill, nil
}

// 清洗账单状态
// 准时还款 / 逾期还款 / 待还
func cleanStatus(str string) bool {
	return strings.Contains(str, STATUS_FINISHED)
}

// 清洗每期还款金额
// ¥0.00/ ¥88.76
func cleanRepayAmount(str string) float64 {
	return utils.FindFloat64(str)
}
