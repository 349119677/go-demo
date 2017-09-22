package assist

import "time"

const (
	CHINA_TIME_LAYOUT   = "2006年01月02日"
	DEFAULT_TIME_LAYOUT = "2006-01-02"
	PRIMARY_TIME_LAYOUT = "2006-01-02 15:04:05"
)

func StringTimeToTimeByRFC3339Layout(timeStr string) (time.Time, error) {
	return StringTimeToTime(time.RFC3339, timeStr)
}

func StringTimeToTimeByChinaLayout(timeStr string) (time.Time, error) {
	return StringTimeToTime(CHINA_TIME_LAYOUT, timeStr)
}

// 字符串类型的时间转换成 time.Time 使用默认 layout ""
func StringTimeToTimeByDefaultLayout(timeStr string) (time.Time, error) {
	return StringTimeToTime(DEFAULT_TIME_LAYOUT, timeStr)
}

// 用正则表达式匹配 yyyy-MM-dd 格式的日期
func FindDefaultLayoutTime(str string) (time.Time, error) {
	timeStr := regDate.FindString(str)
	return StringTimeToTimeByDefaultLayout(timeStr)
}

func StringTimeToTimeByPrimaryTimeLayout(timeStr string) (time.Time, error) {
	return StringTimeToTime(PRIMARY_TIME_LAYOUT, timeStr)
}

// 字符串类型的时间转换成 time.Time
func StringTimeToTime(layout string, timeStr string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")
	return time.ParseInLocation(layout, timeStr, loc)
}

// 时间戳转换到 time.Time
func UnixTimestampToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}

//某一天据今天多少天
func ManyDaysDistanceFromToday(OneDay time.Time) int {
	now := time.Now()
	second := now.Unix() - OneDay.Unix()
	oneDaySecond := int64(60 * 60 * 24)
	return int(second/oneDaySecond) + 1
}

// 获取当前时间的日期格式 2006-01-02 15:04:05 yyyy-MM-dd HH:mm:ss
func GetCurrentTimeDateType() string {
	return time.Now().Format(PRIMARY_TIME_LAYOUT)
}
