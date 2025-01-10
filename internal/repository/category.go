package repository

import (
	"gorm.io/gorm"
	"mygo/internal/model"
)

type CategoryRepository interface {
	CreateCategory(name string) error
	DeleteCategory(id uint) error
	UpdateCategory(id uint, name string) error
	GetCategory(id uint) (*model.Category, error)
	GetAllCategory() ([]model.Category, error)
	CountCategory() (int64, error)
}
type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) CreateCategory(name string) error {
	category := &model.Category{
		Name: name,
	}
	return r.db.Create(category).Error
}

func (r *categoryRepository) DeleteCategory(id uint) error {
	return r.db.Delete(&model.Category{}, id).Error
}

func (r *categoryRepository) UpdateCategory(id uint, name string) error {
	return r.db.Model(&model.Category{}).Where("id = ?", id).Update("name", name).Error
}

func (r *categoryRepository) GetCategory(id uint) (*model.Category, error) {
	var category model.Category
	err := r.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) GetAllCategory() ([]model.Category, error) {
	var categories []model.Category
	err := r.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, err
}

func (r *categoryRepository) CountCategory() (int64, error) {
	var count int64
	err := r.db.Model(&model.Category{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
