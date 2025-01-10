package repository

import (
	"gorm.io/gorm"
	"mygo/internal/model"
)

type MetaDataRepository interface {
	GetAllArticleMetadata() ([]model.ArticleMetadata, error)
	GetArticleMetadataByCategory(categoryID uint) ([]model.ArticleMetadata, error)
	GetArticleMetadataByTag(tagID uint) ([]model.ArticleMetadata, error)
	GetArticleMetadataByPage(limit, offset int) ([]model.ArticleMetadata, error)
	CountAllArticleMetadata() (int64, error)
}
type metaDataRepository struct {
	db *gorm.DB
}

func NewMetaDataRepository(db *gorm.DB) MetaDataRepository {
	return &metaDataRepository{db: db}
}
func (r *metaDataRepository) GetAllArticleMetadata() ([]model.ArticleMetadata, error) {
	var articleMeta []model.ArticleMetadata
	res := r.db.Model(&model.Article{}).
		Select("id,title,slug,description,created_at,updated_at").
		Find(&articleMeta)
	return articleMeta, res.Error
}
func (r *metaDataRepository) GetArticleMetadataByTag(tagID uint) ([]model.ArticleMetadata, error) {
	var articleMeta []model.ArticleMetadata
	res := r.db.Model(&model.Article{}).
		Select("articles.id, articles.title, articles.slug, articles.description, articles.created_at, articles.updated_at").
		Joins("JOIN article_tags ON article_tags.article_id = articles.id").
		Where("article_tags.tag_id = ?", tagID).
		Find(&articleMeta)
	return articleMeta, res.Error
}

func (r *metaDataRepository) GetArticleMetadataByCategory(categoryID uint) ([]model.ArticleMetadata, error) {
	var articleMeta []model.ArticleMetadata
	res := r.db.Model(&model.Article{}).
		Select("id, title, slug, description, created_at, updated_at").
		Where("category_id = ?", categoryID).
		Find(&articleMeta)
	return articleMeta, res.Error
}

func (r *metaDataRepository) GetArticleMetadataByPage(limit, offset int) ([]model.ArticleMetadata, error) {
	var articleMeta []model.ArticleMetadata
	res := r.db.Model(&model.Article{}).
		Select("id, title, slug, description, created_at, updated_at").
		Limit(limit).
		Offset(offset).
		Find(&articleMeta)
	return articleMeta, res.Error
}

// CountAllArticleMetadata 统计总数
func (r *metaDataRepository) CountAllArticleMetadata() (int64, error) {
	var count int64
	// 统计表 article 的记录数
	res := r.db.Model(&model.Article{}).Count(&count)
	return count, res.Error
}
