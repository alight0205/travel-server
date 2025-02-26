package site

import "travel-server/model"

type QueryListReq struct {
	model.PageInfo
	ID            int    `form:"id" json:"id"`
	Name          string `form:"name" json:"name"`
	ProvinceCode  int    `form:"province_code" json:"province_code"`
	CityCode      int    `form:"city_code" json:"city_code"`
	AddressDetail string `form:"address_detail" json:"address_detail"`
}

type CreateReq struct {
	Name          string `form:"name" json:"name"`
	ProvinceCode  int    `form:"province_code" json:"province_code"`
	CityCode      int    `form:"city_code" json:"city_code"`
	AddressDetail string `form:"address_detail" json:"address_detail"`
	Images        string `form:"images" json:"images"`
	Desc          string `form:"desc" json:"desc"`
}

type RemoveReq struct {
	ID int `form:"id" json:"id"`
}
