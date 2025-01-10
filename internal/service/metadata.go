package service

import (
	"mygo/internal/dto"
	"mygo/internal/repository"
	"mygo/pkg/bizerrors"
)

// MetaDataService 定义元数据服务接口
type MetaDataService interface {
	GetAllArticleMetadata() (*dto.ArticleMetadataListResponse, error)
	GetArticleMetadataByCategory(categoryID uint) (*dto.ArticleMetadataListResponse, error)
	GetArticleMetadataByTag(tagID uint) (*dto.ArticleMetadataListResponse, error)
	GetArticleMetadataByPage(req dto.GetArticleMetadataPageRequest) (*dto.ArticleMetadataListResponse, error)
}

type metaDataService struct {
	repo repository.MetaDataRepository
}

func NewMetaDataService(repo repository.MetaDataRepository) MetaDataService {
	return &metaDataService{repo: repo}
}

func (s *metaDataService) GetAllArticleMetadata() (*dto.ArticleMetadataListResponse, error) {
	metadata, err := s.repo.GetAllArticleMetadata()
	if err != nil {
		// 这里可用 CodeMetaDataFetchFailed 或 CodeServerError
		return nil, bizerrors.NewBizError(
			bizerrors.CodeMetaDataFetchFailed,
			"获取所有文章元数据失败: "+err.Error(),
		)
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

func (s *metaDataService) GetArticleMetadataByCategory(categoryID uint) (*dto.ArticleMetadataListResponse, error) {
	metadata, err := s.repo.GetArticleMetadataByCategory(categoryID)
	if err != nil {
		// 同上
		return nil, bizerrors.NewBizError(
			bizerrors.CodeMetaDataFetchFailed,
			"获取分类文章元数据失败: "+err.Error(),
		)
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

func (s *metaDataService) GetArticleMetadataByTag(tagID uint) (*dto.ArticleMetadataListResponse, error) {
	metadata, err := s.repo.GetArticleMetadataByTag(tagID)
	if err != nil {
		return nil, bizerrors.NewBizError(
			bizerrors.CodeMetaDataFetchFailed,
			"获取标签文章元数据失败: "+err.Error(),
		)
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

func (s *metaDataService) GetArticleMetadataByPage(req dto.GetArticleMetadataPageRequest) (*dto.ArticleMetadataListResponse, error) {
	limit := req.PageSize
	offset := (req.Page - 1) * req.PageSize

	metadata, err := s.repo.GetArticleMetadataByPage(limit, offset)
	if err != nil {
		return nil, bizerrors.NewBizError(
			bizerrors.CodeMetaDataFetchFailed,
			"获取文章元数据失败: "+err.Error(),
		)
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

	totalCount, err := s.repo.CountAllArticleMetadata()
	if err != nil {
		return nil, bizerrors.NewBizError(
			bizerrors.CodeMetaDataFetchFailed,
			"统计文章元数据总数失败: "+err.Error(),
		)
	}

	return &dto.ArticleMetadataListResponse{
		Metadata: response,
		Total:    int(totalCount),
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}
