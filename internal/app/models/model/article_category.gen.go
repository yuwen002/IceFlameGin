// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameArticleCategory = "article_category"

// ArticleCategory 文章分类表
type ArticleCategory struct {
	ID        uint32    `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true" json:"id"`
	Fid       uint32    `gorm:"column:fid;type:int(10) unsigned;comment:父级ID" json:"fid"`                     // 父级ID
	Name      string    `gorm:"column:name;type:varchar(80);comment:分类名称" json:"name"`                        // 分类名称
	Remark    string    `gorm:"column:remark;type:varchar(100);comment:备注信息" json:"remark"`                   // 备注信息
	Sort      uint32    `gorm:"column:sort;type:int(255) unsigned;comment:排序顺序" json:"sort"`                  // 排序顺序
	Status    uint32    `gorm:"column:status;type:tinyint(3) unsigned;comment:显示状态（0=显示，1=隐藏）" json:"status"` // 显示状态（0=显示，1=隐藏）
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
}

// TableName ArticleCategory's table name
func (*ArticleCategory) TableName() string {
	return TableNameArticleCategory
}
