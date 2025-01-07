package err_code

import (
	"fmt"
	"net/http"
	"sync"
)

type Error struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Details []string `json:"details"`
}

// 使用 sync.Map 替代普通 map
var codes sync.Map

func NewError(code int, msg string) *Error {
	// 使用 LoadOrStore 确保同一时间只有一个 Goroutine 成功存入相同的错误码
	_, loaded := codes.LoadOrStore(code, msg)
	if loaded {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}

	return &Error{Code: code, Msg: msg}
}

func (e *Error) Msgf(args ...interface{}) string {
	return fmt.Sprintf(e.Msg, args...)
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.Details = []string{}
	for _, d := range details {
		newError.Details = append(newError.Details, d)
	}
	return &newError
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息：%s", e.Code, e.Msg)
}

func (e *Error) StatusCode() int {
	switch e.Code {
	case Success.Code:
		return http.StatusOK
	case InvalidParams.Code:
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code, UnauthorizedTokenError.Code,
		UnauthorizedTokenTimeout.Code, UnauthorizedTokenGenerate.Code:
		return http.StatusUnauthorized
	case TooManyRequests.Code:
		return http.StatusTooManyRequests
	case NotFound.Code:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError // 默认返回 500
	}
}
