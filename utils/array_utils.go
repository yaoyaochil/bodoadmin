package utils

// InArray 判断是否在数组中
func InArray(str string, arr []string) bool {
	for _, v := range arr {
		if str == v {
			return true
		}
	}
	return false
}
