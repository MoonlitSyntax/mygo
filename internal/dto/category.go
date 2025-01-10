package dto

// CreateCategoryRequest 创建分类请求
type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"` // 分类名称
}

// UpdateCategoryRequest 更新分类请求
type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"required"` // 分类名称
}

// CategoryResponse 分类响应
type CategoryResponse struct {
	ID   uint   `json:"id"`   // 分类 ID
	Name string `json:"name"` // 分类名称
}

// CategoryListResponse 分类列表响应
type CategoryListResponse struct {
	Categories []CategoryResponse `json:"categories"` // 分类列表
	Total      int                `json:"total"`      // 总数
}

type DeleteCategoryRequest struct {
	ID uint `json:"id" binding:"required,gte=1"`
}
