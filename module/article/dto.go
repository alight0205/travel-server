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
