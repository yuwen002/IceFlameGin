package repositories

import (
	"errors"
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
	dto "ice_flame_gin/internal/app/dto/d_uc_center"
	"ice_flame_gin/internal/app/models/model"
)

//	rUcSystemMaster
//	@Description: 后台管理员
//
// @Author liuxingyu
// @Date 2024-02-08 23:31:58
type rUcSystemMaster struct {
	DBs map[string]*gorm.DB
	DB  *gorm.DB
}

// NewUcSystemMasterRepository
//
// @Title NewUcSystemMasterRepository
// @Description: 创建一个新的 rUcSystemMaster 仓库实例
// @Author liuxingyu
// @Date 2024-02-08 23:36:52
// @return *rUcSystemMaster 创建一个新的 rUcSystemMaster 仓库实例
func NewUcSystemMasterRepository() *rUcSystemMaster {
	return &rUcSystemMaster{
		DB: db.DB["default"], //默认使用default
	}
}

// Insert
//
// @Title Insert
// @Description: 写入用户数据
// @Author liuxingyu
// @Date 2024-02-15 22:52:32
// @receiver repo
// @param data
// @return error
func (repo *rUcSystemMaster) Insert(data dto.RegisterSystemMasterInput) error {
	tx := repo.DB.Begin()

	account := "SA_" + data.Tel
	// 写入uc_account主表数据
	id, err := db.NewGormCore().SetDefaultTable(model.TableNameUcAccount).InsertAndGetID(&model.UcAccount{
		Username:     account,
		PasswordHash: data.Password,
		Tel:          account,
	})
	if err != nil {
		// 回滚事务
		tx.Rollback()
		return err
	}
	if id == 0 {
		// 回滚事务
		tx.Rollback()
		return errors.New("account数据库插入错误")
	}

	// 写入uc_system_master数据
	err = db.NewGormCore().SetDefaultTable(model.TableNameUcSystemMaster).Insert(&model.UcSystemMaster{
		AccountID: id,
		Name:      data.Name,
		Tel:       data.Tel,
		Email:     data.Email,
	})
	if err != nil {
		// 回滚事务
		tx.Rollback()
		return err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		// 回滚事务
		tx.Rollback()
		return err
	}

	return nil
}

// GetByEmail
//
// @Title GetByEmail
// @Description: 按Email获取用户信息
// @Author liuxingyu
// @Date 2024-02-15 23:01:52
// @receiver repo
// @param email
// @return *models.UcSystemMaster
// @return error
func (repo *rUcSystemMaster) GetByEmail(email string) (*model.UcSystemMaster, error) {
	var systemMaster model.UcSystemMaster
	condition := "email = ?"
	err := db.NewGormCore().QueryOne(&systemMaster, condition, email)
	if err != nil {
		return nil, err
	}

	if systemMaster.ID == 0 {
		return nil, nil
	}

	return &systemMaster, nil
}

// GetByAccountId
//
// @Title GetByAccountId
// @Description: 按EAccountId获取用户信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-02-27 16:25:16
// @receiver repo
// @param id
// @return *model.UcSystemMaster
// @return error
func (repo *rUcSystemMaster) GetByAccountId(AccountID uint32) (*model.UcSystemMaster, error) {
	var systemMaster *model.UcSystemMaster
	condition := "account_id = ?"
	err := db.NewGormCore().QueryOne(&systemMaster, condition, AccountID)
	if err != nil {
		return nil, err
	}

	if systemMaster.ID == 0 {
		return nil, nil
	}

	return systemMaster, nil
}

// UpdateByAccountId
//
// @Title UpdateByAccountId
// @Description: 按ID修改管理员密码
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-06 17:19:35
// @receiver repo
// @param id
// @param in
// @return error
func (repo *rUcSystemMaster) UpdateByAccountId(id uint32, in *model.UcSystemMaster) error {
	err := db.NewGormCore().UpdateByID(id, in)
	if err != nil {
		return err
	}

	return nil
}

// GetAll
//
// @Title GetAll
// @Description: 获取管理员全部信息
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-20 14:39:49
// @receiver repo
// @return []*model.UcSystemMaster
// @return error
func (repo *rUcSystemMaster) GetAll() ([]*model.UcSystemMaster, error) {
	var systemMaster []*model.UcSystemMaster
	err := db.NewGormCore().GetAll(&systemMaster)
	if err != nil {
		return nil, err
	}

	return systemMaster, nil
}
