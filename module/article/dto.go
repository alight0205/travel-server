package article

import (
	"travel-server/model"
)

type ListByAdminReq struct {
	model.PageInfo
	ID    int    `form:"id" json:"id"`
	Title string `form:"title" json:"title"`
	Tag   int    `form:"tag" json:"tag"`
}

type DetailReq struct {
	ID int `form:"id" json:"id"`
}

type RemoveReq struct {
	ID int `form:"id" json:"id"`
}

type ExamineReq struct {
	ID            int `form:"id" json:"id"`
	ExamineStatus int `form:"examine_status" json:"examine_status"`
}

type SetBannerReq struct {
	ID       int `form:"id" json:"id"`
	IsBanner int `form:"is_banner" json:"is_banner"`
}
