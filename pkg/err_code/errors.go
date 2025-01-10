package err_code

import (
	"fmt"
	"net/http"
	"sync"
)

// 通用错误码
var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000, "服务内部错误")
	InvalidParams             = NewError(10001, "入参错误")
	NotFound                  = NewError(10002, "资源不存在")
	UnauthorizedAuthNotExist  = NewError(10003, "鉴权失败，找不到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewError(10004, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout  = NewError(10005, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate = NewError(10006, "鉴权失败，Token 生成失败")
	TooManyRequests           = NewError(10007, "请求过多")
)

// 用户模块错误码
var (
	UserNotFound           = NewError(20001, "用户不存在")
	UserCreateFailed       = NewError(20002, "用户创建失败")
	UserUpdateFailed       = NewError(20003, "用户更新失败")
	UserDeleteFailed       = NewError(20004, "用户删除失败")
	UserUnauthorizedAction = NewError(20005, "用户无权限操作")
)

// 文章模块错误码
var (
	ArticleNotFound         = NewError(21001, "文章不存在")
	ArticleCreateFailed     = NewError(21002, "文章创建失败")
	ArticleUpdateFailed     = NewError(21003, "文章更新失败")
	ArticleDeleteFailed     = NewError(21004, "文章删除失败")
	ArticlePermissionDenied = NewError(21005, "无权限操作文章")
)

// 标签模块错误码
var (
	TagNotFound           = NewError(22001, "标签不存在")
	TagCreateFailed       = NewError(22002, "标签创建失败")
	TagUpdateFailed       = NewError(22003, "标签更新失败")
	TagDeleteFailed       = NewError(22004, "标签删除失败")
	TagUnauthorizedAction = NewError(22005, "无权限操作标签")
)

// 分类模块错误码
var (
	CategoryNotFound           = NewError(23001, "分类不存在")
	CategoryCreateFailed       = NewError(23002, "分类创建失败")
	CategoryUpdateFailed       = NewError(23003, "分类更新失败")
	CategoryDeleteFailed       = NewError(23004, "分类删除失败")
	CategoryUnauthorizedAction = NewError(23005, "无权限操作分类")
)

// 元数据模块错误码
var (
	MetaDataNotFound    = NewError(24001, "文章元数据未找到")
	MetaDataFetchFailed = NewError(24002, "文章元数据获取失败")
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
