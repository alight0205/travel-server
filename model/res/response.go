package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FailConfig struct {
	Abort bool
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}
type ListResponse[T any] struct {
	Total int `json:"total"`
	List  T   `json:"list"`
}

const (
	SUCCESS = 0
	ERROR   = 1
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}
func Ok(c *gin.Context) {
	Result(SUCCESS, nil, "success", c)
}
func OkWithData(data any, c *gin.Context) {
	Result(SUCCESS, data, "success", c)
}
func OkWithMsg(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]any{}, msg, c)
}
func OkWithList(list any, total int, c *gin.Context) {
	OkWithData(ListResponse[any]{
		List:  list,
		Total: total,
	}, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]any{}, "error", c)
}
func FailWithMsg(msg string, c *gin.Context) {
	Result(ERROR, map[string]any{}, msg, c)
}
func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	Result(ERROR, map[string]any{}, "未知错误", c)
}
func FailWithAuth(msg string, c *gin.Context) {
	Result(ERROR, map[string]any{}, msg, c)
}
