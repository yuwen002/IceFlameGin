package repositories

import (
	"errors"
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
	dto "ice_flame_gin/internal/app/dto/d_uc_center"
	models "ice_flame_gin/internal/app/models/m_uc_center"
)

//	UcSystemMasterRepository
//	@Description: 后台管理员
//
// @Author liuxingyu
// @Date 2024-02-08 23:31:58
type UcSystemMasterRepository struct {
	DB *gorm.DB
}

// NewUcSystemMasterRepository
//
// @Title NewUcSystemMasterRepository
// @Description: 创建一个新的 UcSystemMasterRepository 仓库实例
// @Author liuxingyu
// @Date 2024-02-08 23:36:52
// @return *UcSystemMasterRepository 创建一个新的 UcSystemMasterRepository 仓库实例
func NewUcSystemMasterRepository() *UcSystemMasterRepository {
	return &UcSystemMasterRepository{DB: db.DB}
}

func (repo *UcSystemMasterRepository) Insert(data dto.RegisterSystemMasterInput) error {
	tx := repo.DB.Begin()

	// 写入uc_account主表数据
	id, err := db.NewGormCore().SetDefaultTable("uc_account").InsertAndGetID(&models.UcAccount{
		Username:     "SA_" + data.Tel,
		PasswordHash: data.Password,
		Tel:          data.Tel,
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
	err = db.NewGormCore().SetDefaultTable("uc_system_master").Insert(&models.UcSystemMaster{
		AccountID: id,
		Name:      data.Name,
		Tel:       data.Tel,
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
