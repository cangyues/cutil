package td

import (
	"time"
)

var (
	td     time.Time         = time.Now()
	ymd    string            = "2006-01-02"
	ymdhms string            = "2006-01-02 15:04:05"
	month  map[string]string = map[string]string{
		"January":   "01",
		"February":  "02",
		"March":     "03",
		"April":     "04",
		"May":       "05",
		"June":      "06",
		"July":      "07",
		"August":    "08",
		"September": "09",
		"October":   "10",
		"November":  "11",
		"December":  "12",
	}
)

//获取当前年
func GetYear() int {
	return td.Year()
}

//获取当前日
func GetDay() int {
	return td.Day()
}

//获取当前月份
func GetMonth() string {
	return month[td.Month().String()]
}

//获取年月日时分秒
func GetYMDHMS() string {
	return td.Format(ymdhms)
}

//获取年月日
func GetYMD() string {
	return td.Format(ymd)
}

//获取时间戳
func GetUnix() int64 {
	return td.Unix()
}

//日期加减（day +1后一天   -1后一天）
func GetDayPlus(day time.Duration) string {
	return DayOpt(day, "24h")
}

/**
* 日期加减
* duration 加减只
* granularity 加减粒度
 */
func DayOpt(duration time.Duration, granularity string) string {
	d, _ := time.ParseDuration(granularity)
	_td := time.Now()
	_td.Add(d * duration)
	return _td.Format(ymd)
}
