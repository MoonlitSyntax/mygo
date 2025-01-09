package model

type ArticleMetadata struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Slug        string     `json:"slug"`
	Description string     `json:"description"`
	CreatedAt   CustomTime `json:"created_at"`
	UpdatedAt   CustomTime `json:"updated_at"`
}
