package user

import "travel-server/model"

type _QueryListReq struct {
	model.PageInfo
	ID       int    `form:"id" json:"id"`
	UserName string `form:"name" json:"name"`
	Nickname string `form:"nickname" json:"nickname"`
	Role     int    `form:"role" json:"role"`
}
