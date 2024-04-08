package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
	"ice_flame_gin/internal/app/models/model"
)

// rUcAccount
// @Description: 用户中心总表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-01 16:34:24
type rUcAccount struct {
	DBs map[string]*gorm.DB
}

// NewUcAccountRepository
//
// @Title NewUcAccountRepository
// @Description: 创建一个新的UcAccountRepository实例，并使用已初始化的DB
// @Author liuxingyu
// @Date 2024-02-03 00:58:14
// @return *rUcAccount
func NewUcAccountRepository() *rUcAccount {
	return &rUcAccount{
		DBs: db.DB, // 直接使用全局的DB实例
	}
}

// GetByID
//
// @Title GetByID
// @Description: 按ID获取用户信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-02 00:16:58
// @receiver repo
// @param id
// @return out
// @return err
func (repo *rUcAccount) GetByID(id uint32) (out *model.UcAccount, err error) {
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
func (repo *rUcAccount) GetAccountByTel(tel string) (*model.UcAccount, error) {
	var account *model.UcAccount
	condition := "username = ?"
	err := db.NewGormCore().QueryOne(&account, condition, tel)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, nil
	}
	return account, nil
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
func (repo *rUcAccount) UpdatePasswordById(id uint32, newPasswordHash string) error {
	account := model.UcAccount{PasswordHash: newPasswordHash}
	err := db.NewGormCore().UpdateByID(id, account)
	if err != nil {
		return err
	}
	return nil
}

// UpdateStatusById
//
// @Title UpdateStatusById
// @Description: 按ID更改用户状态
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-08 11:39:22
// @receiver repo
// @param id
// @param status
// @return error
func (repo *rUcAccount) UpdateStatusById(id uint32, status uint32) error {
	account := model.UcAccount{Status: status}
	fmt.Println("account:", account)
	err := db.NewGormCore().UpdateColumnsByID(id, account)
	if err != nil {
		return err
	}
	return nil
}
