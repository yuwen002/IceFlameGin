package repositories

import (
	"errors"
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
	dto "ice_flame_gin/internal/app/dto/d_uc_center"
	"ice_flame_gin/internal/app/models/model"
)

//	UcSystemMasterRepository
//	@Description: 后台管理员
//
// @Author liuxingyu
// @Date 2024-02-08 23:31:58
type UcSystemMasterRepository struct {
	DBs map[string]*gorm.DB
	DB  *gorm.DB
}

// NewUcSystemMasterRepository
//
// @Title NewUcSystemMasterRepository
// @Description: 创建一个新的 UcSystemMasterRepository 仓库实例
// @Author liuxingyu
// @Date 2024-02-08 23:36:52
// @return *UcSystemMasterRepository 创建一个新的 UcSystemMasterRepository 仓库实例
func NewUcSystemMasterRepository() *UcSystemMasterRepository {
	return &UcSystemMasterRepository{
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
func (repo *UcSystemMasterRepository) Insert(data dto.RegisterSystemMasterInput) error {
	tx := repo.DB.Begin()

	account := "SA_" + data.Tel
	// 写入uc_account主表数据
	id, err := db.NewGormCore().SetDefaultTable("uc_account").InsertAndGetID(&model.UcAccount{
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
	err = db.NewGormCore().SetDefaultTable("uc_system_master").Insert(&model.UcSystemMaster{
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
func (repo *UcSystemMasterRepository) GetByEmail(email string) (*model.UcSystemMaster, error) {
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
