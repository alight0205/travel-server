package tag

import (
	"travel-server/global"
	"travel-server/middleware"
	"travel-server/model"
	"travel-server/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	g := r.Group("/user/tag", middleware.JWTAuth())
	g.GET("/query_list", utils.WrapHandler(queryList, &QueryListReq{})) // 获取标签列表
}

// @Tags 标签管理
// @Summary 查询标签列表
// @Description 查询标签列表
// @Router /api/user/tag/query_list [get]
// @Produce json
// @Success 200 {object} res.Response{}
func queryList(c *gin.Context, req QueryListReq) (data any, err error) {
	var tags []model.Tag
	query := global.DB.Model(&model.Tag{})

	if req.Name == "" {
		return
	}
	if req.Name != "" {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}

	if err = query.Find(&tags).Error; err != nil {
		return
	}
	data = tags
	return
}
