package common

import "mime/multipart"

type UploadFileReq struct {
	File *multipart.FileHeader `form:"file"`
}
