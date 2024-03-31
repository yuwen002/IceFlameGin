package validators

// SinglePageForm
//
// @Description: 单页信息form
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-03-29 23:36:27
type SinglePageForm struct {
	Title       string `form:"title" binding:"required,max=80" msg:"标题不能为空|标题长度不能超过80个字符"`
	Description string `form:"description"`
	Keyword     string `form:"keyword"`
	Content     string `form:"content"`
	Thumbnail   string `form:"thumbnail"`
	Click       string `form:"click"`
	Status      string `form:"status"`
}
