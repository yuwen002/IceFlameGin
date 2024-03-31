package repositories

import (
	"gorm.io/gorm"
	"ice_flame_gin/internal/app/db"
	"ice_flame_gin/internal/app/dto"
	"ice_flame_gin/internal/app/models/association"
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

// GetListByWhere
//
// @Title GetListByWhere
// @Description:  条件获取管理员访问日志列表
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-27 15:52:17
// @receiver repo
// @param opts
// @return []*association.UcSystemVisitorLogs
// @return error
func (repo *rUcSystemMasterVisitorLogs) GetListByWhere(data dto.ListSystemMasterVisitorLogsInput) ([]*association.UcSystemVisitorLogs, error) {
	var out []*association.UcSystemVisitorLogs
	err := db.NewGormCore().QueryListWithCondition(db.QueryOptions{
		Condition: data.Condition,
		Args:      data.Args,
		Order:     data.Order,
		Preload:   []string{"Master", "VisitCategory"},
		PageType:  2,
		Limit: db.Limit{
			Length: data.Length,
			Offset: data.Start,
		},
	}, &out)

	if err != nil {
		return nil, err
	}

	return out, nil
}

// CountRecords
//
// @Title CountRecords
// @Description: 管理员访问日志列表总条数
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-27 17:15:42
// @receiver repo
// @return int64
// @return error
func (repo *rUcSystemMasterVisitorLogs) CountRecords() (int64, error) {
	totalRecords, err := db.NewGormCore().SetDefaultTable(model.TableNameUcSystemMasterVisitorLog).Count()
	if err != nil {
		return 0, err
	}

	return totalRecords, err
}

// CountWhereRecords
//
// @Title CountWhereRecords
// @Description: 按条件查询管理员访问日志列表总条数
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-27 17:22:04
// @receiver repo
// @param data
// @return int64
// @return error
func (repo *rUcSystemMasterVisitorLogs) CountWhereRecords(data dto.ListSystemMasterVisitorLogsInput) (int64, error) {
	totalRecords, err := db.NewGormCore().SetDefaultTable(model.TableNameUcSystemMasterVisitorLog).CountWhere(data.Condition, data.Args...)
	if err != nil {
		return 0, err
	}

	return totalRecords, err
}
