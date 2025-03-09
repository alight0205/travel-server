package tag

type QueryListReq struct {
	Name string `form:"name" json:"name" binding:"required"`
}
