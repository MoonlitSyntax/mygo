package service

import (
	"mygo/internal/dto"
	"mygo/internal/repository"
	"mygo/pkg/bizerrors"
)

// CategoryService 接口
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

func (s *categoryService) CreateCategory(req dto.CreateCategoryRequest) error {
	err := s.repo.CreateCategory(req.Name)
	if err != nil {
		return bizerrors.NewBizError(
			bizerrors.CodeCategoryCreateFailed,
			"分类创建失败: "+err.Error(),
		)
	}
	return nil
}

func (s *categoryService) UpdateCategory(req dto.UpdateCategoryRequest, id uint) error {
	if req.Name == "" {
		// 参数错误
		return bizerrors.NewBizError(
			bizerrors.CodeInvalidParams,
			"分类名称不能为空",
		)
	}

	err := s.repo.UpdateCategory(id, req.Name)
	if err != nil {
		return bizerrors.NewBizError(
			bizerrors.CodeCategoryUpdateFailed,
			"分类更新失败: "+err.Error(),
		)
	}
	return nil
}

func (s *categoryService) DeleteCategory(req dto.DeleteCategoryRequest) error {
	if req.ID == 0 {
		// 参数错误
		return bizerrors.NewBizError(
			bizerrors.CodeInvalidParams,
			"无效的分类 ID",
		)
	}

	err := s.repo.DeleteCategory(req.ID)
	if err != nil {
		return bizerrors.NewBizError(
			bizerrors.CodeCategoryDeleteFailed,
			"分类删除失败: "+err.Error(),
		)
	}
	return nil
}

func (s *categoryService) GetCategoryByID(id uint) (*dto.CategoryResponse, error) {
	category, err := s.repo.GetCategory(id)
	if err != nil {
		// 也可以判断是否 gorm.ErrRecordNotFound -> CodeCategoryNotFound
		return nil, bizerrors.NewBizError(
			bizerrors.CodeCategoryNotFound,
			"分类不存在: "+err.Error(),
		)
	}

	return &dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}, nil
}

func (s *categoryService) GetAllCategories() ([]dto.CategoryResponse, error) {
	categories, err := s.repo.GetAllCategory()
	if err != nil {
		return nil, bizerrors.NewBizError(
			bizerrors.CodeServerError,
			"获取所有分类失败: "+err.Error(),
		)
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

func (s *categoryService) CountCategories() (int64, error) {
	count, err := s.repo.CountCategory()
	if err != nil {
		return 0, bizerrors.NewBizError(
			bizerrors.CodeServerError,
			"统计分类失败: "+err.Error(),
		)
	}
	return count, nil
}
