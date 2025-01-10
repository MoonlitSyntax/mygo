package bizerrors

import (
	"errors"
	"fmt"
	"net/http"
)

/*
 1. BizError：业务层错误结构体，仅包含错误码和错误消息；
 2. 预定义错误码常量；
 3. 错误码 -> 默认消息；
 4. 错误码 -> HTTP 状态码；
 5. 辅助方法：获取默认消息，获取 HTTP 状态码等
*/

// ---------------- 1. BizError 结构体 ----------------
type BizError struct {
	Code    int    // 业务错误码
	Message string // 业务错误描述
}

// 实现 go 原生 error 接口
func (e *BizError) Error() string {
	return e.Message
}

// NewBizError 根据业务错误码和可选的 message 构造 BizError
// 如果传空字符串，会在返回时用默认文案
func NewBizError(code int, message string) error {
	return &BizError{
		Code:    code,
		Message: message,
	}
}

// WrapBizError 用来在捕获原始错误时，包装成 BizError。
// 注意：这是可选写法，也可以直接在业务逻辑中返回 BizError。
func WrapBizError(err error, code int, message string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w", &BizError{
		Code:    code,
		Message: message + ": " + err.Error(),
	})
}

// IsBizError 判断 error 是否为 BizError
// 这样我们可以用 errors.As() / errors.Is() 等方法去区分是不是业务错误。
func IsBizError(err error) bool {
	var be *BizError
	return errors.As(err, &be)
}

// ---------------- 2. 预定义错误码常量 ----------------
// 你可以按需要继续往下补充
const (
	CodeSuccess = 0

	// 通用错误
	CodeServerError   = 10000
	CodeInvalidParams = 10001
	CodeNotFound      = 10002

	// 鉴权相关
	CodeUnauthorizedAuthNotExist  = 10003
	CodeUnauthorizedTokenError    = 10004
	CodeUnauthorizedTokenTimeout  = 10005
	CodeUnauthorizedTokenGenerate = 10006
	CodeTooManyRequests           = 10007

	// 用户模块
	CodeUserNotFound           = 20001
	CodeUserCreateFailed       = 20002
	CodeUserUpdateFailed       = 20003
	CodeUserDeleteFailed       = 20004
	CodeUserUnauthorizedAction = 20005

	// 文章模块
	CodeArticleNotFound         = 21001
	CodeArticleCreateFailed     = 21002
	CodeArticleUpdateFailed     = 21003
	CodeArticleDeleteFailed     = 21004
	CodeArticlePermissionDenied = 21005

	// 标签模块
	CodeTagNotFound           = 22001
	CodeTagCreateFailed       = 22002
	CodeTagUpdateFailed       = 22003
	CodeTagDeleteFailed       = 22004
	CodeTagUnauthorizedAction = 22005

	// 分类模块
	CodeCategoryNotFound           = 23001
	CodeCategoryCreateFailed       = 23002
	CodeCategoryUpdateFailed       = 23003
	CodeCategoryDeleteFailed       = 23004
	CodeCategoryUnauthorizedAction = 23005

	// 元数据模块
	CodeMetaDataNotFound    = 24001
	CodeMetaDataFetchFailed = 24002
)

// ---------------- 3. 错误码 -> 默认消息 ----------------
var codeToMessage = map[int]string{
	CodeSuccess:                   "成功",
	CodeServerError:               "服务内部错误",
	CodeInvalidParams:             "入参错误",
	CodeNotFound:                  "资源不存在",
	CodeUnauthorizedAuthNotExist:  "鉴权失败，找不到对应的 AppKey 和 AppSecret",
	CodeUnauthorizedTokenError:    "鉴权失败，Token 错误",
	CodeUnauthorizedTokenTimeout:  "鉴权失败，Token 超时",
	CodeUnauthorizedTokenGenerate: "鉴权失败，Token 生成失败",
	CodeTooManyRequests:           "请求过多",

	// 用户模块
	CodeUserNotFound:           "用户不存在",
	CodeUserCreateFailed:       "用户创建失败",
	CodeUserUpdateFailed:       "用户更新失败",
	CodeUserDeleteFailed:       "用户删除失败",
	CodeUserUnauthorizedAction: "用户无权限操作",

	// 文章模块
	CodeArticleNotFound:         "文章不存在",
	CodeArticleCreateFailed:     "文章创建失败",
	CodeArticleUpdateFailed:     "文章更新失败",
	CodeArticleDeleteFailed:     "文章删除失败",
	CodeArticlePermissionDenied: "无权限操作文章",

	// 标签模块
	CodeTagNotFound:           "标签不存在",
	CodeTagCreateFailed:       "标签创建失败",
	CodeTagUpdateFailed:       "标签更新失败",
	CodeTagDeleteFailed:       "标签删除失败",
	CodeTagUnauthorizedAction: "无权限操作标签",

	// 分类模块
	CodeCategoryNotFound:           "分类不存在",
	CodeCategoryCreateFailed:       "分类创建失败",
	CodeCategoryUpdateFailed:       "分类更新失败",
	CodeCategoryDeleteFailed:       "分类删除失败",
	CodeCategoryUnauthorizedAction: "无权限操作分类",

	// 元数据模块
	CodeMetaDataNotFound:    "文章元数据未找到",
	CodeMetaDataFetchFailed: "文章元数据获取失败",
}

// GetDefaultMessage 根据错误码获取默认提示信息
func GetDefaultMessage(code int) string {
	if msg, ok := codeToMessage[code]; ok {
		return msg
	}
	return "未知错误"
}

// ---------------- 4. 错误码 -> HTTP 状态码 ----------------
var codeToHTTPStatus = map[int]int{
	CodeSuccess:       http.StatusOK,
	CodeInvalidParams: http.StatusBadRequest,
	CodeNotFound:      http.StatusNotFound,

	CodeUnauthorizedAuthNotExist:  http.StatusUnauthorized,
	CodeUnauthorizedTokenError:    http.StatusUnauthorized,
	CodeUnauthorizedTokenTimeout:  http.StatusUnauthorized,
	CodeUnauthorizedTokenGenerate: http.StatusInternalServerError,
	CodeTooManyRequests:           http.StatusTooManyRequests,

	CodeUserNotFound:           http.StatusNotFound,
	CodeUserUnauthorizedAction: http.StatusForbidden,

	CodeArticleNotFound:         http.StatusNotFound,
	CodeArticlePermissionDenied: http.StatusForbidden,

	CodeTagNotFound:           http.StatusNotFound,
	CodeTagUnauthorizedAction: http.StatusForbidden,

	CodeCategoryNotFound:           http.StatusNotFound,
	CodeCategoryUnauthorizedAction: http.StatusForbidden,

	CodeMetaDataNotFound:    http.StatusNotFound,
	CodeMetaDataFetchFailed: http.StatusInternalServerError,
}

// GetHTTPStatus 根据业务错误码，映射到对应的 HTTP 状态码
func GetHTTPStatus(code int) int {
	if status, ok := codeToHTTPStatus[code]; ok {
		return status
	}
	return http.StatusInternalServerError
}
