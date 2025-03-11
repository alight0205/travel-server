package user

import (
	"travel-server/global"
	"travel-server/middleware"
	"travel-server/model"
	"travel-server/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	_g := r.Group("/admin/user", middleware.JWTAuth(), middleware.AdminAuth())
	_g.GET("/query_list", utils.WrapHandler(_queryList, &_QueryListReq{})) // 查询用户列表
	_g.POST("/update", utils.WrapHandler(_update, nil))                    // 更新用户

	g := r.Group("/user/user", middleware.JWTAuth())
	g.GET("/detail", utils.WrapHandler(detail, &DetailReq{})) // 获取用户详情
	g.POST("/update", utils.WrapHandler(update, nil))         // 更新用户
}

// @Summary 用户列表
// @Tags 用户管理
// @Produce  application/json
// @Param data query _QueryListReq    true  "查询参数"
// @Router /api/admin/user/query_list [get]
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}]
func _queryList(c *gin.Context, req _QueryListReq) (data any, err error) {
	var users []model.User
	var total int64

	query := global.DB.Model(&model.User{})
	if req.ID != 0 {
		query = query.Where("id = ?", req.ID)
	}
	if req.UserName != "" {
		query = query.Where("username like ?", "%"+req.UserName+"%")
	}
	if req.Nickname != "" {
		query = query.Where("nickname like ?", "%"+req.Nickname+"%")
	}
	if req.Role != 0 {
		query = query.Where("role = ?", req.Role)
	}

	query.Count(&total)

	if err = query.Order("created_at desc").Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(&users).Error; err != nil {
		return
	}
	data = map[string]any{
		"list":  users,
		"total": total,
	}
	return
}

// @Summary 管理员更新用户
// @Tags 用户管理
// @Produce  application/json
// @Param data body UpdateReq    true  "更新参数"
// @Router /api/admin/user/update [post]
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func _update(c *gin.Context, req map[string]any) (data any, err error) {
	id := int(req["id"].(float64))
	utils.FilterProps(req, []string{"status", "role"})

	if err = global.DB.Model(&model.User{}).Where("id = ?", id).Updates(&req).Error; err != nil {
		return
	}
	return
}

// @Summary 更新用户
// @Tags 用户管理
// @Produce  application/json
// @Param data body UpdateReq    true  "更新参数"
// @Router /api/admin/user/update [post]
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func update(c *gin.Context, req map[string]any) (data any, err error) {
	id := int(req["id"].(float64))
	utils.FilterProps(req, []string{"password", "nickname", "avatar", "desc", "email"})

	if err = global.DB.Model(&model.Site{}).Where("id = ?", id).Updates(&req).Error; err != nil {
		return
	}
	return
}

func detail(c *gin.Context, req DetailReq) (data any, err error) {
	var user model.User
	if err = global.DB.Select("id", "username", "nickname", "avatar", "desc", "email", "created_at").Where("id = ?", req.ID).First(&user).Error; err != nil {
		return
	}
	data = user
	return
}
