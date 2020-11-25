package util

import "time"

const (
	layout        = "2006-01-02 15:04:05"
	layoutYmd     = "20060102"
	chinaTimeZone = "Asia/Shanghai"
)

func GetCurrentTime() string {
	loc, err := time.LoadLocation(chinaTimeZone)
	if err != nil {
		panic(err)

	}
	return time.Now().In(loc).Format(layout)
}

func GetCurrentDay() string {
	loc, err := time.LoadLocation(chinaTimeZone)
	if err != nil {
		panic(err)

	}
	return time.Now().In(loc).Format(layoutYmd)
}
