package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"io"
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

// EncryptAES128
//
// @Title EncryptAES128
// @Description: 使用AES-128对数据进行加密
// @Author liuxingyu
// @Date 2024-02-22 00:09:04
// @param key
// @param plaintext
// @return []byte
// @return error
func EncryptAES128(key []byte, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

// DecryptAES128
//
// @Title DecryptAES128
// @Description: 使用AES-128对数据进行解密
// @Author liuxingyu
// @Date 2024-02-22 00:09:58
// @param key
// @param ciphertext
// @return []byte
// @return error
func DecryptAES128(key []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext is too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return ciphertext, nil
}

// EncryptAES256
//
// @Title EncryptAES256
// @Description: 使用AES-256对数据进行加密
// @Author liuxingyu
// @Date 2024-02-22 00:10:16
// @param key
// @param plaintext
// @return []byte
// @return error
func EncryptAES256(key []byte, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

// DecryptAES256
//
// @Title DecryptAES256
// @Description: 使用AES-256对数据进行解密
// @Author liuxingyu
// @Date 2024-02-22 00:10:33
// @param key
// @param ciphertext
// @return []byte
// @return error
func DecryptAES256(key []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext is too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return ciphertext, nil
}
