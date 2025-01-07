package err_code

var (
	ArticleNotFound         = NewError(21001, "文章不存在")
	ArticleCreateFailed     = NewError(21002, "文章创建失败")
	ArticleUpdateFailed     = NewError(21003, "文章更新失败")
	ArticleDeleteFailed     = NewError(21004, "文章删除失败")
	ArticlePermissionDenied = NewError(21005, "无权限操作文章")
)
