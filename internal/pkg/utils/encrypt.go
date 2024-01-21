package utils

import (
	"crypto/md5"
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
