package repository

import (
	"errors"
	"mygo/internal/db"
	"mygo/internal/model"
)

func CreateUser(user *model.User) error {
	if user.Username == "" || user.Password == "" || user.Email == "" {
		return errors.New("用户名 密码 邮箱 不能为空")
	}
	return db.DB.Create(&user).Error
}

func DeleteUserById(id uint) error {
	return db.DB.Delete(&model.User{}, id).Error
}

func UpdateUser(id uint, update map[string]interface{}) error {
	if len(update) == 0 {
		return errors.New("用户更新空字段")
	}

	return db.DB.Where("id = ?", id).Updates(update).Error
}

func GetUserById(id uint) (*model.User, error) {
	var user model.User
	err := db.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
