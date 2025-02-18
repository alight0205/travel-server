package article

import (
	"time"
	"travel-server/model"
)

type ListByAdminReq struct {
	model.PageInfo
	ID    int    `form:"id" json:"id"`
	Title string `form:"title" json:"title"`
	Tag   int    `form:"tag" json:"tag"`
}

type QueryListReq struct {
	model.PageInfo
	ID        int    `form:"id" json:"id"`
	Title     string `form:"title" json:"title"`
	Category  int    `form:"category" json:"category"`
	Tag       int    `form:"tag" json:"tag"`
	Published bool   `form:"published" json:"published"`
	IsBanner  bool   `form:"is_banner" json:"is_banner"`
}

type QueryDetailReq struct {
	ID int `form:"id" json:"id"`
}

type ReadReq struct {
	ID int `form:"id" json:"id"`
}

type CreateReq struct {
	Title       string     `form:"title" json:"title"`
	Desc        string     `form:"desc" json:"desc"`
	Cover       string     `form:"cover" json:"cover"`
	Content     string     `form:"content" json:"content"`
	ReadNum     int        `form:"read_num" json:"read_num"`
	LikeNum     int        `form:"like_num" json:"like_num"`
	IsBanner    int        `form:"is_banner" json:"is_banner"`
	PublishedAt *time.Time `form:"published_at" json:"published_at"`
	UserID      int        `form:"user_id" json:"user_id"`
	CategoryID  int        `form:"category_id" json:"category_id"`
	Tags        []any      `form:"tags" json:"tags"`
}

type DeleteReq struct {
	ID int `form:"id" json:"id"`
}

type UpdateReq struct {
	ID          int        `form:"id" json:"id"`
	Title       string     `form:"title" json:"title"`
	Desc        string     `form:"desc" json:"desc"`
	Cover       string     `form:"cover" json:"cover"`
	Content     string     `form:"content" json:"content"`
	ReadNum     int        `form:"read_num" json:"read_num"`
	LikeNum     int        `form:"like_num" json:"like_num"`
	IsBanner    int        `form:"is_banner" json:"is_banner"`
	PublishedAt *time.Time `form:"published_at" json:"published_at"`
	UserID      int        `form:"user_id" json:"user_id"`
	CategoryID  int        `form:"category_id" json:"category_id"`
}
