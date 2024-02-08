package services

import (
	dto "ice_flame_gin/internal/app/dto/d_uc_center"
	repositories "ice_flame_gin/internal/app/repositories/r_uc_center"
	"ice_flame_gin/internal/pkg/utils"
	"ice_flame_gin/internal/system"
)

//	sUcSystemMaster
//	@Description:
//
// @Author liuxingyu
// @Date 2024-02-08 21:50:16
type sUcSystemMaster struct {
	prefix string
}

// NewUcSystemMasterService
//
// @Title NewUcSystemMasterService
// @Description:
// @Author liuxingyu
// @Date 2024-02-08 21:12:49
// @return *sUcSystemMaster
func NewUcSystemMasterService() *sUcSystemMaster {
	return &sUcSystemMaster{
		prefix: "SA_",
	}
}

// LoginTelPassword
//
// @Title LoginTelPassword
// @Description:
// @Author liuxingyu
// @Date 2024-02-08 21:46:35
// @receiver s
// @param in
// @return *system.ApiResponse
func (s sUcSystemMaster) LoginTelPassword(in dto.LoginTelPasswordInput) *system.SysResponse {
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
		Code:    1,
		Message: "success",
		Data:    out,
	}
}
