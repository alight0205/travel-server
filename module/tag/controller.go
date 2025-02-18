package tag

import (
	"travel-server/global"
	"travel-server/middleware"
	"travel-server/model"
	"travel-server/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	g := r.Group("/tag")
	g.GET("/query_list", utils.WrapHandler(queryList, &QueryListReq{}))              // 查询标签列表
	g.GET("/query_detail", utils.WrapHandler(queryDetail, &QueryDetailReq{}))        // 查询标签详情
	g.POST("/create", middleware.JWTAuth(), utils.WrapHandler(create, &CreateReq{})) // 创建标签
	g.POST("/delete", middleware.JWTAuth(), utils.WrapHandler(delete, &DeleteReq{})) // 删除标签
	g.POST("/update", middleware.JWTAuth(), utils.WrapHandler(update, &UpdateReq{})) // 更新标签
}

// @Tags 标签管理
// @Summary 查询标签列表
// @Description 查询标签列表
// @Router /api/tag/query_list [get]
// @Param data query QueryListReq    true  "查询参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func queryList(c *gin.Context, req QueryListReq) (data any, err error) {
	var tags []model.Tag
	var total int64
	query := global.DB.Model(&model.Tag{})

	if req.ID != 0 {
		query = query.Where("id = ?", req.ID)
	}
	if req.Name != "" {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}
	query.Count(&total)

	if err = query.Order("created_at desc").Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(&tags).Error; err != nil {
		return
	}

	data = map[string]any{
		"list":  tags,
		"total": total,
	}
	return
}

// @Tags 标签管理
// @Summary 查询标签详情
// @Description 查询标签详情
// @Router /api/tag/query_detail [get]
// @Param data query QueryDetailReq    true  "查询参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func queryDetail(c *gin.Context, req QueryDetailReq) (data any, err error) {
	var tag model.Tag
	if err = global.DB.Preload("Articles.Comments").Preload("Articles.Category").Where("id = ?", req.ID).First(&tag).Error; err != nil {
		return
	}
	return tag, nil
}

// @Tags 标签管理
// @Summary 创建标签
// @Description 创建标签
// @Router /api/tag/create [post]
// @Param data body CreateReq    true  "创建参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func create(c *gin.Context, req CreateReq) (data any, err error) {
	tag := model.Tag{
		Name: req.Name,
	}
	if err = global.DB.Create(&tag).Error; err != nil {
		return
	}
	return
}

// @Tags 标签管理
// @Summary 删除标签
// @Description 删除标签
// @Router /api/tag/delete [post]
// @Param data body DeleteReq    true  "删除参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func delete(c *gin.Context, req DeleteReq) (data any, err error) {
	if err = global.DB.Where("id = ?", req.ID).Delete(&model.Tag{}).Error; err != nil {
		return
	}
	return
}

// @Tags 标签管理
// @Summary 更新标签
// @Description 更新标签
// @Router /api/tag/update [post]
// @Param data body UpdateReq    true  "更新参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func update(c *gin.Context, req UpdateReq) (data any, err error) {
	var tag model.Tag
	if err = global.DB.Model(&tag).Where("id = ?", req.ID).Updates(req).Error; err != nil {
		return
	}
	data = tag
	return
}
