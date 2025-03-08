package article

import (
	"travel-server/model"
)

type _QueryListReq struct {
	model.PageInfo
	ID    int    `form:"id" json:"id"`
	Title string `form:"title" json:"title"`
	Tag   int    `form:"tag" json:"tag"`
}

type _DetailReq struct {
	ID int `form:"id" json:"id"`
}

type _RemoveReq struct {
	ID int `form:"id" json:"id"`
}

type _ExamineReq struct {
	ID            int `form:"id" json:"id"`
	ExamineStatus int `form:"examine_status" json:"examine_status"`
}

type _SetBannerReq struct {
	ID       int `form:"id" json:"id"`
	IsBanner int `form:"is_banner" json:"is_banner"`
}

type QueryMyListReq struct {
	model.PageInfo
	Title string `form:"title" json:"title"`
	Tag   int    `form:"tag" json:"tag"`
}

type QueryListReq struct {
	model.PageInfo
	ID      int    `form:"id" json:"id"`
	Title   string `form:"title" json:"title"`
	Tag     int    `form:"tag" json:"tag"`
	Creator int    `form:"creator" json:"creator"`
}

type CreateReq struct {
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
	Tags    string `form:"tags" json:"tags" binding:"required"`
	Images  string `form:"images" json:"images"`
	Desc    string `form:"desc" json:"desc"`
	Cover   string `form:"cover" json:"cover"`
	TagIDs  string `form:"tag_ids" json:"tag_ids" binding:"required"`
}

type DetailReq struct {
	ID int `form:"id" json:"id"`
}
