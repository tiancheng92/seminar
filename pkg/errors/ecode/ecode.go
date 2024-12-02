package ecode

import "github.com/tiancheng92/seminar/pkg/errors"

const (
	Success int = iota + 100000
	ErrUnknown
	ErrGet
	ErrCreate
	ErrUpdate
	ErrDelete
	ErrParam
	ErrAuth
	ErrTokenInvalid
	ErrDataNotFound
	ErrPageNotFound
	ErrValidation
	ErrTimeOut
	ErrDuplicateKey
	ErrServiceBusy
)

func init() {
	errors.Register(Success, 200, "OK")
	errors.Register(ErrUnknown, 500, "服务端错误")
	errors.Register(ErrGet, 500, "数据获取失败")
	errors.Register(ErrCreate, 500, "创建失败")
	errors.Register(ErrUpdate, 500, "更新失败")
	errors.Register(ErrDelete, 500, "删除失败")
	errors.Register(ErrParam, 400, "参数异常")
	errors.Register(ErrAuth, 403, "权限异常")
	errors.Register(ErrTokenInvalid, 401, "Token无效")
	errors.Register(ErrDataNotFound, 404, "数据未找到")
	errors.Register(ErrPageNotFound, 404, "页面未找到")
	errors.Register(ErrValidation, 400, "参数验证失败")
	errors.Register(ErrTimeOut, 503, "服务端响应超时")
	errors.Register(ErrDuplicateKey, 400, "数据已存在或与现有数据冲突")
	errors.Register(ErrServiceBusy, 503, "服务繁忙")
}
