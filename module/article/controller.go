package article

import (
	"travel-server/global"
	"travel-server/middleware"
	"travel-server/model"
	"travel-server/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	g := r.Group("/article")
	g.GET("/list_by_admin", middleware.JWTAuth(), middleware.AdminAuth(), utils.WrapHandler(listByAdmin, &ListByAdminReq{})) // 查询已发布文章列表
}

// @Tags 文章管理
// @Summary 管理员查询文章列表
// @Description 管理员查询文章列表
// @Router /api/article/list_by_admin [get]
// @Param data query ListByAdminReq    true  "查询参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func listByAdmin(c *gin.Context, req ListByAdminReq) (data any, err error) {
	var articles []model.Article
	var total int64
	query := global.DB.Model(&model.Article{}).Preload("Tags").Joins(
		"left join article_tag on article_tag.article_id = article.id",
	)
	if req.ID != 0 {
		query = query.Where("id = ?", req.ID)
	}
	if req.Title != "" {
		query = query.Where("title like ?", "%"+req.Title+"%")
	}
	if req.Tag != 0 {
		query = query.
			Where("article_tag.tag_id = ?", req.Tag)
	}
	query.Count(&total)
	if err = query.Select("DISTINCT(article.id)", "article.*").Order("created_at desc").Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(&articles).Error; err != nil {
		return
	}
	data = map[string]any{
		"list":  articles,
		"total": total,
	}
	return
}
