package repository

import (
	"mygo/internal/db"
	"mygo/internal/model"
	"mygo/internal/util"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateArticle(t *testing.T) {
	err := os.Chdir("/Users/dionysus/code/GolandProjects/mygo")
	util.InitAll()
	db.DB.Migrator().DropTable("blog_service")

	testArticle := &model.Article{
		Title:       "Test Article",
		Content:     "This is a test article content.",
		Slug:        "test-article",
		Description: "A short description for test article.",
		Top:         false,
		BelongTo:    "public",
		UserID:      1,
		CategoryID:  1,
	}

	var num uint
	num = 1
	// 调用 CreateArticle 函数
	err = CreateArticle(testArticle)

	// 验证结果
	assert.NoError(t, err, "创建文章失败")

	// 检查数据库中的文章是否创建成功
	var createdArticle model.Article
	err = db.DB.Preload("User").Preload("Category").First(&createdArticle, testArticle.ID).Error
	assert.NoError(t, err, "查询文章失败")
	assert.Equal(t, "Test Article", createdArticle.Title)

	assert.Equal(t, num, createdArticle.CategoryID)
	assert.Equal(t, "test-article", createdArticle.Slug)
}

func TestDeleteArticleByID(t *testing.T) {
	_ = os.Chdir("/Users/dionysus/code/GolandProjects/mygo")
	util.InitAll()
	DeleteArticleByID(uint(7))
}
