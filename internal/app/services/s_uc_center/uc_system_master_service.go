package services

import (
	"encoding/hex"
	"fmt"
	config2 "ice_flame_gin/config"
	dto "ice_flame_gin/internal/app/dto/d_uc_center"
	repositories "ice_flame_gin/internal/app/repositories/r_uc_center"
	"ice_flame_gin/internal/pkg/utils"
	"ice_flame_gin/internal/system"
	"time"
)

//	sUcSystemMaster
//	@Description:  表示 UcSystemMaster 服务
//
// @Author liuxingyu
// @Date 2024-02-08 21:50:16
type sUcSystemMaster struct {
	prefix    string
	aes128Key string
}

// NewUcSystemMasterService
//
// @Title NewUcSystemMasterService
// @Description: 创建一个新的 UcSystemMaster 服务实例
// @Author liuxingyu
// @Date 2024-02-08 21:12:49
// @return *sUcSystemMaster 返回一个指向 UcSystemMaster 服务实例的指针
func NewUcSystemMasterService() *sUcSystemMaster {
	return &sUcSystemMaster{
		prefix:    "SA_",
		aes128Key: "IceFlame-G84bWx5",
	}
}

// LoginTelPassword
//
// @Title LoginTelPassword
// @Description: 管理员电话密码登入
// @Author liuxingyu
// @Date 2024-02-08 21:46:35
// @receiver s
// @param in
// @return *system.ApiResponse
func (s *sUcSystemMaster) LoginTelPassword(in dto.LoginTelPasswordInput) *system.SysResponse {
	tel := s.prefix + in.Tel
	out, err := repositories.NewUcAccountRepository().GetAccountByTel(tel)

	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	if out == nil {
		return &system.SysResponse{
			Code:    1,
			Message: "用户不存在",
			Data:    nil,
		}
	}

	if utils.PasswordVerify(in.Password, out.PasswordHash) == false {
		return &system.SysResponse{
			Code:    1,
			Message: "密码错误",
			Data:    nil,
		}
	}

	return &system.SysResponse{
		Code:    0,
		Message: "success",
		Data:    out,
	}
}

// Register
//
// @Title Register
// @Description: 管理员用户注册
// @Author liuxingyu
// @Date 2024-02-11 00:04:28
// @receiver s
// @param in
// @return *system.SysResponse
func (s *sUcSystemMaster) Register(in dto.RegisterSystemMasterInput) *system.SysResponse {
	// 查询电话号是否被注册
	tel := "SA_" + in.Tel
	telAccount, err := repositories.NewUcAccountRepository().GetAccountByTel(tel)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}
	if telAccount != nil {
		return &system.SysResponse{
			Code:    1,
			Message: "用户已存在",
			Data:    nil,
		}
	}

	//  查询email是否被使用
	emailSystemMaster, err := repositories.NewUcSystemMasterRepository().GetByEmail(in.Email)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	if emailSystemMaster != nil {
		return &system.SysResponse{
			Code:    1,
			Message: "用户Email已存在",
			Data:    nil,
		}
	}

	// 密码加密
	in.Password, err = utils.PasswordHash(in.Password)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	err = repositories.NewUcSystemMasterRepository().Insert(in)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	return &system.SysResponse{
		Code:    0,
		Message: "数据写入成功",
		Data:    nil,
	}
}

// ForgotPassword
//
// @Title ForgotPassword
// @Description: 忘记密码，并处理发送邮件
// @Author liuxingyu
// @Date 2024-02-15 22:44:36
// @receiver s
// @param in
// @return *system.SysResponse
func (s *sUcSystemMaster) ForgotPassword(in dto.ForgotPasswordSystemMasterInput) *system.SysResponse {
	// 查询Email信息
	systemMaster, err := repositories.NewUcSystemMasterRepository().GetByEmail(in.Email)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	if systemMaster == nil {
		return &system.SysResponse{
			Code:    1,
			Message: "用户Email不存在",
			Data:    nil,
		}
	}

	// 当前时间戳
	timestamp := time.Now().Unix()
	// 构建明文
	plaintext := in.Email + "|" + fmt.Sprint(timestamp)
	// 加密
	ciphertext, err := utils.EncryptAES128([]byte(s.aes128Key), []byte(plaintext))
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: "加密失败:" + err.Error(),
			Data:    nil,
		}
	}

	// 将密文转换为十六进制字符串进行打印
	ciphertextHex := hex.EncodeToString(ciphertext)
	// 邮件内容信息 @todo 网址需要调整
	body := "<p>这是邮件内容。点击 <a href=\"https://www.example.com/reset-password?token=" + ciphertextHex + "\">此处</a> 重置您的密码。</p>"
	// 配置
	email := config2.GlobalConfig.Email["smtp"]
	config := utils.GoEmailConfig{
		SMTPHost:     email.Host,
		SMTPPort:     email.Port,
		SMTPUsername: email.Username,
		SMTPPassword: email.Password,
		From:         email.Username,
		To:           in.Email,
		Subject:      "Ice Flame 后台用户密码找回",
		Body:         body,
		ContentType:  "text/html",
	}

	err = config.SendMail()
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: "邮件发送不成功",
			Data:    nil,
		}
	}

	return &system.SysResponse{
		Code:    0,
		Message: "邮件发送成功",
		Data:    nil,
	}
}

func (s *sUcSystemMaster) PasswordRecovery() {
	// 解密
	decryptedText, err := utils.DecryptAES128([]byte(s.aes128Key), ciphertext)
	if err != nil {
		fmt.Println("解密失败:", err)
		return
	}

	fmt.Println("解密后的明文:", string(decryptedText))
}
