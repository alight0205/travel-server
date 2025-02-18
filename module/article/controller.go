package article

import (
	"reflect"
	"travel-server/global"
	"travel-server/middleware"
	"travel-server/model"
	"travel-server/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(r *gin.RouterGroup) {
	g := r.Group("/article")
	g.GET("/list_by_admin", middleware.JWTAuth(), middleware.AdminAuth(), utils.WrapHandler(listByAdmin, &ListByAdminReq{})) // 查询已发布文章列表
	g.GET("/query_detail", utils.WrapHandler(queryDetail, &QueryDetailReq{}))                                                // 查询文章详情
	g.GET("/read", utils.WrapHandler(read, &ReadReq{}))
	g.POST("/create", middleware.JWTAuth(), utils.WrapHandler(create, &CreateReq{}))      // 创建文章
	g.POST("/delete", middleware.JWTAuth(), utils.WrapHandler(remove, &DeleteReq{}))      // 删除文章
	g.POST("/update", middleware.JWTAuth(), utils.WrapHandler(update, &map[string]any{})) // 更新文章
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

// @Tags 文章管理
// @Summary 查询已发布文章列表
// @Description 查询已发布文章列表
// @Router /api/article/query_list [get]
// @Param data query QueryListReq    true  "查询参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func queryList(c *gin.Context, req QueryListReq) (data any, err error) {
	var articles []model.Article
	var total int64
	query := global.DB.Preload("User").Preload("Tags").Preload("Comments").
		Joins(
			"left join article_tag on article_tag.article_id = article.id",
		).
		Model(&model.Article{})

	if req.ID != 0 {
		query = query.Where("id = ?", req.ID)
	}
	if req.Title != "" {
		query = query.Where("title like ?", "%"+req.Title+"%")
	}
	if req.Category != 0 {
		query = query.Where("category_id = ?", req.Category)
	}
	if req.Tag != 0 {
		query = query.Where("article_tag.tag_id = ?", req.Tag)
	}
	if req.Published {
		query = query.Where("published_at is not null")
	}
	if !req.Published {
		query = query.Where("published_at is null")
	}
	if req.IsBanner {
		query = query.Where("is_banner = 1")
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
// @Summary 查询文章详情
// @Description 查询文章详情
// @Router /api/article/query_detail [get]
// @Param data query QueryDetailReq    true  "查询参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func queryDetail(c *gin.Context, req QueryDetailReq) (data any, err error) {
	var article model.Article
	if err = global.DB.Preload("Category").Preload("Tags").Preload("Comments").Where("id = ?", req.ID).First(&article).Error; err != nil {
		return
	}
	data = article
	return
}

func read(c *gin.Context, req ReadReq) (data any, err error) {
	var article model.Article
	if err = global.DB.Model(&article).Where("id = ?", req.ID).UpdateColumn("read_num", gorm.Expr("read_num + ?", 1)).Error; err != nil {
		return
	}
	return
}

// @Tags 文章管理
// @Summary 创建文章
// @Description 创建文章
// @Router /api/article/create [post]
// @Param data body CreateReq    true  "创建参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func create(c *gin.Context, req CreateReq) (data any, err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) (err error) {
		article := model.Article{
			Title:    req.Title,
			Desc:     req.Desc,
			Cover:    req.Cover,
			Content:  req.Content,
			ReadNum:  req.ReadNum,
			IsBanner: req.IsBanner,
			UserID:   c.GetInt("userId"),
		}
		if err = tx.Create(&article).Error; err != nil {
			return
		}
		// 文章关联标签
		for _, tag := range req.Tags {
			var articleTag model.ArticleTag
			if reflect.TypeOf(tag).Kind() == reflect.Float64 {
				articleTag = model.ArticleTag{
					ArticleID: article.ID,
					TagID:     int(tag.(float64)),
				}
			}
			if reflect.TypeOf(tag).Kind() == reflect.String {
				tag := model.Tag{
					Name: tag.(string),
				}
				global.DB.Create(&tag)
				articleTag = model.ArticleTag{
					ArticleID: article.ID,
					TagID:     tag.ID,
				}
			}
			if err = tx.Create(&articleTag).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

// @Tags 文章管理
// @Summary 删除文章
// @Description 删除文章
// @Router /api/article/delete [post]
// @Param data body DeleteReq    true  "删除参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func remove(c *gin.Context, req DeleteReq) (data any, err error) {
	if err = global.DB.Where("id = ?", req.ID).Delete(&model.Article{}).Error; err != nil {
		return
	}
	return
}

// @Tags 文章管理
// @Summary 更新文章
// @Description 更新文章
// @Router /api/article/update [post]
// @Param data body UpdateReq    true  "更新参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func update(c *gin.Context, req map[string]any) (data any, err error) {
	id := int(req["id"].(float64))
	utils.FilterProps(req, []string{"category_id", "title", "desc", "cover", "content", "read_num", "is_banner", "published_at", "tags"})

	err = global.DB.Transaction(func(tx *gorm.DB) (err error) {
		if req["tags"] != nil {
			// 删除文章标签
			if err = tx.Where("article_id = ?", id).Delete(&model.ArticleTag{}).Error; err != nil {
				return
			}
			// 创建文章标签
			for _, tag := range req["tags"].([]any) {
				var articleTag model.ArticleTag
				if reflect.TypeOf(tag).Kind() == reflect.Float64 {
					articleTag = model.ArticleTag{
						ArticleID: id,
						TagID:     int(tag.(float64)),
					}
				}
				if reflect.TypeOf(tag).Kind() == reflect.String {
					tag := model.Tag{
						Name: tag.(string),
					}
					global.DB.Create(&tag)
					articleTag = model.ArticleTag{
						ArticleID: id,
						TagID:     tag.ID,
					}
				}
				if err = tx.Create(&articleTag).Error; err != nil {
					return
				}
			}
			delete(req, "tags")
		}
		// 更新文章
		if err = tx.Model(&model.Article{}).Where("id = ?", id).Updates(&req).Error; err != nil {
			return
		}
		return
	})
	return
}
