package article

import (
	"travel-server/global"
	"travel-server/middleware"
	"travel-server/model"
	"travel-server/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(r *gin.RouterGroup) {
	_g := r.Group("/admin/article", middleware.JWTAuth(), middleware.AdminAuth())
	_g.GET("/query_list", utils.WrapHandler(_queryList, &_QueryListReq{}))  // 管理员查看文章列表
	_g.GET("/detail", utils.WrapHandler(_detail, &_DetailReq{}))            // 管理员查看文章详情
	_g.POST("/remove", utils.WrapHandler(_remove, &_RemoveReq{}))           // 管理员删除文章
	_g.POST("/examine", utils.WrapHandler(_examine, &_ExamineReq{}))        // 管理员审核文章
	_g.POST("/set_banner", utils.WrapHandler(_setBanner, &_SetBannerReq{})) // 管理员设置banner

	g := r.Group("/user/article", middleware.JWTAuth())
	g.GET("/query_list", utils.WrapHandler(queryList, &QueryListReq{}))        // 查询用户文章列表
	g.GET("/query_my_list", utils.WrapHandler(queryMyList, &QueryMyListReq{})) // 查询我的文章列表
	g.GET("/detail", utils.WrapHandler(detail, &DetailReq{}))                  // 获取文章详情
	g.POST("/create", utils.WrapHandler(create, &CreateReq{}))                 // 创建文章
}

// @Tags 文章管理
// @Summary 管理员查询文章列表
// @Description 管理员查询文章列表
// @Router /api/admin/article/query_list [get]
// @Param data query _QueryListReq    true  "查询参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func _queryList(c *gin.Context, req _QueryListReq) (data any, err error) {
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

	if err = query.Select("COUNT(DISTINCT article.id)").Count(&total).Error; err != nil {
		return
	}

	if err = query.Select("DISTINCT(article.id)", "article.*").Order("article.created_at desc").Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(&articles).Error; err != nil {
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
func _detail(c *gin.Context, req _DetailReq) (data any, err error) {
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
// @Param data body _RemoveReq    true  "删除参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func _remove(c *gin.Context, req _RemoveReq) (data any, err error) {
	if err = global.DB.Where("id = ?", req.ID).Delete(&model.Article{}).Error; err != nil {
		return
	}
	return
}

// @Tags 文章管理
// @Summary 管理员审核文章
// @Description 管理员审核文章
// @Router /api/admin/article/examine [post]
// @Param data body _ExamineReq    true  "审核参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func _examine(c *gin.Context, req _ExamineReq) (data any, err error) {
	if err = global.DB.Model(&model.Article{}).Where("id = ?", req.ID).Update("examine_status", req.ExamineStatus).Error; err != nil {
		return
	}
	return
}

// @Tags 文章管理
// @Summary 管理员设置banner
// @Description 管理员设置banner
// @Router /api/admin/article/set_banner [post]
// @Param data body _SetBannerReq    true  "设置参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func _setBanner(c *gin.Context, req _SetBannerReq) (data any, err error) {
	if err = global.DB.Model(&model.Article{}).Where("id = ?", req.ID).Update("is_banner", req.IsBanner).Error; err != nil {
		return
	}
	return
}

// @Tags 文章管理
// @Summary 查询我的文章列表
// @Description 查询文章列表
// @Router /api/user/article/query_my_list [get]
// @Param data query QueryMyListReq    true  "查询参数"
// @Produce json
// @Success 200 {object} res.Response{}
func queryMyList(c *gin.Context, req QueryMyListReq) (data any, err error) {
	userId := c.GetInt("user_id")
	var articles []model.Article
	var total int64

	query := global.DB.Model(&model.Article{}).Preload("Tags").Joins(
		"left join article_tag on article_tag.article_id = article.id",
	)

	query = query.Where("creator = ?", userId)

	if req.Title != "" {
		query = query.Where("title like ?", "%"+req.Title+"%")
	}
	if req.Tag != 0 {
		query = query.
			Where("article_tag.tag_id = ?", req.Tag)
	}

	if err = query.Select("COUNT(DISTINCT article.id)").Count(&total).Error; err != nil {
		return
	}

	if err = query.Select("DISTINCT(article.id)", "article.*").Order("article.created_at desc").Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(&articles).Error; err != nil {
		return
	}
	data = map[string]any{
		"list":  articles,
		"total": total,
	}
	return
}

// @Tags 文章管理
// @Summary 查询文章列表
// @Description 查询文章列表
// @Router /api/user/article/query_list [get]
// @Param data query QueryListReq    true  "查询参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func queryList(c *gin.Context, req QueryListReq) (data any, err error) {
	var articles []model.Article
	var total int64
	query := global.DB.Model(&model.Article{}).Preload("Tags").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "nickname", "avatar") // 只选择需要的字段
	})
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
	if req.Creator != 0 {
		query = query.Where("creator = ?", req.Creator)
	}
	if req.IsBanner != 0 {
		query = query.Where("is_banner = ?", req.IsBanner)
	}

	if err = query.Select("COUNT(DISTINCT article.id)").Count(&total).Error; err != nil {
		return
	}

	if err = query.Select("article.*").Order("article.created_at desc").Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(&articles).Error; err != nil {
		return
	}
	data = map[string]any{
		"list":  articles,
		"total": total,
	}
	return
}

// @Tags 文章管理
// @Summary 查询文章详情
// @Description 查询文章详情
// @Router /api/user/article/detail [get]
// @Param data query DetailReq    true  "查询参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func detail(c *gin.Context, req DetailReq) (data any, err error) {
	var article model.Article
	if err = global.DB.Where("id = ?", req.ID).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "nickname", "avatar") // 只选择需要的字段
	}).Preload("Tags").First(&article).Error; err != nil {
		return
	}
	data = article
	return
}

// @Tags 文章管理
// @Summary 创建文章
// @Description 创建文章
// @Router /api/user/article/create [post]
// @Param data body CreateReq    true  "创建参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func create(c *gin.Context, req CreateReq) (data any, err error) {
	userId := c.GetInt("user_id")

	article := model.Article{
		Creator:      userId,
		Title:        req.Title,
		Desc:         req.Desc,
		Content:      req.Content,
		Cover:        req.Cover,
		ProvinceCode: req.ProvinceCode,
		CityCode:     req.CityCode,
	}

	if err = global.DB.Create(&article).Error; err != nil {
		return
	}
	return
}
