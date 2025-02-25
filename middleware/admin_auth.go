package middleware

import (
	"errors"
	"travel-server/global"
	"travel-server/model"
	"travel-server/model/res"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 鉴权用户是否是管理员
func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetInt("user_id")
		user := &model.User{}
		if err := global.DB.First(&user, userId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				res.FailWithMsg("此用户不存在！", c)
				c.Abort()
				return
			}
			res.FailWithMsg("数据库错误！", c)
			c.Abort()
			return
		}
		if user.Role != 1 {
			res.FailWithMsg("您不是管理员", c)
			c.Abort()
			return
		}
	}
}
