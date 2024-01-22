package utils

import "time"

// GetCurrentTime 获取当前时间
func GetCurrentTime() int64 {
	// 获取当前时间戳 精确到秒
	return time.Now().Unix()
}
