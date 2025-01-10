package repository

import (
	"gorm.io/gorm"
	"mygo/internal/model"
)

type TagRepository interface {
	CreateTag(name string) error
	DeleteTag(id uint) error
	UpdateTag(id uint, name string) error
	GetTagById(id uint) (*model.Tag, error)
	GetTagsByPage(limit, offset int) ([]model.Tag, error)
	GetAllTags() ([]model.Tag, error)
	CountTag() (int64, error)
}
type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{db: db}
}
func (r *tagRepository) CreateTag(name string) error {
	tag := model.Tag{Name: name}
	return r.db.Create(&tag).Error
}

func (r *tagRepository) DeleteTag(id uint) error {
	return r.db.Delete(&model.Tag{}, id).Error
}

func (r *tagRepository) UpdateTag(id uint, name string) error {
	return r.db.Model(&model.Tag{}).Where("id = ?", id).Update("name", name).Error
}

func (r *tagRepository) GetTagById(id uint) (*model.Tag, error) {
	var tag model.Tag
	err := r.db.First(&tag, id).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *tagRepository) GetTagsByPage(limit, offset int) ([]model.Tag, error) {
	var tags []model.Tag
	err := r.db.Limit(limit).Offset(offset).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *tagRepository) GetAllTags() ([]model.Tag, error) {
	var tags []model.Tag
	err := r.db.Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}
func (r *tagRepository) CountTag() (int64, error) {
	var count int64
	err := r.db.Model(&model.Tag{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
