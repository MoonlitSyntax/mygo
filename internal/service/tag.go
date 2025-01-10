package service

import (
	"fmt"
	"mygo/internal/dto"
	"mygo/internal/repository"
)

type TagService interface {
	CreateTag(req dto.CreateTagRequest) error
	UpdateTag(req dto.UpdateTagRequest, id uint) error
	DeleteTag(req dto.DeleteTagRequest) error
	GetTagByID(id uint) (*dto.TagResponse, error)
	GetTagsByPage(limit, offset int) (*dto.TagListResponse, error)
	GetAllTags() ([]dto.TagResponse, error)
	CountTags() (int64, error)
}

type tagService struct {
	repo repository.TagRepository
}

func NewTagService(repo repository.TagRepository) TagService {
	return &tagService{repo: repo}
}

// CreateTag 创建标签
func (s *tagService) CreateTag(req dto.CreateTagRequest) error {
	return s.repo.CreateTag(req.Name)
}

// UpdateTag 更新标签
func (s *tagService) UpdateTag(req dto.UpdateTagRequest, id uint) error {
	if req.Name == "" {
		return fmt.Errorf("标签名称不能为空")
	}
	return s.repo.UpdateTag(id, req.Name)
}

// DeleteTag 删除标签
func (s *tagService) DeleteTag(req dto.DeleteTagRequest) error {
	if req.ID == 0 {
		return fmt.Errorf("无效的标签 ID")
	}
	return s.repo.DeleteTag(req.ID)
}

// GetTagByID 根据 ID 获取标签
func (s *tagService) GetTagByID(id uint) (*dto.TagResponse, error) {
	tag, err := s.repo.GetTagById(id)
	if err != nil {
		return nil, err
	}

	return &dto.TagResponse{
		ID:   tag.ID,
		Name: tag.Name,
	}, nil
}

// GetTagsByPage 分页获取标签
func (s *tagService) GetTagsByPage(limit, offset int) (*dto.TagListResponse, error) {
	tags, err := s.repo.GetTagsByPage(limit, offset)
	if err != nil {
		return nil, err
	}

	var tagResponses []dto.TagResponse
	for _, tag := range tags {
		tagResponses = append(tagResponses, dto.TagResponse{
			ID:   tag.ID,
			Name: tag.Name,
		})
	}

	total, err := s.repo.CountTag()
	if err != nil {
		return nil, err
	}

	return &dto.TagListResponse{
		Tags:  tagResponses,
		Total: int(total),
	}, nil
}

// GetAllTags 获取所有标签
func (s *tagService) GetAllTags() ([]dto.TagResponse, error) {
	tags, err := s.repo.GetAllTags()
	if err != nil {
		return nil, err
	}

	var tagResponses []dto.TagResponse
	for _, tag := range tags {
		tagResponses = append(tagResponses, dto.TagResponse{
			ID:   tag.ID,
			Name: tag.Name,
		})
	}

	return tagResponses, nil
}

// CountTags 统计标签总数
func (s *tagService) CountTags() (int64, error) {
	return s.repo.CountTag()
}
