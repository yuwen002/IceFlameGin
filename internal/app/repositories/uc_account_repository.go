package repositories

import (
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
	"ice_flame_gin/internal/app/models/uc_center"
)

// UcAccountRepository
// @Description:
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

// InsertAccount 插入用户账户
func (repo *UcAccountRepository) InsertAccount(account *uc_center.UcAccount) error {
	return repo.DB.Create(account).Error
}

// GetAccountByUsername 根据用户名获取用户账户
func (repo *UcAccountRepository) GetAccountByUsername(username string) (*uc_center.UcAccount, error) {
	account := &uc_center.UcAccount{}
	err := repo.DB.Where("username = ?", username).First(account).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return account, nil
}

// UpdateAccount 更新用户账户
func (repo *UcAccountRepository) UpdateAccount(account *uc_center.UcAccount) error {
	return repo.DB.Save(account).Error
}

// DeleteAccount 删除用户账户
func (repo *UcAccountRepository) DeleteAccount(account *uc_center.UcAccount) error {
	return repo.DB.Delete(account).Error
}
