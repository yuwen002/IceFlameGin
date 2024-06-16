package validators

import "time"

// ArticleCategoryForm
//
// @Description: 文章分类验证
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-17 23:38:35
type ArticleCategoryForm struct {
	Name   string `form:"name" binding:"required,min=2,max=20" msg:"分类名称不能为空|分类名称长度不能小于2个字符|分类名称长度不能超过20个字符"`
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
	Name   string `form:"name" binding:"required,min=2,max=20" msg:"分类名称不能为空|分类名称长度不能小于2个字符|分类名称长度不能超过20个字符"`
	Remark string `form:"remark"`
	Sort   string `form:"sort"`
	Status string `form:"status"`
}

// ArticleTagForm
//
// @Description: 文章标签验证
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-27 00:07:02
type ArticleTagForm struct {
	Name   string `form:"name" binding:"required,min=2,max=20" msg:"标签名称不能为空|标签名称长度不能小于2个字符|标签名称长度不能超过20个字符"`
	Remark string `form:"remark"`
	Sort   string `form:"sort"`
	Status string `form:"status"`
}

// ArticleForm
//
// @Description: 文章验证
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-06-16 23:40:21
type ArticleForm struct {
	Title       string    `form:"title" binding:"required,min=2,max=120" msg:"文章标题不能为空|标签名称长度不能小于2个字符|标签名称长度不能超过120个字符"`
	CategoryID  int       `form:"category_id" binding:"required" msg:"文章分类不能为空"`
	ChannelID   int       `form:"channel_id"`
	Keyword     string    `form:"keyword"`
	Description string    `form:"description"`
	Link        string    `form:"link"`
	Author      string    `form:"author"`
	Tags        []string  `form:"tags"`
	PubDate     time.Time `form:"pub_date"`
	Summary     string    `form:"summary"`
	Content     string    `form:"content"`
	Thumbnail   string    `form:"thumbnail"`
	Click       int       `form:"click"`
	Status      int       `form:"status"`
}
