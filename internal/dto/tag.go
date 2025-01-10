package dto

type CreateTagRequest struct {
	Name string `json:"name" binding:"required"` // 标签名称
}

type UpdateTagRequest struct {
	Name string `json:"name,omitempty"` // 标签名称
}
type DeleteTagRequest struct {
	ID uint `json:"id" binding:"required,gte=1"`
}

type TagResponse struct {
	ID   uint   `json:"id"`   // 标签 ID
	Name string `json:"name"` // 标签名称
}

type TagListResponse struct {
	Tags  []TagResponse `json:"tags"`  // 标签列表
	Total int           `json:"total"` // 标签总数
}
