package util

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func GenerateRandomFileName(ext string) string {
	const length = 30
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result strings.Builder
	for i := 0; i < length-len(ext); i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result.WriteByte(charset[num.Int64()])
	}
	return result.String() + ext
}
