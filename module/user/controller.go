package user

import (
	"errors"
	"travel-server/global"
	"travel-server/middleware"
	"travel-server/model"
	"travel-server/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(r *gin.RouterGroup) {
	g := r.Group("/user")
	g.GET("/query_list", middleware.JWTAuth(), utils.WrapHandler(queryList, &QueryListReq{}))     // 获取用户列表
	g.GET("/detail", utils.WrapHandler(detail, &DetailReq{}))                                     // 获取用户详情
	g.POST("/create", utils.WrapHandler(create, &CreateReq{}))                                    // 创建用户
	g.POST("/update", middleware.JWTAuth(), middleware.JWTAuth(), utils.WrapHandler(update, nil)) // 更新用户
}

// 获取用户列表
// @Tags 用户管理
// @Summary 获取用户列表
// @Description 获取用户列表
// @Router /api/user/query_list [get]
// @Param data query QueryListReq    true  "用户信息"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{data=[]UserInfo}
func queryList(c *gin.Context, req QueryListReq) (data any, err error) {
	var users []model.User
	var total int64
	query := global.DB.Preload("Theme").Model(&model.User{}).Where("username like ?", "%"+req.Username+"%")
	if err = query.Count(&total).Error; err != nil {
		return
	}
	if err = query.Offset((req.PageNum - 1) * req.PageSize).Limit(req.PageSize).Scan(&users).Error; err != nil {
		return
	}
	data = map[string]any{
		"list":  users,
		"total": total,
	}
	return
}

// 查询当前登录用户信息
// @Tags 用户管理
// @Summary 获取当前登录用户信息
// @Description 获取当前登录用户信息
// @Router /api/user/detail [get]
// @Param data query DetailReq    true  "用户信息"
// @Produce json
// @Success 200 {object} res.Response{}
func detail(c *gin.Context, req DetailReq) (data any, err error) {
	user := &model.User{}
	if err = global.DB.Omit("password").First(&user, req.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("此用户不存在！")
			return
		}
		return
	}
	data = user
	return
}

// 创建用户
// @Tags 用户管理
// @Summary 创建用户
// @Description 创建用户
// @Router /api/user/create [post]
// @Param data body CreateReq    true  "用户信息"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response
func create(c *gin.Context, req CreateReq) (data any, err error) {
	if err = Create(req); err != nil {
		return
	}
	return
}

// 更新用户
// @Tags 用户管理
// @Summary 更新用户
// @Description 更新用户
// @Router /api/user/update [post]
// @Param data body UpdateReq    true  "用户信息"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response
func update(c *gin.Context, req map[string]any) (data any, err error) {
	userId := c.GetInt("userId")
	utils.FilterProps(req, []string{"username", "password", "nickname", "avatar", "desc", "email", "about", "theme_id"})
	if err = global.DB.Model(&model.User{}).Where("id = ?", userId).Updates(&req).Error; err != nil {
		return
	}
	return
}
