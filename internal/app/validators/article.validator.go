package validators

// ArticleCategoryForm
//
// @Description: 文章分类验证
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-17 23:38:35
type ArticleCategoryForm struct {
	Name   string `form:"name" binding:"required,min=2,max=20" msg:"分类名称不能为空|分类名称长度不能小于2个字符|分类名称长度不能超过20个字符 "`
	Remark string `form:"remark"`
	Fid    string `form:"fid"`
	Sort   string `form:"sort"`
	Status string `form:"status"`
}

// ArticleChannelForm
//
// @Description: 文章频道验证
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-17 23:38:50
type ArticleChannelForm struct {
	Name   string `form:"name" binding:"required,min=2,max=20" msg:"分类名称不能为空|分类名称长度不能小于2个字符|分类名称长度不能超过20个字符 "`
	Remark string `form:"remark"`
	Sort   string `form:"sort"`
	Status string `form:"status"`
}
