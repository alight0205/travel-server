package article

import (
	"travel-server/global"
	"travel-server/middleware"
	"travel-server/model"
	"travel-server/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	g1 := r.Group("/admin/article", middleware.JWTAuth(), middleware.AdminAuth())
	g1.GET("/list", utils.WrapHandler(listByAdmin, &ListByAdminReq{}))    // 管理员查看文章列表
	g1.GET("/detail", utils.WrapHandler(detailByAdmin, &DetailReq{}))     // 管理员查看文章详情
	g1.POST("/remove", utils.WrapHandler(removeByAdmin, &RemoveReq{}))    // 管理员删除文章
	g1.POST("/examine", utils.WrapHandler(examine, &ExamineReq{}))        // 管理员审核文章
	g1.POST("/set_banner", utils.WrapHandler(setBanner, &SetBannerReq{})) // 管理员设置banner
}

// @Tags 文章管理
// @Summary 管理员查询文章列表
// @Description 管理员查询文章列表
// @Router /api/admin/article/list [get]
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

// @Tags 文章管理
// @Summary 管理员查看文章详情
// @Description 管理员查看文章详情
// @Router /api/admin/detail [get]
// @Param data query DetailReq    true  "查询参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func detailByAdmin(c *gin.Context, req DetailReq) (data any, err error) {
	var article model.Article
	if err = global.DB.Where("id = ?", req.ID).Preload("Tags").First(&article).Error; err != nil {
		return
	}
	data = article
	return
}

// @Tags 文章管理
// @Summary 管理员删除文章
// @Description 管理员删除文章
// @Router /api/admin/article/remove [post]
// @Param data body RemoveReq    true  "删除参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func removeByAdmin(c *gin.Context, req RemoveReq) (data any, err error) {
	if err = global.DB.Where("id = ?", req.ID).Delete(&model.Article{}).Error; err != nil {
		return
	}
	return
}

// @Tags 文章管理
// @Summary 管理员审核文章
// @Description 管理员审核文章
// @Router /api/admin/article/examine [post]
// @Param data body ExamineReq    true  "审核参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func examine(c *gin.Context, req ExamineReq) (data any, err error) {
	if err = global.DB.Where("id = ?", req.ID).Update("examine_status", req.ExamineStatus).Error; err != nil {
		return
	}
	return
}

// @Tags 文章管理
// @Summary 管理员设置banner
// @Description 管理员设置banner
// @Router /api/admin/article/set_banner [post]
// @Param data body SetBannerReq    true  "设置参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func setBanner(c *gin.Context, req SetBannerReq) (data any, err error) {
	if err = global.DB.Where("id = ?", req.ID).Update("is_banner", req.IsBanner).Error; err != nil {
		return
	}
	return
}
