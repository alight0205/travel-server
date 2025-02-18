package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"travel-server/global"
	"travel-server/model/res"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

func WrapHandler[T any](fn func(*gin.Context, T) (data any, err error), param *T) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req T
		switch c.Request.Method {
		case "POST":
			if err := c.ShouldBindJSON(&req); err != nil {
				res.FailWithMsg(err.Error(), c)
				return
			}
		case "GET":
			if err := c.ShouldBindQuery(&req); err != nil {
				res.FailWithMsg(err.Error(), c)
				return
			}
		default:
			res.FailWithMsg("请求方式错误", c)
			return
		}

		data, err := fn(c, req)

		if err != nil {
			res.FailWithMsg(err.Error(), c)
			return
		}
		if data != nil {
			res.OkWithData(data, c)
			return
		}
		res.Ok(c)
	}
}

func FilterProps(req map[string]any, props []string) {
	for key, _ := range req {
		if !lo.Contains[string](props, key) {
			delete(req, key)
			continue
		}
	}
}

func GetClientIP(c *gin.Context) string {
	ip := c.Request.Header.Get("X-Forwarded-For")
	if ip != "" {
		ips := strings.Split(ip, ",")
		if len(ips) > 0 {
			ip = strings.TrimSpace(ips[0]) // 取第一个IP
		}
	}
	if ip == "" {
		ip = c.Request.Header.Get("X-Real-IP")
	}
	if ip == "" {
		ip = c.Request.RemoteAddr
		if strings.Contains(ip, ":") {
			ip = strings.Split(ip, ":")[0]
		}
	}
	return ip
}

// BaiduMapResponse 表示百度地图API的响应结构
type BaiduMapResponse struct {
	Content struct {
		AddressDetail struct {
			Province string `json:"province"`
			City     string `json:"city"`
		} `json:"address_detail"`
	} `json:"content"`
}

func GetIpLocation(ip string) BaiduMapResponse {
	// 获取ip所在城市
	url := fmt.Sprintf("https://api.map.baidu.com/location/ip?ak=%s&ip=%s&coor=bd09ll", global.Config.BaiDuMap.AppKey, ip)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("请求百度地图API失败:", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应数据失败:", err)
	}

	var baiduMap BaiduMapResponse
	err = json.Unmarshal(body, &baiduMap)
	if err != nil {
		fmt.Println("解析百度地图API响应失败:", err)
	}
	return baiduMap
}
