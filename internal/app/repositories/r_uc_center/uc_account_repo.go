package repositories

import (
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
	"ice_flame_gin/internal/app/models/model"
)

// rUcAccountRepository
// @Description: 用户中心总表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-01 16:34:24
type rUcAccountRepository struct {
	DBs map[string]*gorm.DB
}

// NewUcAccountRepository
//
// @Title NewUcAccountRepository
// @Description: 创建一个新的UcAccountRepository实例，并使用已初始化的DB
// @Author liuxingyu
// @Date 2024-02-03 00:58:14
// @return *rUcAccountRepository
func NewUcAccountRepository() *rUcAccountRepository {
	return &rUcAccountRepository{
		DBs: db.DB, // 直接使用全局的DB实例
	}
}

// GetById
//
// @Title GetById
// @Description: 按ID获取用户信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-02 00:16:58
// @receiver repo
// @param id
// @return out
// @return err
func (repo *rUcAccountRepository) GetById(id uint32) (out *model.UcAccount, err error) {
	err = db.NewGormCore().GetByID(id, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
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
func (repo *rUcAccountRepository) GetAccountByTel(tel string) (*model.UcAccount, error) {
	var account model.UcAccount
	condition := "username = ?"
	err := db.NewGormCore().QueryOne(&account, condition, tel)
	if err != nil {
		return nil, err
	}

	if account.ID == 0 {
		return nil, nil
	}
	return &account, nil
}

// UpdatePasswordById
//
// @Title UpdatePasswordById
// @Description: 按ID更新用户密码
// @Author liuxingyu
// @Date 2024-02-25 01:37:25
// @receiver repo
// @param id
// @param newPasswordHash
// @return error
func (repo *rUcAccountRepository) UpdatePasswordById(id uint32, newPasswordHash string) error {
	account := model.UcAccount{PasswordHash: newPasswordHash}
	err := db.NewGormCore().UpdateByID(id, account)
	if err != nil {
		return err
	}
	return nil
}
