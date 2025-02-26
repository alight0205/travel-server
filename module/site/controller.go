package site

import (
	"travel-server/global"
	"travel-server/middleware"
	"travel-server/model"
	"travel-server/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	g1 := r.Group("/admin/site", middleware.JWTAuth(), middleware.AdminAuth())

	g1.GET("/query_list", utils.WrapHandler(queryList, &QueryListReq{})) // 查询评论列表
	g1.POST("/create", utils.WrapHandler(create, &CreateReq{}))          // 创建景点
	g1.POST("/update", utils.WrapHandler(update, nil))                   // 更新景点
	g1.POST("/remove", utils.WrapHandler(remove, &RemoveReq{}))          // 管理员删除文章
}

// @Tags 景点管理
// @Summary 查询景点列表
// @Produce  application/json
// @Param data query QueryListReq    true  "查询参数"
// @Success 200 {object} res.Response{}
// @Router /api/admin/site/query_list [get]
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func queryList(c *gin.Context, req QueryListReq) (data any, err error) {
	var sites []model.Site
	var total int64

	query := global.DB.Model(&model.Site{})
	if req.ID != 0 {
		query = query.Where("id = ?", req.ID)
	}
	if req.Name != "" {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}
	if req.ProvinceCode != 0 {
		query = query.Where("province_code = ?", req.ProvinceCode)
	}
	if req.CityCode != 0 {
		query = query.Where("city_code = ?", req.CityCode)
	}
	if req.AddressDetail != "" {
		query = query.Where("address_detail like ?", "%"+req.AddressDetail+"%")
	}

	query.Count(&total)

	if err = query.Order("created_at desc").Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(&sites).Error; err != nil {
		return
	}

	data = map[string]any{
		"list":  sites,
		"total": total,
	}
	return
}

// @Tags 景点管理
// @Summary 创建景点
// @Produce  application/json
// @Param data body CreateReq    true  "创建参数"
// @Router /api/admin/site/create [post]
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func create(c *gin.Context, req CreateReq) (data any, err error) {
	site := model.Site{
		Name:          req.Name,
		ProvinceCode:  req.ProvinceCode,
		CityCode:      req.CityCode,
		AddressDetail: req.AddressDetail,
		Images:        req.Images,
		Desc:          req.Desc,
	}
	if err = global.DB.Create(&site).Error; err != nil {
		return
	}
	return
}

// @Tags 景点管理
// @Summary 更新景点
// @Produce  application/json
// @Param data body UpdateReq    true  "更新参数"
// @Router /api/admin/site/update [post]
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func update(c *gin.Context, req map[string]any) (data any, err error) {
	id := int(req["id"].(float64))
	utils.FilterProps(req, []string{"name", "province_code", "city_code", "address_detail", "images", "desc"})

	if err = global.DB.Model(&model.Site{}).Where("id = ?", id).Updates(&req).Error; err != nil {
		return
	}
	return
}

// @Tags 景点管理
// @Summary 删除景点
// @Produce  application/json
// @Param data body RemoveReq    true  "删除参数"
// @Router /api/admin/site/remove [post]
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func remove(c *gin.Context, req RemoveReq) (data any, err error) {
	if err = global.DB.Where("id = ?", req.ID).Delete(&model.Site{}).Error; err != nil {
		return
	}
	return
}
