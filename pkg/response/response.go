package response

import (
	"github.com/gin-gonic/gin"
	"mygo/pkg/err_code"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func NewResponse(ctx *gin.Context, err *err_code.Error, data interface{}) {
	ctx.JSON(err.StatusCode(), Response{
		Code: err.Code,
		Msg:  err.Msg,
		Data: data,
	})
}
