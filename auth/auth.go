package auth

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

// Encrypt encrypts the plain text with bcrypt.
func GeneratePassword(source string) (string, string) {
	salt := GetRandomSalt()
	password := MD5(source + salt)
	return password, salt
}

// Compare compares the encrypted text with the plain text if it's the same.
func ComparePassword(hashedPassword, password, salt string) bool {
	return hashedPassword == MD5(password+salt)
}

// 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// return len=8  salt
func GetRandomSalt() string {
	return GetRandomString(8)
}

//生成随机字符串
func GetRandomString(size int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
