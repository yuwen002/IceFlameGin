package dto

import "ice_flame_gin/internal/app/models/model"

// SinglePageInput
//
// @Description: 单页信息输入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-02 00:15:07
type SinglePageInput struct {
	Title       string
	Description string
	Keyword     string
	Content     string
	Thumbnail   string
	Click       uint32
	Status      uint32
}

// ListSinglePageInput
//
// @Description: 单页信息列表输入
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-02 00:15:48
type ListSinglePageInput struct {
	Order  string
	Start  int
	Length int
}

// ListSinglePageOutput
//
// @Description: 单页信息列表输出
// @Author liuxingyu <yuwen002@163.com>
// @Date 2024-04-02 00:25:57
type ListSinglePageOutput struct {
	List  []*model.SinglePage
	Total int64
}
