package service

import (
	"mygo/internal/dto"
	"mygo/internal/repository"
	"mygo/pkg/bizerrors"
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

func (s *userService) CreateUser(req dto.CreateUserRequest) error {
	err := s.repo.CreateUser(req.Username, req.Password, req.Email, req.Role)
	if err != nil {
		return bizerrors.NewBizError(
			bizerrors.CodeUserCreateFailed,
			"创建用户失败: "+err.Error(),
		)
	}
	return nil
}

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
		return bizerrors.NewBizError(
			bizerrors.CodeInvalidParams,
			"没有需要更新的字段",
		)
	}

	err := s.repo.UpdateUser(id, updates)
	if err != nil {
		return bizerrors.NewBizError(
			bizerrors.CodeUserUpdateFailed,
			"更新用户失败: "+err.Error(),
		)
	}
	return nil
}

func (s *userService) DeleteUser(req dto.DeleteUserRequest) error {
	if req.ID == 0 {
		return bizerrors.NewBizError(
			bizerrors.CodeInvalidParams,
			"无效的用户 ID",
		)
	}
	err := s.repo.DeleteUser(req.ID)
	if err != nil {
		return bizerrors.NewBizError(
			bizerrors.CodeUserDeleteFailed,
			"用户删除失败: "+err.Error(),
		)
	}
	return nil
}

func (s *userService) GetUserByID(id uint) (*dto.UserResponse, error) {
	user, err := s.repo.GetUserById(id)
	if err != nil {
		return nil, bizerrors.NewBizError(
			bizerrors.CodeUserNotFound,
			"用户不存在: "+err.Error(),
		)
	}
	return &dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}, nil
}

func (s *userService) CountUsers() (int64, error) {
	count, err := s.repo.CountUser()
	if err != nil {
		return 0, bizerrors.NewBizError(
			bizerrors.CodeServerError,
			"统计用户总数失败: "+err.Error(),
		)
	}
	return count, nil
}
