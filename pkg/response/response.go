package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	// 假设你的项目 module 名是 your_project_path
	// 请自行替换为自己项目的正确 import
	"mygo/pkg/bizerrors"
)

// Response 统一返回结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// NewResponse 用于在 Controller 中直接返回
//   - 如果 bizErr == nil, 表示成功，返回 code=0
//   - 如果是 BizError, 则转换成相应的错误码和 http.Status
//   - 如果是其他错误, 则默认返回服务内部错误
func NewResponse(ctx *gin.Context, bizErr error, data interface{}) {
	if bizErr == nil {
		ctx.JSON(http.StatusOK, Response{
			Code:    bizerrors.CodeSuccess,
			Message: bizerrors.GetDefaultMessage(bizerrors.CodeSuccess),
			Data:    data,
		})
		return
	}
	var be *bizerrors.BizError
	// 判断是否为 BizError
	if errors.As(bizErr, &be) {
		httpCode := bizerrors.GetHTTPStatus(be.Code)

		// 如果使用者没有自定义 message 则读取默认文案
		msg := be.Message
		if msg == "" {
			msg = bizerrors.GetDefaultMessage(be.Code)
		}

		ctx.JSON(httpCode, Response{
			Code:    be.Code,
			Message: msg,
			Data:    data,
		})
		return
	}

	// 其他错误 统一视为服务内部错误
	ctx.JSON(http.StatusInternalServerError, Response{
		Code:    bizerrors.CodeServerError,
		Message: bizerrors.GetDefaultMessage(bizerrors.CodeServerError) + ": " + bizErr.Error(),
		Data:    nil,
	})
}
