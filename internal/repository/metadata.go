package repository

import (
	"mygo/internal/db"
	"mygo/internal/model"
)

func GetAllArticleMetadata() ([]model.ArticleMetadata, error) {
	var articleMeta []model.ArticleMetadata
	res := db.DB.Model(&model.Article{}).
		Select("id,title,slug,description,created_at,updated_at").
		Find(&articleMeta)
	return articleMeta, res.Error
}
