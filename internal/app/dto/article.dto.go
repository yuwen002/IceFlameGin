package dto

// ArticleCategoryInput
//
// @Description: 文章分类信息写入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-10 00:02:56
type ArticleCategoryInput struct {
	Fid    uint32
	Name   string
	Remark string
	Sort   uint32
	Status uint32
}
