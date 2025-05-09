// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameArticle = "article"

// Article mapped from table <article>
type Article struct {
	ID            int       `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	Title         string    `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Desc          string    `gorm:"column:desc;type:varchar(255)" json:"desc"`
	Cover         string    `gorm:"column:cover;type:varchar(255)" json:"cover"`
	Content       string    `gorm:"column:content;type:mediumtext" json:"content"`
	ReadNum       int       `gorm:"column:read_num;type:mediumint;not null" json:"read_num"`
	IsBanner      int       `gorm:"column:is_banner;type:tinyint;not null" json:"is_banner"`
	Creator        int       `gorm:"column:creator;type:int" json:"creator"`
	ProvinceCode  int       `gorm:"column:province_code;type:int" json:"province_code"`
	CityCode      int       `gorm:"column:city_code;type:int" json:"city_code"`
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	ExamineStatus int       `gorm:"column:examine_status;type:int;default:1;comment:1审核通过2审核不通过" json:"examine_status"` // 1审核通过2审核不通过
	User          User     `gorm:"foreignKey:Creator" json:"user"`
	Tags          []Tag    `gorm:"many2many:article_tag" json:"tags"`
	Comments      []Comment `gorm:"foreignKey:ArticleID" json:"comments"`
}

// TableName Article's table name
func (*Article) TableName() string {
	return TableNameArticle
}
