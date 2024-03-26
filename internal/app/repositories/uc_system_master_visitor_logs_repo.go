package repositories

import (
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
	"ice_flame_gin/internal/app/dto"
	"ice_flame_gin/internal/app/models/model"
)

// rUcSystemMasterVisitorLogs
//
// @Description: 管理员访问记录
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-26 21:08:30
type rUcSystemMasterVisitorLogs struct {
	DBs map[string]*gorm.DB
	DB  *gorm.DB
}

// NewUcSystemMasterVisitorLogs
//
// @Title NewUcSystemMasterVisitorLogs
// @Description: 创建一个新的 rUcSystemMasterVisitorLogs 仓库实例
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-26 21:08:49
// @return *rUcSystemMasterVisitorLogs
func NewUcSystemMasterVisitorLogs() *rUcSystemMasterVisitorLogs {
	return &rUcSystemMasterVisitorLogs{
		DB: db.DB["default"],
	}
}

// Insert
//
// @Title Insert
// @Description: 管理员访问记录写入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-26 21:34:52
// @receiver repo
// @param data
// @return error
func (repo *rUcSystemMasterVisitorLogs) Insert(data dto.SystemMasterVisitorLogsInput) error {
	err := db.NewGormCore().Insert(&model.UcSystemMasterVisitorLog{
		AccountID:     data.AccountID,
		OsCategory:    data.OsCategory,
		VisitCategory: data.VisitCategory,
		UnionID:       data.UnionID,
		Description:   data.Description,
		IPLong:        data.IPLong,
		IP:            data.IP,
	})

	if err != nil {
		return err
	}

	return nil
}
