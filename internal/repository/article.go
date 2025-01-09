package repository

import (
	"errors"
	"mygo/internal/db"
	"mygo/internal/model"
)

// CreateArticle 创建文章
func CreateArticle(article *model.Article) error {
	if article.Title == "" || article.Content == "" || article.CategoryID == 0 {
		return errors.New("标题、内容和分类ID不能为空")
	}

	res := db.DB.Create(article)
	return res.Error
}

// GetArticleByID 根据 ID 获取文章
func GetArticleByID(id uint) (*model.Article, error) {
	var article model.Article
	err := db.DB.First(&article, id).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

// DeleteArticleByID 根据 ID 软删除文章
func DeleteArticleByID(id uint) error {
	res := db.DB.Delete(&model.Article{}, id)
	return res.Error
}

// UpdateArticle 更新文章
func UpdateArticle(id uint, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return errors.New("文章更新空字段")
	}
	res := db.DB.Model(&model.Article{}).Where("id = ?", id).Updates(updates)
	return res.Error
}

func GetArticles(limit, offset int) ([]model.Article, error) {
	var articles []model.Article
	err := db.DB.Limit(limit).Offset(offset).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}
