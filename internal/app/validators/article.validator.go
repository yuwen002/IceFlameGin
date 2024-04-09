package validators

type ArticleCategoryForm struct {
	Name   string `form:"name" binding:"required,min=2,max=20" msg:"分类名称不能为空|分类名称长度不能小于2个字符|分类名称长度不能超过20个字符 "`
	Remark string `form:"remark"`
	Fid    string `form:"fid"`
	Sort   string `form:"sort"`
	Status string `form:"status"`
}
