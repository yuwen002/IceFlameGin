package dto

import "ice_flame_gin/internal/app/models/model"

// SystemMasterVisitCategoryInput
// @Description: 访问型类分类输入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 16:14:29
type SystemMasterVisitCategoryInput struct {
	Title string
}

// ListSystemMasterVisitCategoryInput
//
// @Description: 访问型类分类列表输入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 22:01:13
type ListSystemMasterVisitCategoryInput struct {
	Order  string
	Start  int
	Length int
}

// ListSystemMasterVisitCategoryOutput
//
// @Description: 访问型类分类列表输出
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 22:06:57
type ListSystemMasterVisitCategoryOutput struct {
	List  []*model.UcSystemMasterVisitCategory
	Total int64
}

// SystemMasterVisitorLogsInput
//
// @Description: 管理员访问记录写入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-26 21:33:45
type SystemMasterVisitorLogsInput struct {
	AccountID     uint32
	OsCategory    uint32
	VisitCategory uint32
	UnionID       uint32
	Description   string
	IPLong        string
	IP            string
}
