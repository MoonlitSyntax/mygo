package repository

import (
	"errors"
	"mygo/internal/db"
	"mygo/internal/model"
)

func CreateTag(t *model.Tag) error {
	if t.Name == "" {
		return errors.New("标签名不能为空")
	}
	res := db.DB.Create(t)
	return res.Error
}

func DeleteTagById(id uint) error {
	return db.DB.Delete(&model.Tag{}, id).Error
}

func UpdateTag(id uint, name string) error {
	return db.DB.Model(&model.Tag{}).Where("id = ?", id).Update("name", name).Error
}

func GetTagById(id uint) (*model.Tag, error) {
	var tag model.Tag
	err := db.DB.First(&tag, id).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func GetTagsByPage(limit, offset int) ([]model.Tag, error) {
	var tags []model.Tag
	err := db.DB.Limit(limit).Offset(offset).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func GetAllTags() ([]model.Tag, error) {
	var tags []model.Tag
	err := db.DB.Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}
