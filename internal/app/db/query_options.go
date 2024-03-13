package db

// QueryOptions 结构体用于封装查询的条件、分页信息和排序信息
// @Description:
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-13 15:17:11
type QueryOptions struct {
	Field     string
	Condition string
	Args      []interface{}
	Order     string
	Page      int
	PageSize  int
}
