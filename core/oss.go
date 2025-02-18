package core

import (
	"log"
	"travel-server/global"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func InitOSS() *oss.Client {
	OSSClient, err := oss.New(global.Config.AliOSS.Endpoint, global.Config.AliOSS.AccessKey, global.Config.AliOSS.SecretKey)
	if err != nil {
		global.Log.Errorf("oss connect error: %s", err)
		return nil
	}
	log.Println("oss load Init success")
	return OSSClient
}
