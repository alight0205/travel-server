package auth

type LoginReq struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

type RegisterReq struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}
