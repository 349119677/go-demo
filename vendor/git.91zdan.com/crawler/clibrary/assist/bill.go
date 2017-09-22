package assist

import (
	"encoding/base64"
	"net/http"
	"time"
)

// 已还金额，单期账单的已还金额是根据还款状态
// 如果已还款，已还金额等于借款总金额，否则等于 0
func HadRepayAmount(status bool, amount float64) float64 {
	if status {
		return amount
	} else {
		return 0
	}
}

// 已还期数，单期账单的已还期数是根据还款状态
// 如果已还款，已还期数等于1，否则等于 0
func HadRepayPeriod(status bool) uint8 {
	if status {
		return 1
	} else {
		return 0
	}
}

// 单期账单总期数
func SingleTotalPeriod() uint8 {
	return 1
}

// 获取Authorization值
func SetAuthorization(userPhone string, token string, req *http.Request, authorization string) {
	str := userPhone + ":" + token
	req.Header.Add(authorization, "Basic "+base64.StdEncoding.EncodeToString([]byte(str)))
}

// 计算 ID，有些贷款平台对于账单没有一个唯一的 ID 返回，我们将用户名和申请时间/还款时间作为这条订单的唯一标识
// MD5(platform + phone + datestr)
func ComputeIdByPhoneAndDateStr(platform string, phone string, datestr string) string {
	return MD5(platform + phone + datestr)
}

// date layout 2006-01-02
// 弃用，建议用 ComputeIdByPhoneAndDatePrimaryLayout 方法
func ComputeIdByPhoneAndDate(platform string, phone string, date time.Time) string {
	return ComputeIdByPhoneAndDateStr(platform, phone, date.Format(DEFAULT_TIME_LAYOUT))
}

// date layout 2006-01-02 15:04:05
func ComputeIdByPhoneAndDatePrimaryLayout(platform string, phone string, date time.Time) string {
	return ComputeIdByPhoneAndDateStr(platform, phone, date.Format(PRIMARY_TIME_LAYOUT))
}
