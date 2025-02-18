package site

import (
	"travel-server/global"
	"travel-server/middleware"
	"travel-server/model"
	"travel-server/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	g := r.Group("/site")
	g.GET("/query_list", utils.WrapHandler(queryList, &QueryListReq{}))              // 查询景点列表
	g.GET("/query_detail", utils.WrapHandler(queryDetail, &QueryDetailReq{}))        // 查询景点详情
	g.POST("/create", middleware.JWTAuth(), utils.WrapHandler(create, &CreateReq{})) // 创建景点
	g.POST("/delete", middleware.JWTAuth(), utils.WrapHandler(delete, &DeleteReq{})) // 删除景点
	g.POST("/update", middleware.JWTAuth(), utils.WrapHandler(update, &UpdateReq{})) // 更新景点
}

// @Tags 景点管理
// @Summary 查询景点列表
// @Description 查询景点列表
// @Router /api/site/query_list [get]
// @Param data query QueryListReq    true  "查询参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func queryList(c *gin.Context, req QueryListReq) (data any, err error) {
	var tags []model.Site
	var total int64
	query := global.DB.Model(&model.Site{})

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

// @Tags 景点管理
// @Summary 查询景点详情
// @Description 查询景点详情
// @Router /api/site/query_detail [get]
// @Param data query QueryDetailReq    true  "查询参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func queryDetail(c *gin.Context, req QueryDetailReq) (data any, err error) {
	var site model.Site
	if err = global.DB.Where("id = ?", req.ID).First(&site).Error; err != nil {
		return
	}
	return site, nil
}

// @Tags 景点管理
// @Summary 创建景点
// @Description 创建景点
// @Router /api/site/create [post]
// @Param data body CreateReq    true  "创建参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func create(c *gin.Context, req CreateReq) (data any, err error) {
	site := model.Site{
		Name: req.Name,
	}
	if err = global.DB.Create(&site).Error; err != nil {
		return
	}
	return
}

// @Tags 景点管理
// @Summary 删除景点
// @Description 删除景点
// @Router /api/site/delete [post]
// @Param data body DeleteReq    true  "删除参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func delete(c *gin.Context, req DeleteReq) (data any, err error) {
	if err = global.DB.Where("id = ?", req.ID).Delete(&model.Site{}).Error; err != nil {
		return
	}
	return
}

// @Tags 景点管理
// @Summary 更新景点
// @Description 更新景点
// @Router /api/site/update [post]
// @Param data body UpdateReq    true  "更新参数"
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func update(c *gin.Context, req UpdateReq) (data any, err error) {
	var site model.Site
	if err = global.DB.Model(&site).Where("id = ?", req.ID).Updates(req).Error; err != nil {
		return
	}
	data = site
	return
}
