package util

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GetRandStringWithCharset 获取指定字符集 下 指定长度的随机字符串
func GetRandStringWithCharset(length int, charset string) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// GetRandString 获取指定长度的随机字符串
func GetRandString(length int) string {
	return GetRandStringWithCharset(length, charset)
}
