package middleware

import (
	"travel-server/model/res"
	"travel-server/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			res.FailWithMsg("未携带token", c)
			c.Abort()
			return
		}
		claims, err := utils.JWTVerify(token)
		if err != nil {
			res.FailWithMsg("token错误", c)
			c.Abort()
			return
		}
		// 登录的用户
		c.Set("user_id", claims.UserId)
	}
}
