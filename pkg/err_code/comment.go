package err_code

var (
	CommentNotFound     = NewError(22001, "评论不存在")
	CommentCreateFailed = NewError(22002, "评论创建失败")
	CommentDeleteFailed = NewError(22003, "评论删除失败")
)
