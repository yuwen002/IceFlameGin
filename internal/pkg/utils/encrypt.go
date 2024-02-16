package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// Md5
//
// @Title Md5加密函数
// @Description: Md5加密函数
// @Author liuxingyu
// @Date 2024-01-21 19:05:43
// @param data
// @return string
func Md5(data string) string {
	// 创建一个新的 MD5 对象
	hash := md5.New()
	// 将数据写入哈希器
	hash.Write([]byte(data))
	// 计算哈希值并转换为十六进制字符串
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	return hashString
}

// HashWithSHA1
//
// @Title HashWithSHA1
// @Description: 使用 SHA-1 进行哈希计算
// @Author liuxingyu
// @Date 2024-02-16 18:30:10
// @param data
// @return string
func HashWithSHA1(data string) string {
	hash := sha1.New()
	hash.Write([]byte(data))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

// HashWithSHA256
//
// @Title HashWithSHA256
// @Description: 使用 SHA-256 进行哈希计算
// @Author liuxingyu
// @Date 2024-02-16 18:20:15
// @param data
// @return string
func HashWithSHA256(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

// HashSHA512
//
// @Title HashSHA512
// @Description: 使用 SHA-512 进行哈希计算
// @Author liuxingyu
// @Date 2024-02-16 18:21:35
// @param data
// @return string
func HashSHA512(data string) string {
	hash := sha512.New()
	hash.Write([]byte(data))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
