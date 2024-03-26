package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net"
)

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

// IPToLong
//
// @Title IPToLong
// @Description:  IP 地址转换为一个 uint64 类型
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-27 00:01:26
// @param ipString
// @return uint64
// @return error
func IPToLong(ipString string) (uint64, error) {
	ip := net.ParseIP(ipString)
	if ip == nil {
		return 0, fmt.Errorf("invalid IP address: %s", ipString)
	}

	ipInt := uint64(0)
	if ip4 := ip.To4(); ip4 != nil {
		for i := 0; i < net.IPv4len; i++ {
			ipInt += uint64(ip4[i]) << ((net.IPv4len - 1 - i) * 8)
		}
	} else {
		ip16 := ip.To16()
		for i := 0; i < net.IPv6len; i++ {
			ipInt += uint64(ip16[i]) << ((net.IPv6len - 1 - i) * 8)
		}
	}

	return ipInt, nil
}

// LongToIP
//
// @Title LongToIP
// @Description: uint64 长整型表示转换回原始的 net.IP 类型的 IP 地址
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-27 00:02:34
// @param ipInt
// @return net.IP
func LongToIP(ipInt uint64) net.IP {
	ip := make(net.IP, net.IPv6len)
	for i := net.IPv6len - 1; i >= 0; i-- {
		ip[i] = byte(ipInt & 0xFF)
		ipInt >>= 8
	}

	return ip
}
