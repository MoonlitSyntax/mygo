package repository

import (
	"gorm.io/gorm"
	"mygo/internal/model"
)

type UserRepository interface {
	CreateUser(username, password, email, role string) error
	DeleteUser(id uint) error
	UpdateUser(id uint, update map[string]interface{}) error
	GetUserById(id uint) (*model.User, error)
	CountUser() (int64, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
func (r *userRepository) CreateUser(username, password, email, role string) error {

	user := &model.User{
		Email:    email,
		Username: username,
		Password: password,
		Role:     role,
	}
	return r.db.Create(user).Error
}

func (r *userRepository) DeleteUser(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}

func (r *userRepository) UpdateUser(id uint, update map[string]interface{}) error {
	return r.db.Where("id = ?", id).Updates(update).Error
}

func (r *userRepository) GetUserById(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CountUser() (int64, error) {
	var count int64
	err := r.db.Model(&model.User{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
