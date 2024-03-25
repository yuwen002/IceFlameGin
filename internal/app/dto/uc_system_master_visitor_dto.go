package dto

import "ice_flame_gin/internal/app/models/model"

// SystemMasterVisitorCategoryInput
// @Description: 访问型类分类输入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 16:14:29
type SystemMasterVisitorCategoryInput struct {
	Title string
}

// ListSystemMasterVisitorCategoryInput
//
// @Description: 访问型类分类列表输入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 22:01:13
type ListSystemMasterVisitorCategoryInput struct {
	Order  string
	Start  int
	Length int
}

// ListSystemMasterVisitorCategoryOutput
//
// @Description: 访问型类分类列表输出
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-25 22:06:57
type ListSystemMasterVisitorCategoryOutput struct {
	List  []*model.UcSystemMasterVisitCategory
	Total int64
}
