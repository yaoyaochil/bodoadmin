package utils

import (
	"crypto/rand"
	"strings"
)

func GenerateCustomRandomPassword(prefix string, length int) (string, error) {
	charset := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsetLength := len(charset)

	// 生成随机字符
	randomChars := make([]byte, length-len(prefix))
	_, err := rand.Read(randomChars)
	if err != nil {
		return "", err
	}

	// 构建密码
	password := strings.Builder{}
	password.WriteString(prefix)
	for _, c := range randomChars {
		password.WriteByte(charset[int(c)%charsetLength])
	}

	return password.String(), nil
}
