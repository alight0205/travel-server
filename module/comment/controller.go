package comment

import (
	"travel-server/global"
	"travel-server/middleware"
	"travel-server/model"
	"travel-server/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	g1 := r.Group("/admin/comment", middleware.JWTAuth(), middleware.AdminAuth())
	g2 := r.Group("/user/comment", middleware.JWTAuth(), middleware.AdminAuth())

	g1.GET("/query_list", utils.WrapHandler(queryList, &QueryListReq{})) // 查询评论列表
	g1.POST("/delete", utils.WrapHandler(remove, &DeleteReq{}))          // 删除评论
	g1.POST("/examine", utils.WrapHandler(examine, &ExamineReq{}))       // 管理员审核文章

	g2.POST("/create", utils.WrapHandler(create, &CreateReq{})) // 创建评论
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

	if req.UserId != 0 {
		query = query.Where("user_id = ?", req.UserId)
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
		UserID:    c.GetInt("user_id"),
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

// @Tags 文章管理
// @Summary 管理员审核文章
// @Description 管理员审核文章
// @Router /api/admin/comment/examine [post]
// @Param data body ExamineReq    true  "审核参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func examine(c *gin.Context, req ExamineReq) (data any, err error) {
	if err = global.DB.Model(&model.Comment{}).Where("id = ?", req.ID).Update("examine_status", req.ExamineStatus).Error; err != nil {
		return
	}
	return
}
