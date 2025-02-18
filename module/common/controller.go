package common

import (
	"path/filepath"
	"travel-server/global"
	"travel-server/middleware"
	"travel-server/model/res"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func InitRouter(r *gin.RouterGroup) {
	g := r.Group("/common")
	g.GET("/query_oss_config", middleware.JWTAuth(), queryOssConfig) // 获取oss配置
	g.POST("/upload_file", middleware.JWTAuth(), uploadFile)         // 上传文件
}

// @Tags OSS
// @Summary 获取oss配置
// @Description 获取oss配置
// @Router /api/common/query_oss_config [get]
// @Param file formData file true "文件"
// @Produce json
// @Success 200 {object} res.Response
func queryOssConfig(c *gin.Context) {
	// 创建sts客户端
	client, err := sts.NewClientWithAccessKey(global.Config.AliOSS.RegionId, global.Config.AliOSS.AccessKey, global.Config.AliOSS.SecretKey)
	if err != nil {
		panic(err)
	}

	// 创建请求
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "http"
	request.RoleArn = "acs:ram::300418737907777388:role/RamOssUser"
	request.RoleSessionName = "blog-session"
	request.Policy = `{
		 "Version": "1",
		 "Statement": [{
			 "Effect": "Allow",
			 "Action": [
				 "oss:*"
			 ],
			 "Resource": [
				 "acs:oss:*:*:course-blog",
				 "acs:oss:*:*:course-blog/*"
			 ]
		 }]
	 }`

	response, err := client.AssumeRole(request)

	if err != nil {
		panic(err)
	}

	var info = map[string]any{
		"host":       global.Config.AliOSS.CDN,
		"policy":     response,
		"secret_key": global.Config.AliOSS.SecretKey,
		"access_key": global.Config.AliOSS.AccessKey,
	}

	res.OkWithData(info, c)
}

// @Tags OSS
// @Summary 上传文件
// @Description 上传文件
// @Router /api/common/upload_file [post]
// @Param file formData file true "文件"
// @Produce json
// @Success 200 {object} res.Response
func uploadFile(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}

	fileName := uuid.NewV4().String() + filepath.Ext(fileHeader.Filename)
	file, err := fileHeader.Open()
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}

	bucket, err := global.AliOSS.Bucket(global.Config.AliOSS.Bucket)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}

	if err := bucket.PutObject(fileName, file); err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OkWithData(filepath.Join(global.Config.AliOSS.CDN, fileName), c)
}
