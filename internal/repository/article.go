package repository

import (
	"gorm.io/gorm"
	"mygo/internal/model"
)

type ArticleRepository interface {
	CreateArticle(title, content, slug, description string, top bool, belongTo string, userID, categoryID uint, tagsID []uint) error
	GetArticleByID(id uint) (*model.Article, error)
	GetArticlesByPage(limit, offset int) ([]model.Article, error)
	UpdateArticle(id uint, updates map[string]interface{}, tagsID []uint) error
	DeleteArticle(id uint) error
	CountArticle() (int64, error)
	CountArticleByCategory(categoryID uint) (int64, error)
	CountArticleByTag(tagID uint) (int64, error)
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db: db}
}

// CreateArticle 创建文章
func (r *articleRepository) CreateArticle(title, content, slug, description string, top bool, belongTo string, userID, categoryID uint, tagsID []uint) error {
	article := &model.Article{
		Title:       title,
		Content:     content,
		Slug:        slug,
		Description: description,
		Top:         top,
		BelongTo:    belongTo,
		UserID:      userID,
		CategoryID:  categoryID,
	}
	if len(tagsID) > 0 {
		var tags []model.Tag
		if err := r.db.Where("id IN ?", tagsID).Find(&tags).Error; err != nil {
			return err
		}
		article.Tags = tags
	}
	return r.db.Create(article).Error
}

// GetArticleByID 根据 ID 获取文章
func (r *articleRepository) GetArticleByID(id uint) (*model.Article, error) {
	var article model.Article
	if err := r.db.Preload("User").Preload("Category").Preload("Tags").First(&article, id).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

// DeleteArticle  根据 ID 软删除文章
func (r *articleRepository) DeleteArticle(id uint) error {
	res := r.db.Delete(&model.Article{}, id)
	return res.Error
}

// UpdateArticle 更新文章
func (r *articleRepository) UpdateArticle(id uint, updates map[string]interface{}, tagsID []uint) error {
	// 开始事务
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新文章字段
	if len(updates) > 0 {
		if err := tx.Model(&model.Article{}).Where("id = ?", id).Updates(updates).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 更新 Tags 关联
	if len(tagsID) > 0 {
		var tags []model.Tag
		if err := tx.Where("id IN ?", tagsID).Find(&tags).Error; err != nil {
			tx.Rollback()
			return err
		}
		if err := tx.Model(&model.Article{GormModel: model.GormModel{ID: id}}).
			Association("Tags").Replace(tags); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *articleRepository) GetArticlesByPage(limit, offset int) ([]model.Article, error) {
	var articles []model.Article
	err := r.db.Preload("User").Preload("Category").Preload("Tags").
		Limit(limit).Offset(offset).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *articleRepository) CountArticle() (int64, error) {
	var count int64
	err := r.db.Model(&model.Article{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *articleRepository) CountArticleByCategory(categoryID uint) (int64, error) {
	var count int64
	err := r.db.Model(&model.Article{}).Where("category_id = ?", categoryID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *articleRepository) CountArticleByTag(tagID uint) (int64, error) {
	var count int64
	err := r.db.Model(&model.Article{}).
		Joins("JOIN article_tags ON article_tags.article_id = articles.id").
		Where("article_tags.tag_id = ?", tagID).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
