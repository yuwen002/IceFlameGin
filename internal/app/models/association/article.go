package association

import "ice_flame_gin/internal/app/models/model"

type Article struct {
	model.Article
	Category model.ArticleCategory `gorm:"foreignKey:CategoryID;references:ID" json:"category"`
	Channel  model.ArticleChannel  `gorm:"foreignKey:ChannelID;references:ID" json:"channel"`
}
