package dto

import "mygo/internal/model"

// ArticleMetadataResponse 文章元数据响应
type ArticleMetadataResponse struct {
	ID          uint             `json:"id"`
	Title       string           `json:"title"`
	Slug        string           `json:"slug"`
	Description string           `json:"description"`
	CreatedAt   model.CustomTime `json:"created_at"`
	UpdatedAt   model.CustomTime `json:"updated_at"`
}

// ArticleMetadataListResponse 文章元数据列表响应
type ArticleMetadataListResponse struct {
	Metadata []ArticleMetadataResponse `json:"metadata"`
	Total    int                       `json:"total"`     // 符合条件的数据总数
	Page     int                       `json:"page"`      // 当前页码
	PageSize int                       `json:"page_size"` // 每页条数
}

type GetArticleMetadataPageRequest struct {
	Page     int `form:"page" binding:"gte=1"`      // 页码
	PageSize int `form:"page_size" binding:"gte=1"` // 每页条数
}
