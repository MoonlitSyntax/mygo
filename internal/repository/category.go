package repository

import (
	"errors"
	"mygo/internal/db"
	"mygo/internal/model"
)

func CreateCategory(c *model.Category) error {
	if c.Name == "" {
		return errors.New("类别名不能为空")
	}

	return db.DB.Create(c).Error
}

func DeleteCategoryById(id uint) error {
	return db.DB.Delete(&model.Category{}, id).Error
}

func UpdateCategory(id uint, name string) error {
	return db.DB.Model(&model.Category{}).Where("id = ?", id).Update("name", name).Error
}

func GetCategory(id uint) (*model.Category, error) {
	var category model.Category
	err := db.DB.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func GetAllCategory() ([]model.Category, error) {
	var categories []model.Category
	err := db.DB.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, err
}
