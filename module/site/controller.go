package site

import (
	"travel-server/global"
	"travel-server/middleware"
	"travel-server/model"
	"travel-server/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.RouterGroup) {
	_g := r.Group("/admin/site", middleware.JWTAuth(), middleware.AdminAuth())
	_g.GET("/query_list", utils.WrapHandler(_queryList, &_QueryListReq{})) // 查询评论列表
	_g.POST("/create", utils.WrapHandler(_create, &_CreateReq{}))          // 创建景点
	_g.POST("/update", utils.WrapHandler(_update, nil))                    // 更新景点
	_g.POST("/remove", utils.WrapHandler(_remove, &_RemoveReq{}))          // 管理员删除文章

	g := r.Group("/user/site", middleware.JWTAuth())
	g.GET("/query_list", utils.WrapHandler(queryList, &QueryListReq{})) // 获取景点列表
	g.GET("/detail", utils.WrapHandler(detail, &DetailReq{}))           // 获取景点详情
}

// @Tags 景点管理
// @Summary 查询景点列表
// @Produce  application/json
// @Param data query _QueryListReq    true  "查询参数"
// @Success 200 {object} res.Response{}
// @Router /api/admin/site/query_list [get]
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func _queryList(c *gin.Context, req _QueryListReq) (data any, err error) {
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
// @Param data body _CreateReq    true  "创建参数"
// @Router /api/admin/site/create [post]
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func _create(c *gin.Context, req _CreateReq) (data any, err error) {
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
func _update(c *gin.Context, req map[string]any) (data any, err error) {
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
// @Param data body _RemoveReq    true  "删除参数"
// @Router /api/admin/site/remove [post]
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func _remove(c *gin.Context, req _RemoveReq) (data any, err error) {
	if err = global.DB.Where("id = ?", req.ID).Delete(&model.Site{}).Error; err != nil {
		return
	}
	return
}

// @Tags 景点管理
// @Summary 查询景点列表
// @Produce  application/json
// @Param data query QueryListReq    true  "查询参数"
// @Router /api/user/site/query_list [get]
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func queryList(c *gin.Context, req QueryListReq) (data any, err error) {
	var sites []model.Site

	query := global.DB.Model(&model.Site{})

	if req.ProvinceCode != 0 {
		query = query.Where("province_code = ?", req.ProvinceCode)
	}
	if req.CityCode != 0 {
		query = query.Where("city_code = ?", req.CityCode)
	}
	if req.AddressDetail != "" {
		query = query.Where("address_detail like ?", "%"+req.AddressDetail+"%")
	}

	if err = query.Order("created_at desc").Find(&sites).Error; err != nil {
		return
	}
	data = sites
	return
}

// @Tags 景点管理
// @Summary 查询景点详情
// @Produce  application/json
// @Param data query DetailReq    true  "查询参数"
// @Router /api/user/site/detail [get]
// @Param Authorization header string true "Authorization"
// @Produce json
// @Success 200 {object} res.Response{}
func detail(c *gin.Context, req DetailReq) (data any, err error) {
	var site model.Site
	if err = global.DB.Where("id = ?", req.ID).First(&site).Error; err != nil {
		return
	}
	data = site
	return
}
