package err_code

var (
	UserNotFound           = NewError(20001, "用户不存在")
	UserCreateFailed       = NewError(20002, "用户创建失败")
	UserUpdateFailed       = NewError(20003, "用户更新失败")
	UserDeleteFailed       = NewError(20004, "用户删除失败")
	UserUnauthorizedAction = NewError(20005, "用户无权限操作")
)
