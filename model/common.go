package model

import "time"

type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PageInfo struct {
	PageNum  int `form:"page_num" default:"1"`
	PageSize int `form:"page_size" default:"20"`
}
