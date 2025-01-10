package service

import (
	"mygo/internal/dto"
	"mygo/internal/repository"
)

// MetaDataService 定义元数据服务接口
type MetaDataService interface {
	GetAllArticleMetadata() (*dto.ArticleMetadataListResponse, error)
	GetArticleMetadataByCategory(categoryID uint) (*dto.ArticleMetadataListResponse, error)
	GetArticleMetadataByTag(tagID uint) (*dto.ArticleMetadataListResponse, error)
}
type metaDataService struct {
	repo repository.MetaDataRepository
}

func NewMetaDataService(repo repository.MetaDataRepository) MetaDataService {
	return &metaDataService{repo: repo}
}

// GetAllArticleMetadata 获取所有文章的元数据
func (s *metaDataService) GetAllArticleMetadata() (*dto.ArticleMetadataListResponse, error) {
	metadata, err := s.repo.GetAllArticleMetadata()
	if err != nil {
		return nil, err
	}

	response := make([]dto.ArticleMetadataResponse, len(metadata))
	for i, meta := range metadata {
		response[i] = dto.ArticleMetadataResponse{
			ID:          meta.ID,
			Title:       meta.Title,
			Slug:        meta.Slug,
			Description: meta.Description,
			CreatedAt:   meta.CreatedAt,
			UpdatedAt:   meta.UpdatedAt,
		}
	}

	return &dto.ArticleMetadataListResponse{
		Metadata: response,
		Total:    len(response),
	}, nil
}

// GetArticleMetadataByCategory 根据分类获取文章的元数据
func (s *metaDataService) GetArticleMetadataByCategory(categoryID uint) (*dto.ArticleMetadataListResponse, error) {
	metadata, err := s.repo.GetArticleMetadataByCategory(categoryID)
	if err != nil {
		return nil, err
	}

	response := make([]dto.ArticleMetadataResponse, len(metadata))
	for i, meta := range metadata {
		response[i] = dto.ArticleMetadataResponse{
			ID:          meta.ID,
			Title:       meta.Title,
			Slug:        meta.Slug,
			Description: meta.Description,
			CreatedAt:   meta.CreatedAt,
			UpdatedAt:   meta.UpdatedAt,
		}
	}

	return &dto.ArticleMetadataListResponse{
		Metadata: response,
		Total:    len(response),
	}, nil
}

// GetArticleMetadataByTag 根据标签获取文章的元数据
func (s *metaDataService) GetArticleMetadataByTag(tagID uint) (*dto.ArticleMetadataListResponse, error) {
	metadata, err := s.repo.GetArticleMetadataByTag(tagID)
	if err != nil {
		return nil, err
	}

	response := make([]dto.ArticleMetadataResponse, len(metadata))
	for i, meta := range metadata {
		response[i] = dto.ArticleMetadataResponse{
			ID:          meta.ID,
			Title:       meta.Title,
			Slug:        meta.Slug,
			Description: meta.Description,
			CreatedAt:   meta.CreatedAt,
			UpdatedAt:   meta.UpdatedAt,
		}
	}

	return &dto.ArticleMetadataListResponse{
		Metadata: response,
		Total:    len(response),
	}, nil
}
