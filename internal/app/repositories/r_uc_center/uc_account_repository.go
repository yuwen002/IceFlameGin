package repositories

import (
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
	models "ice_flame_gin/internal/app/models/m_uc_center"
)

// UcAccountRepository
// @Description: 用户中心总表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-01 16:34:24
type UcAccountRepository struct {
	DB *gorm.DB
}

// NewUcAccountRepository
//
// @Title NewUcAccountRepository
// @Description: 创建一个新的UcAccountRepository实例，并使用已初始化的DB
// @Author liuxingyu
// @Date 2024-02-03 00:58:14
// @return *UcAccountRepository
func NewUcAccountRepository() *UcAccountRepository {
	return &UcAccountRepository{
		DB: db.DB, // 直接使用全局的DB实例
	}
}

// GetAccountByTel
//
// @Title GetAccountByTel
// @Description: 按电话查找用户信息
// @Author liuxingyu
// @Date 2024-02-08 01:36:18
// @receiver repo
// @param tel
// @return *r_uc_center.UcAccount
// @return error
func (repo *UcAccountRepository) GetAccountByTel(tel string) (*models.UcAccount, error) {
	var account models.UcAccount
	condition := "username = ?"
	err := db.NewGormCore().QueryOne(&account, condition, tel)
	if err != nil {
		return nil, err
	}
	return &account, nil
}
