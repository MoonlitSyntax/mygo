package repository

import (
	"fmt"
	"gorm.io/gorm"
	"mygo/internal/util"
	"os"
	"testing"
)

func ClearDatabase(db *gorm.DB) error {
	// 获取所有表
	tables, err := db.Migrator().GetTables()
	if err != nil {
		return fmt.Errorf("failed to get tables: %w", err)
	}

	// 删除每个表
	for _, table := range tables {
		if err := db.Migrator().DropTable(table); err != nil {
			return fmt.Errorf("failed to drop table %s: %w", table, err)
		}
	}
	return nil
}

func TestCreateArticle(t *testing.T) {
	_ = os.Chdir("/Users/dionysus/code/GolandProjects/mygo")

	util.InitAll()
	//ClearDatabase(db.DB)

	//var articles []model.Article
	//articleRepository := NewArticleRepository(db.DB)
	//articles, _ = articleRepository.GetArticles(10, 0)
	//fmt.Println(articles)
}
