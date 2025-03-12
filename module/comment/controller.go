package comment

import (
	"travel-server/global"
	"travel-server/middleware"
	"travel-server/model"
	"travel-server/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	_g := r.Group("/admin/comment", middleware.JWTAuth(), middleware.AdminAuth())
	_g.GET("/query_list", utils.WrapHandler(_queryList, &_QueryListReq{})) // 查询评论列表
	_g.POST("/delete", utils.WrapHandler(_remove, &_DeleteReq{}))          // 删除评论
	_g.POST("/examine", utils.WrapHandler(_examine, &_ExamineReq{}))       // 管理员审核文章

	g := r.Group("/user/comment", middleware.JWTAuth(), middleware.AdminAuth())
	g.GET("/query_list", utils.WrapHandler(queryList, &QueryListReq{}))                              // 获取我的评论列表
	g.GET("/query_list_by_article", utils.WrapHandler(queryListByArticle, &QueryListByArticleReq{})) // 获取文章评论列表
	g.POST("/create", utils.WrapHandler(create, &CreateReq{}))                                       // 创建评论
	g.POST("/remove", utils.WrapHandler(remove, &RemoveReq{}))                                       // 删除评论
}

// @Tags 评论管理
// @Summary 查询评论列表
// @Description 查询评论列表
// @Router /api/comment/query_list [get]
// @Param data query _QueryListReq    true  "查询参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func _queryList(c *gin.Context, req _QueryListReq) (data any, err error) {
	var comments []model.Comment
	var total int64
	query := global.DB.Model(&model.Comment{})

	if req.Creator != 0 {
		query = query.Where("creator = ?", req.Creator)
	}
	if req.IP != "" {
		query = query.Where("ip = ?", req.IP)
	}
	if req.Content != "" {
		query = query.Where("content like ?", "%"+req.Content+"%")
	}
	if req.Province != "" {
		query = query.Where("province like ?", req.Province)
	}
	if req.City != "" {
		query = query.Where("city like ?", req.City)
	}
	if req.ExamineStatus != 0 {
		query = query.Where("examine = ?", req.ExamineStatus)
	}

	query.Count(&total)

	if err = query.Order("created_at desc").Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(&comments).Error; err != nil {
		return
	}

	data = map[string]any{
		"list":  comments,
		"total": total,
	}
	return
}

// @Tags 评论管理
// @Summary 创建评论
// @Description 创建评论
// @Router /api/comment/create [post]
// @Param data body CreateReq    true  "创建参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func create(c *gin.Context, req CreateReq) (data any, err error) {
	// 获取客户端IP地址
	ip := c.ClientIP()
	if ip == "127.0.0.1" || ip == "::1" {
		ip = "182.150.30.50"
	}
	// 获取ip所在城市
	baiduMap := utils.GetIpLocation(ip)

	comment := model.Comment{
		Creator:   c.GetInt("user_id"),
		Content:   req.Content,
		ArticleID: req.ArticleID,
		CommentID: req.CommentID,
		IP:        ip,
		Province:  baiduMap.Content.AddressDetail.Province,
		City:      baiduMap.Content.AddressDetail.City,
	}
	if err = global.DB.Create(&comment).Error; err != nil {
		return
	}
	return
}

// @Tags 评论管理
// @Summary 删除评论
// @Description 删除评论
// @Router /api/comment/delete [post]
// @Param data body _DeleteReq    true  "删除参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func _remove(c *gin.Context, req _DeleteReq) (data any, err error) {
	if err = global.DB.Where("id = ?", req.ID).Delete(&model.Comment{}).Error; err != nil {
		return
	}
	return
}

// @Tags 评论管理
// @Summary 管理员审核评论
// @Description 管理员审核评论
// @Router /api/admin/comment/examine [post]
// @Param data body _ExamineReq    true  "审核参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func _examine(c *gin.Context, req _ExamineReq) (data any, err error) {
	if err = global.DB.Model(&model.Comment{}).Where("id = ?", req.ID).Update("examine_status", req.ExamineStatus).Error; err != nil {
		return
	}
	return
}

// @Tags 评论管理
// @Summary 查询评论列表
// @Description 查询评论列表
// @Router /api/user/comment/query_list [get]
// @Param data query QueryListReq    true  "查询参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func queryList(c *gin.Context, req QueryListReq) (data any, err error) {
	var results []map[string]any
	var total int64
	query := global.DB.Select("comment.*", "article.title", "user.nickname", "user.avatar").Table("comment").Joins("left join article on comment.article_id = article.id").Joins("left join user on comment.creator = user.id")

	if req.Creator != 0 {
		query = query.Where("comment.creator = ?", req.Creator)
	}

	query.Count(&total)

	if err = query.Order("comment.created_at desc").Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(&results).Error; err != nil {
		return
	}

	data = map[string]any{
		"list":  results,
		"total": total,
	}
	return
}

// @Tags 评论管理
// @Summary 查询文章评论列表
// @Description 查询文章评论列表
// @Router /api/user/comment/query_list_by_article [get]
// @Param data query QueryListReq    true  "查询参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func queryListByArticle(c *gin.Context, req QueryListByArticleReq) (data any, err error) {
	var results []map[string]interface{}
	query := global.DB.Select("comment.*", "article.title", "user.nickname", "user.avatar").Table("comment").Joins("left join article on comment.article_id = article.id").Joins("left join user on comment.creator = user.id")

	query.Where("comment.examine_status = 1")
	if req.ArticleID != 0 {
		query = query.Where("comment.article_id = ?", req.ArticleID)
	}

	if err = query.Order("comment.created_at desc").Find(&results).Error; err != nil {
		return
	}
	data = results
	return
}

// @Tags 评论管理
// @Summary 删除评论
// @Description 删除评论
// @Router /api/user/comment/remove [post]
// @Param data body RemoveReq    true  "删除参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func remove(c *gin.Context, req RemoveReq) (data any, err error) {
	if err = global.DB.Where("id = ?", req.ID).Delete(&model.Comment{}).Error; err != nil {
		return
	}
	return
}
