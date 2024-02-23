package utils

import (
	"bytes"
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
	if len(key) != 16 {
		return nil, errors.New("AES-128 key size must be 16 bytes")
	}

	return EncryptAES(key, plaintext, 16)
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
	if len(key) != 16 {
		return nil, errors.New("AES-128 key size must be 16 bytes")
	}

	return DecryptAES(key, ciphertext, 16)
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
	if len(key) != 32 {
		return nil, errors.New("AES-256 key size must be 32 bytes")
	}

	return EncryptAES(key, plaintext, 32)
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
	if len(key) != 32 {
		return nil, errors.New("AES-256 key size must be 32 bytes")
	}

	return DecryptAES(key, ciphertext, 32)
}

// pkcs7Pad
//
// @Title pkcs7Pad
// @Description: 使用PKCS#7填充对数据进行填充
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-23 16:02:14
// @param data
// @param blockSize
// @return []byte
func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7Unpad
//
// @Title pkcs7Unpad
// @Description: 使用PKCS#7填充对数据进行去除填充
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-23 16:02:29
// @param data
// @return []byte
// @return error
func pkcs7Unpad(data []byte) ([]byte, error) {
	length := len(data)
	unpadding := int(data[length-1])
	if unpadding > length {
		return nil, errors.New("invalid padding")
	}
	return data[:(length - unpadding)], nil
}

// EncryptAES
//
// @Title EncryptAES
// @Description: 使用AES对数据进行加密
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-23 16:02:53
// @param key
// @param plaintext
// @param keySize
// @return []byte
// @return error
func EncryptAES(key []byte, plaintext []byte, keySize int) ([]byte, error) {
	if len(key) != keySize {
		return nil, errors.New("AES key size is incorrect")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	plaintext = pkcs7Pad(plaintext, blockSize)

	ciphertext := make([]byte, blockSize+len(plaintext))
	iv := ciphertext[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[blockSize:], plaintext)

	return ciphertext, nil
}

// DecryptAES
//
// @Title DecryptAES
// @Description: 使用AES对数据进行解密
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-23 16:03:05
// @param key
// @param ciphertext
// @param keySize
func DecryptAES(key []byte, ciphertext []byte, keySize int) ([]byte, error) {
	if len(key) != keySize {
		return nil, errors.New("AES key size is incorrect")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	if len(ciphertext)%blockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	iv := ciphertext[:blockSize]
	ciphertext = ciphertext[blockSize:]

	plaintext := make([]byte, len(ciphertext))

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)

	plaintext, err = pkcs7Unpad(plaintext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
