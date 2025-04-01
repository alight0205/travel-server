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
	ID       int    `form:"id" json:"id"`
	Title    string `form:"title" json:"title"`
	Creator  int    `form:"creator" json:"creator"`
	IsBanner int    `form:"is_banner" json:"is_banner"`
}

type DetailReq struct {
	ID int `form:"id" json:"id"`
}

type CreateReq struct {
	Title        string   `form:"title" json:"title" binding:"required"`
	Desc         string   `form:"desc" json:"desc"`
	Content      string   `form:"content" json:"content" binding:"required"`
	Cover        string   `form:"cover" json:"cover"`
	ProvinceCode int      `form:"province_code" json:"province_code"`
	CityCode     int      `form:"city_code" json:"city_code"`
	Tags         []string `form:"tags" json:"tags"`
}
type RemoveReq struct {
	ID int `form:"id" json:"id"`
}
