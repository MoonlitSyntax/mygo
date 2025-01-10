package dto

import "mygo/internal/model"

type CreateArticleRequest struct {
	Title       string `json:"title" binding:"required"`       // 文章标题
	Content     string `json:"content" binding:"required"`     // 文章内容
	Slug        string `json:"slug" binding:"required"`        // 唯一标识
	Description string `json:"description" binding:"required"` // 描述
	Top         bool   `json:"top"`                            // 是否置顶
	BelongTo    string `json:"belong_to" binding:"required"`   // 归属
	UserID      uint   `json:"user_id" binding:"required"`     // 用户 ID
	CategoryID  uint   `json:"category_id" binding:"required"` // 分类 ID
	TagsID      []uint `json:"tags_id" binding:"required"`     // 标签 ID 列表
}
type GetArticlesByPageRequest struct {
	Page     int `json:"page" binding:"required,gte=1"`      // 页码，从 1 开始
	PageSize int `json:"page_size" binding:"required,gte=1"` // 每页条数
}

type UpdateArticleRequest struct {
	ID      uint                   `json:"id" binding:"required,gte=1"`
	Updates map[string]interface{} `json:"updates" binding:"required"` // 需要更新的字段及值
	TagsID  []uint                 `json:"tags_id"`                    // 标签 ID 列表
}
type DeleteArticleRequest struct {
	ID uint `json:"id" binding:"required,gte=1"`
}

type ArticleResponse struct {
	ID          uint             `json:"id"`
	Title       string           `json:"title"`
	Content     string           `json:"content"`
	Slug        string           `json:"slug"`
	Description string           `json:"description"`
	Top         bool             `json:"top"`
	BelongTo    string           `json:"belong_to"`
	UserID      uint             `json:"user_id"`
	CategoryID  uint             `json:"category_id"`
	Tags        []TagResponse    `json:"tags"`       // 嵌套的标签列表
	CreatedAt   model.CustomTime `json:"created_at"` // 创建时间
	UpdatedAt   model.CustomTime `json:"updated_at"` // 更新时间
}

type ArticleListResponse struct {
	Articles []ArticleResponse `json:"articles"`  // 文章列表
	Total    int               `json:"total"`     // 符合条件的总数
	Page     int               `json:"page"`      // 当前第几页（可选）
	PageSize int               `json:"page_size"` // 每页条数（可选）

}
