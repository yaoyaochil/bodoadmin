package util

import "time"

// GetYesterday 获取昨日日期
func GetYesterday() string {
	now := time.Now()
	yes := now.AddDate(0, 0, -1)
	return yes.Format("2006-01-02")
}

// GetYearFirstDay 获取今年第一天
func GetYearFirstDay() string {
	now := time.Now()
	return now.Format("2006-01-02")[:4] + "-01-01"
}

// GetToday 获取今日日期
func GetToday() string {
	now := time.Now()
	return now.Format("2006-01-02")
}
