package comment

import (
	"travel-server/global"
	"travel-server/middleware"
	"travel-server/model"
	"travel-server/utils"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

func InitRouter(r *gin.RouterGroup) {
	g := r.Group("/comment")
	g.GET("/query_list", utils.WrapHandler(queryList, &QueryListReq{})) // 查询评论列表
	g.GET("/query_by_article", utils.WrapHandler(queryByArticle, &QueryByArticleReq{}))
	g.POST("/create", utils.WrapHandler(create, &CreateReq{}))                            // 创建评论
	g.POST("/delete", middleware.JWTAuth(), utils.WrapHandler(remove, &DeleteReq{}))      // 删除评论
	g.POST("/update", middleware.JWTAuth(), utils.WrapHandler(update, &map[string]any{})) // 更新评论
}

// @Tags 评论管理
// @Summary 查询评论列表
// @Description 查询评论列表
// @Router /api/comment/query_list [get]
// @Param data query QueryListReq    true  "查询参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func queryList(c *gin.Context, req QueryListReq) (data any, err error) {
	var comments []model.Comment
	var total int64
	query := global.DB.Model(&model.Comment{})

	if req.IP != "" {
		query = query.Where("ip = ?", req.IP)
	}
	if req.Account != "" {
		query = query.Where("account like ?", "%"+req.Account+"%")
	}
	if req.Email != "" {
		query = query.Where("email like ?", "%"+req.Email+"%")
	}
	if req.Content != "" {
		query = query.Where("content like ?", "%"+req.Content+"%")
	}
	if req.Province != "" {
		query = query.Where("province like ?", "%"+req.Province+"%")
	}
	if req.City != "" {
		query = query.Where("city like ?", "%"+req.City+"%")
	}
	if req.Examine != 0 {
		query = query.Where("examine = ?", req.Examine)
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
// @Summary 根据文章查询评论
// @Description 根据文章查询评论
// @Router /api/comment/query_by_article [get]
// @Param data query QueryByArticleReq    true  "查询参数"
// @Produce json
// @Success 200 {object} res.Response{}
func queryByArticle(c *gin.Context, req QueryByArticleReq) (data any, err error) {
	var comments []model.Comment
	if err = global.DB.Model(&comments).Where("article_id = ? and examine = 1", req.ArticleID).Find(&comments).Error; err != nil {
		return
	}
	data = comments
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
		IP:        ip,
		Content:   req.Content,
		ArticleID: req.ArticleID,
		CommentID: req.CommentID,
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
// @Param data body DeleteReq    true  "删除参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func remove(c *gin.Context, req DeleteReq) (data any, err error) {
	if err = global.DB.Where("id = ?", req.ID).Delete(&model.Comment{}).Error; err != nil {
		return
	}
	return
}

// @Tags 评论管理
// @Summary 更新评论
// @Description 更新评论
// @Router /api/comment/update [post]
// @Param data body UpdateReq    true  "更新参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func update(c *gin.Context, req map[string]any) (data any, err error) {
	id := int(req["id"].(float64))
	for k, _ := range req {
		if !lo.Contains[string]([]string{"ip", "nickname", "avatar", "content", "province", "city", "examine", "article_id", "comment_id"}, k) {
			delete(req, k)
			continue
		}
	}
	if err = global.DB.Model(&model.Comment{}).Where("id = ?", id).Updates(req).Error; err != nil {
		return
	}
	return
}
