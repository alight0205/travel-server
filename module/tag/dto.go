package tag

import "travel-server/model"

type QueryListReq struct {
	model.PageInfo
	ID   int    `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
}

type QueryDetailReq struct {
	ID int `form:"id" json:"id"`
}

type CreateReq struct {
	Name string `form:"name" json:"name"`
}

type DeleteReq struct {
	ID int `form:"id" json:"id"`
}

type UpdateReq struct {
	ID   int    `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
}
