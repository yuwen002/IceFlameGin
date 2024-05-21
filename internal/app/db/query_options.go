package db

import "gorm.io/gorm"

// Limit
// @Description: 分页
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-15 09:38:56
type Limit struct {
	Offset int
	Length int
}

// Pagination
// @Description: 分页
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-15 09:39:47
type Pagination struct {
	Page int
	Size int
}

// QueryOptions 结构体用于封装查询的条件、分页信息和排序信息
// @Description:
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-13 15:17:11
type QueryOptions struct {
	Field       string
	Condition   string
	Args        []interface{}
	Order       string
	Omit        string
	Preload     []string
	PreloadFunc map[string]func(db *gorm.DB) *gorm.DB
	PageType    int8
	Pagination  // 分页，页码
	Limit       // 分页，偏移量
}

type GetListOptions struct {
	Order  string
	Start  int
	Length int
}
