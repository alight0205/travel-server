package auth

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
	g := r.Group("/auth")
	{
		g.POST("/login_by_account", utils.WrapHandler(loginByAccount, &LoginReq{}))          // 账号密码登录
		g.GET("/get_login_info", middleware.JWTAuth(), utils.WrapHandler(getLoginInfo, nil)) // 获取当前用户信息

	}
}

// @Tags 鉴权
// @Summary 账号登录
// @Description 账号登录
// @Param data body LoginReq    true  "登录信息"
// @Router /api/auth/login_by_account [post]
// @Produce json
// @Success 200 {object} res.Response{}
func loginByAccount(c *gin.Context, req LoginReq) (data any, err error) {
	user := &model.User{}
	if err = global.DB.Where("username = ? AND password = ? ", req.Username, req.Password).First(&user).Error; err != nil {
		err = errors.New("用户名或密码错误！")
		return
	}
	token, err := utils.JWTSign(user.ID)
	if err != nil {
		return
	}
	data = map[string]any{
		"token": token,
	}
	return
}

// @Tags 鉴权
// @Summary 获取当前登录用户信息
// @Description 获取当前登录用户信息
// @Router /api/auth/get_login_info [get]
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func getLoginInfo(c *gin.Context, _ any) (data any, err error) {
	userId := c.GetInt("user_id")
	user := &model.User{}
	if err = global.DB.First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("此用户不存在！")
			return
		}
		return
	}
	data = map[string]any{
		"user": user,
	}
	return
}
