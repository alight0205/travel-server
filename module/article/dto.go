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
