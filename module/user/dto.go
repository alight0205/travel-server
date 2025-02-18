package user

import (
	"time"
	"travel-server/model"
)

type UserInfo struct {
	Username  string
	Password  string
	Nickname  string
	Avatar    string
	Desc      string
	Email     string
	About     string
	CreatedAt time.Time
}

type QueryListReq struct {
	model.PageInfo
	Username string `form:"username" json:"username"`
}

type DetailReq struct {
	ID int `form:"id" json:"id"`
}

type CreateReq struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type UpdateReq struct {
	Id       int    `form:"id" json:"id"`
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Nickname string `form:"nickname" json:"nickname"`
	Avatar   string `form:"avatar" json:"avatar"`
	Desc     string `form:"desc" json:"desc"`
	Email    string `form:"email" json:"email"`
	About    string `form:"about" json:"about"`
	ThemeId  int    `form:"theme_id" json:"theme_id"`
}

type DeleteReq struct {
	Id int
}
