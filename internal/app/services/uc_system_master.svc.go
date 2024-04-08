package services

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"ice_flame_gin/config"
	"ice_flame_gin/internal/app/constants"
	"ice_flame_gin/internal/app/dto"
	"ice_flame_gin/internal/app/models/model"
	"ice_flame_gin/internal/app/repositories"
	"ice_flame_gin/internal/pkg/utils"
	"ice_flame_gin/internal/system"
	"sync"
	"time"
)

//	sUcSystemMaster
//	@Description:  表示 UcSystemMaster 服务
//
// @Author liuxingyu
// @Date 2024-02-08 21:50:16
type sUcSystemMaster struct {
	prefix     string
	aes128Key  string
	usedTokens map[string]int64
	mu         sync.Mutex
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
		prefix:     "SA_",
		aes128Key:  "IceFlame-G84bWx5",
		usedTokens: make(map[string]int64),
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
func (svc *sUcSystemMaster) LoginTelPassword(in dto.LoginTelPasswordSystemMasterInput) *system.SysResponse {
	tel := svc.prefix + in.Tel
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

	// 查询管理员其他信息
	outExt, err := repositories.NewUcSystemMasterRepository().GetByAccountID(out.ID)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	if outExt == nil {
		return &system.SysResponse{
			Code:    1,
			Message: "Master查询错误",
			Data:    nil,
		}
	}

	token, err := utils.CreateToken(utils.CustomClaimsInput{
		Id: out.ID,
		UserInfo: gin.H{
			"username":      out.Username,
			"name":          outExt.Name,
			"tel":           outExt.Tel,
			"supper_master": outExt.SuperMaster,
		},
		Expire:  60 * 60 * 24,
		Issuer:  "系统管理员:" + outExt.Name,
		Subject: "后台管理",
	}, constants.MasterSecretKey)

	return &system.SysResponse{
		Code:    0,
		Message: "success",
		Data: dto.SystemMasterTokenOutput{
			Token:     token,
			AccountID: out.ID,
		},
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
func (svc *sUcSystemMaster) Register(in dto.RegisterSystemMasterInput) *system.SysResponse {
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
func (svc *sUcSystemMaster) ForgotPassword(in dto.ForgotPasswordInput) *system.SysResponse {
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

	//  加密数据
	result := svc.EncryptToken(in.Email)
	if result.Code == 1 {
		return result
	}
	output, ok := result.Data.(dto.EncryptTokenOutput)
	if !ok {
		// 类型断言失败，处理错误
		return &system.SysResponse{
			Code:    1,
			Message: "数据转换失败",
			Data:    nil,
		}
	}
	ciphertextHex := output.CiphertextHex
	// 邮件内容信息 @todo 网址需要调整
	body := "<p>这是邮件内容。点击 <a href=\"https://www.example.com/reset-password?token=" + ciphertextHex + "\">此处</a> 重置您的密码。</p>"
	// 配置
	email := config.GlobalConfig.Email["smtp"]
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

// EncryptToken
//
// @Title EncryptToken
// @Description: 加密token
// @Author liuxingyu
// @Date 2024-02-25 00:51:29
// @receiver svc
// @param email
// @return *system.SysResponse
func (svc *sUcSystemMaster) EncryptToken(email string) *system.SysResponse {
	// 当前时间戳
	timestamp := time.Now().Unix()
	// 构建明文
	plaintext := email + "|" + fmt.Sprint(timestamp)
	// 加密
	ciphertext, err := utils.EncryptAES128([]byte(svc.aes128Key), []byte(plaintext))
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: "加密失败:" + err.Error(),
			Data:    nil,
		}
	}

	// 将密文转换为十六进制字符串进行打印
	ciphertextHex := hex.EncodeToString(ciphertext)

	return &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data:    dto.EncryptTokenOutput{CiphertextHex: ciphertextHex},
	}
}

// DecryptToken
//
// @Title DecryptToken
// @Description: 解密token
// @Author liuxingyu
// @Date 2024-02-25 00:28:45
// @receiver svc
// @param token
// @return *system.SysResponse
func (svc *sUcSystemMaster) DecryptToken(token string) *system.SysResponse {
	// 将十六进制字符串表示的密文解码为字节切片
	ciphertextBytes, err := hex.DecodeString(token)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: "解码失败:" + err.Error(),
			Data:    nil,
		}
	}

	// 解密
	decryptedTextBytes, err := utils.DecryptAES128([]byte(svc.aes128Key), ciphertextBytes)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: "解密失败:" + err.Error(),
			Data:    nil,
		}
	}

	result := bytes.Split(decryptedTextBytes, []byte("|"))
	email := string(result[0])
	timestamp := string(result[1])
	// 将时间戳转换为整数
	t, err := utils.ToInt64(timestamp)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: "时间戳解析错误:" + err.Error(),
			Data:    nil,
		}
	}

	// 将时间戳转换为时间类型
	createdAt := time.Unix(t, 0)
	// 获取当前时间
	now := time.Now()
	// 计算时间差
	diff := now.Sub(createdAt)
	// 判断时间差是否超过2小时
	if diff > 2*time.Hour {
		return &system.SysResponse{
			Code:    1,
			Message: "密码重置超时:" + err.Error(),
			Data:    nil,
		}
	}

	data := dto.PasswordRecoveryOutput{
		Email:     email,
		Timestamp: t,
	}
	return &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data:    data,
	}
}

// RecoverPassword
//
// @Title RecoverPassword
// @Description: 找回密码
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-22 10:24:39
// @receiver s
// @param token
// @return *system.SysResponse

func (svc *sUcSystemMaster) RecoverPassword(token string, newPassword string) *system.SysResponse {
	// 检查 token 是否已被使用或过期
	_, ok := svc.usedTokens[token]
	if ok {
		return &system.SysResponse{
			Code:    1,
			Message: "token 是否已被使用或过期",
			Data:    nil,
		}
	}

	result := svc.DecryptToken(token)
	if result.Code == 1 {
		return result
	}

	output, ok := result.Data.(dto.PasswordRecoveryOutput)
	if !ok {
		// 类型断言失败，处理错误
		return &system.SysResponse{
			Code:    1,
			Message: "数据转换失败",
			Data:    nil,
		}
	}

	// 密码加密
	newPasswordHash, err := utils.PasswordHash(newPassword)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	//  查询email是否存在
	emailSystemMaster, err := repositories.NewUcSystemMasterRepository().GetByEmail(output.Email)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	if emailSystemMaster == nil {
		return &system.SysResponse{
			Code:    1,
			Message: "用户Email不存在",
			Data:    nil,
		}
	}

	//  更新用户密码
	err = repositories.NewUcAccountRepository().UpdatePasswordById(emailSystemMaster.ID, newPasswordHash)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	svc.mu.Lock()
	defer svc.mu.Unlock()

	// 删除过期的 token
	for key, expiresAt := range svc.usedTokens {
		if expiresAt+2*60*60 <= time.Now().Unix() {
			delete(svc.usedTokens, key)
		}
	}

	// 将 token 存储为生成时间
	svc.usedTokens[token] = output.Timestamp

	return &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data:    nil,
	}
}

// ChangePassword
//
// @Title ChangePassword
// @Description: 修改密码
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-01 23:28:29
// @receiver svc
// @param in
func (svc *sUcSystemMaster) ChangePassword(in dto.ChangePasswordInput) *system.SysResponse {
	// 密码hash
	passwordHash, err := utils.PasswordHash(in.Password)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	// 更新密码
	err = repositories.NewUcAccountRepository().UpdatePasswordById(in.ID, passwordHash)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	return &system.SysResponse{
		Code:    0,
		Message: "Success",
		Data:    nil,
	}
}

// ChangeOwnPassword
//
// @Title ChangeOwnPassword
// @Description: 修改自己的密码
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-01 23:43:54
// @receiver svc
// @param in
// @return *system.SysResponse
func (svc *sUcSystemMaster) ChangeOwnPassword(in dto.ChangeOwnPasswordInput) *system.SysResponse {
	out, err := repositories.NewUcAccountRepository().GetByID(in.ID)
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
			Message: "没有查询到用户信息",
			Data:    nil,
		}
	}

	//  验证密码
	if utils.PasswordVerify(in.OldPassword, out.PasswordHash) == false {
		return &system.SysResponse{
			Code:    1,
			Message: "旧密码错误",
			Data:    nil,
		}
	}

	return svc.ChangePassword(dto.ChangePasswordInput{
		ID:       in.ID,
		Password: in.NewPassword,
	})
}

// GetMasterInfoByAccountId
//
// @Title GetMasterInfoByAccountId
// @Description: 按Account ID 查询master用户信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-06 00:48:17
// @receiver svc
// @param id
// @return *system.SysResponse
func (svc *sUcSystemMaster) GetMasterInfoByAccountId(id uint32) *system.SysResponse {
	// 查询用户信息
	master, err := repositories.NewUcSystemMasterRepository().GetByAccountID(id)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	if master == nil {
		return &system.SysResponse{
			Code:    1,
			Message: "用户数据为空",
			Data:    nil,
		}
	}

	return &system.SysResponse{
		Code:    0,
		Message: "success",
		Data:    master,
	}
}

// ChangeMasterInfoById
//
// @Title ChangeMasterInfoById
// @Description: 按ID修改个人信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-06 17:20:26
// @receiver svc
// @param in
// @return *system.SysResponse
func (svc *sUcSystemMaster) ChangeMasterInfoById(in dto.ChangeMasterInfoInput) *system.SysResponse {
	// 修改个人信息
	err := repositories.NewUcSystemMasterRepository().UpdateByAccountId(in.ID, &model.UcSystemMaster{
		Email: in.Email,
		Name:  in.Name,
		Tel:   in.Tel,
	})

	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	return &system.SysResponse{
		Code:    0,
		Message: "success",
		Data:    nil,
	}
}

// ShowMasterAll
//
// @Title ShowMasterAll
// @Description: 获取全部管理员信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-20 14:59:07
// @receiver svc
// @return *system.SysResponse
func (svc *sUcSystemMaster) ShowMasterAll() *system.SysResponse {
	output, err := repositories.NewUcSystemMasterRepository().GetAll()
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	var data []*dto.SelectOptionOutput
	for _, v := range output {
		d := dto.SelectOptionOutput{
			Key:   v.AccountID,
			Value: v.Name,
		}
		data = append(data, &d)
	}

	return &system.SysResponse{
		Code:    0,
		Message: "success",
		Data:    data,
	}
}

func (svc *sUcSystemMaster) ShowSystemMaster(in dto.ListSystemMasterInput) *system.SysResponse {
	output, err := repositories.NewUcSystemMasterRepository().GetList(in)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	// 角色信息记录的数量
	totalRecords, err := repositories.NewUcSystemMasterRepository().CountRecords()
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	return &system.SysResponse{
		Code:    0,
		Message: "success",
		Data: dto.ListSystemMasterOutput{
			List:  output,
			Total: totalRecords,
		},
	}
}

// GetMasterInfoById
//
// @Title GetMasterInfoById
// @Description: 按ID获取管理员信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-24 23:36:59
// @receiver svc
// @param id
// @return *system.SysResponse
func (svc *sUcSystemMaster) GetMasterInfoById(id uint32) *system.SysResponse {
	// 查询用户信息
	master, err := repositories.NewUcSystemMasterRepository().GetByID(id)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	if master == nil {
		return &system.SysResponse{
			Code:    1,
			Message: "用户数据为空",
			Data:    nil,
		}
	}

	return &system.SysResponse{
		Code:    0,
		Message: "success",
		Data:    master,
	}
}

// ChangeMasterStatusByID
//
// @Title ChangeMasterStatusByID
// @Description: 按ID修改用户状态
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-08 14:00:55
// @receiver svc
// @param id
// @param status
// @return *system.SysResponse
func (svc *sUcSystemMaster) ChangeMasterStatusByID(id uint32, status uint32) *system.SysResponse {
	err := repositories.NewUcAccountRepository().UpdateStatusById(id, status)
	if err != nil {
		return &system.SysResponse{
			Code:    1,
			Message: err.Error(),
			Data:    nil,
		}
	}

	return &system.SysResponse{
		Code:    0,
		Message: "success",
		Data:    nil,
	}
}
