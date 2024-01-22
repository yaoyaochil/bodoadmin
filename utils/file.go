package utils

// TrimSuffix 去除文件后缀
func TrimSuffix(fileName string) string {
	return fileName[:len(fileName)-len(GetFileSuffix(fileName))]
}

// GetFileSuffix 获取文件后缀
func GetFileSuffix(fileName string) string {
	return fileName[len(fileName)-4:]
}
