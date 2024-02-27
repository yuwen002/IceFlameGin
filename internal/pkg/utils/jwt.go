package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type CustomClaimsInput struct {
	Id            uint32                 // 用户ID
	UserInfo      map[string]interface{} // 用户信息
	Expire        int64                  // 过期时间
	Before        int64                  // 生效时间
	Issuer        string                 // 签发人
	Subject       string                 // 签发主题
	SigningMethod jwt.SigningMethod
}

// CustomClaims
// @Description: 自定义输入信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-27 15:31:38
type CustomClaims struct {
	Id       uint32                 `json:"user_id"`
	UserInfo map[string]interface{} `json:"user_info"`
	jwt.RegisteredClaims
}

// JwtClaims
// @Description: 用户签名验证
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-27 15:40:42
type JwtClaims struct {
	SecretKey string
}

// VerifySecretKey
//
// @Title VerifySecretKey
// @Description: 验证SecretKey
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-27 15:41:11
// @receiver jc
// @return error
func (jc *JwtClaims) VerifySecretKey() error {
	if jc.SecretKey == "" {
		err := errors.New("密钥信息错误")
		return err
	}
	return nil
}

// CreateToken
//
// @Title CreateToken
// @Description:
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-27 15:42:42
// @receiver jc
// @param in
// @return string
// @return error
func (jc *JwtClaims) CreateToken(in CustomClaimsInput) (string, error) {
	// 验证秘钥
	err := jc.VerifySecretKey()
	if err != nil {
		return "", err
	}

	if in.Id == 0 {
		err := errors.New("ID信息错误")
		return "", err
	}

	// 判断过期时间
	var expiresAt *jwt.NumericDate
	if in.Expire > 0 {
		expiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(in.Expire) * time.Second))
	} else {
		expiresAt = jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour * time.Duration(1)))
	}

	// 签发时间
	now := jwt.NewNumericDate(time.Now())

	// 判断生效时间
	var notBefore *jwt.NumericDate
	if in.Before > 0 {
		notBefore = jwt.NewNumericDate(time.Now().Add(time.Duration(in.Before) * 24 * time.Hour))
	} else {
		notBefore = now
	}

	// 判断加密方式
	if in.SigningMethod == nil {
		in.SigningMethod = jwt.SigningMethodHS256
	}

	// 初始化加密数据
	customClaims := CustomClaims{
		Id:       in.Id,
		UserInfo: in.UserInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    in.Issuer,
			Subject:   in.Subject,
			ExpiresAt: expiresAt,
			NotBefore: notBefore,
			IssuedAt:  now,
		},
	}
	// 生成token
	claims := jwt.NewWithClaims(in.SigningMethod, customClaims)

	return claims.SignedString([]byte(jc.SecretKey))
}

// ParseToken
//
// @Title ParseToken
// @Description: 解析Token
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-27 15:44:28
// @receiver jc
// @param token
// @return *CustomClaims
// @return error
func (jc *JwtClaims) ParseToken(token string) (*CustomClaims, error) {
	// 验证秘钥
	err := jc.VerifySecretKey()
	if err != nil {
		return nil, err
	}

	// 解析token
	withClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jc.SecretKey), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("token不正确")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token已经过期")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token未激活")
			} else {
				return nil, errors.New("token错误")
			}
		}
	}

	if claims, ok := withClaims.Claims.(*CustomClaims); ok && withClaims.Valid {
		return claims, nil
	}

	return nil, errors.New("无法处理此token")
}

// CreateToken
//
// @Title CreateToken
// @Description: 创建token
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-27 15:51:53
// @param in
// @param secretKey
// @return string
// @return error
func CreateToken(in CustomClaimsInput, secretKey string) (string, error) {
	jwt := JwtClaims{SecretKey: secretKey}
	return jwt.CreateToken(in)
}

// ParseToken
//
// @Title ParseToken
// @Description: 解析Token
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-27 15:52:22
// @param token
// @param secretKey
// @return *CustomClaims
// @return error
func ParseToken(token string, secretKey string) (*CustomClaims, error) {
	jwt := JwtClaims{SecretKey: secretKey}
	return jwt.ParseToken(token)
}
