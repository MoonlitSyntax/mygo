package service

import (
	"fmt"
	"mygo/internal/dto"
	"mygo/internal/repository"
)

type UserService interface {
	CreateUser(req dto.CreateUserRequest) error
	UpdateUser(req dto.UpdateUserRequest, id uint) error
	DeleteUser(req dto.DeleteUserRequest) error
	GetUserByID(id uint) (*dto.UserResponse, error)
	CountUsers() (int64, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// CreateUser 创建用户
func (s *userService) CreateUser(req dto.CreateUserRequest) error {
	return s.repo.CreateUser(req.Username, req.Password, req.Email, req.Role)
}

// UpdateUser 更新用户
func (s *userService) UpdateUser(req dto.UpdateUserRequest, id uint) error {
	updates := map[string]interface{}{}
	if req.Username != "" {
		updates["username"] = req.Username
	}
	if req.Email != "" {
		updates["email"] = req.Email
	}
	if req.Role != "" {
		updates["role"] = req.Role
	}
	if req.Password != "" {
		updates["password"] = req.Password
	}
	if len(updates) == 0 {
		return fmt.Errorf("没有需要更新的字段")
	}
	return s.repo.UpdateUser(id, updates)
}

// DeleteUser 删除用户
func (s *userService) DeleteUser(req dto.DeleteUserRequest) error {
	if req.ID == 0 {
		return fmt.Errorf("无效的用户 ID")
	}
	return s.repo.DeleteUser(req.ID)
}

// GetUserByID 根据 ID 获取用户
func (s *userService) GetUserByID(id uint) (*dto.UserResponse, error) {
	user, err := s.repo.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}, nil
}

// CountUsers 统计用户总数
func (s *userService) CountUsers() (int64, error) {
	return s.repo.CountUser()
}
