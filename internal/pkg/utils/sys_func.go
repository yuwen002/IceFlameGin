package utils

import "golang.org/x/crypto/bcrypt"

// PasswordHash
//
// @Title PasswordHash 创建密码的散列
// @Description: 密码hash
// @Author liuxingyu
// @Date 2024-02-08 21:26:10
// @param password
// @return string
// @return error
func PasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), err
}

// PasswordVerify
//
// @Title PasswordVerify 验证密码是否和散列值匹配
// @Description: 验证密码
// @Author liuxingyu
// @Date 2024-02-08 21:26:55
// @param password
// @param hash
// @return bool
func PasswordVerify(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}

	return true
}
