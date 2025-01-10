package service

import (
	"fmt"
	"mygo/internal/dto"
	"mygo/internal/repository"
)

type CategoryService interface {
	CreateCategory(req dto.CreateCategoryRequest) error
	UpdateCategory(req dto.UpdateCategoryRequest, id uint) error
	DeleteCategory(req dto.DeleteCategoryRequest) error
	GetCategoryByID(id uint) (*dto.CategoryResponse, error)
	GetAllCategories() ([]dto.CategoryResponse, error)
	CountCategories() (int64, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

// CreateCategory 创建分类
func (s *categoryService) CreateCategory(req dto.CreateCategoryRequest) error {
	return s.repo.CreateCategory(req.Name)
}

// UpdateCategory 更新分类
func (s *categoryService) UpdateCategory(req dto.UpdateCategoryRequest, id uint) error {
	if req.Name == "" {
		return fmt.Errorf("分类名称不能为空")
	}
	return s.repo.UpdateCategory(id, req.Name)
}

// DeleteCategory 删除分类
func (s *categoryService) DeleteCategory(req dto.DeleteCategoryRequest) error {
	if req.ID == 0 {
		return fmt.Errorf("无效的分类 ID")
	}
	return s.repo.DeleteCategory(req.ID)
}

// GetCategoryByID 根据 ID 获取分类
func (s *categoryService) GetCategoryByID(id uint) (*dto.CategoryResponse, error) {
	category, err := s.repo.GetCategory(id)
	if err != nil {
		return nil, err
	}

	return &dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}, nil
}

// GetAllCategories 获取所有分类
func (s *categoryService) GetAllCategories() ([]dto.CategoryResponse, error) {
	categories, err := s.repo.GetAllCategory()
	if err != nil {
		return nil, err
	}

	var responses []dto.CategoryResponse
	for _, category := range categories {
		responses = append(responses, dto.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return responses, nil
}

// CountCategories 统计分类总数
func (s *categoryService) CountCategories() (int64, error) {
	return s.repo.CountCategory()
}
