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
	Total    int                       `json:"total"`
}
