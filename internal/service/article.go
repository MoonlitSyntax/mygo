package service

import (
	"fmt"
	"mygo/internal/dto"
	"mygo/internal/repository"
)

type ArticleService interface {
	CreateArticle(req dto.CreateArticleRequest) error
	UpdateArticle(req dto.UpdateArticleRequest) error
	DeleteArticle(req dto.DeleteArticleRequest) error
	GetArticleByID(id uint) (*dto.ArticleResponse, error)
	GetArticlesByPage(req dto.GetArticlesByPageRequest) (*dto.ArticleListResponse, error)
	CountArticles() (int64, error)
}

type articleService struct {
	repo repository.ArticleRepository
}

func NewArticleService(repo repository.ArticleRepository) ArticleService {
	return &articleService{repo: repo}
}

func (s *articleService) CreateArticle(req dto.CreateArticleRequest) error {
	return s.repo.CreateArticle(
		req.Title,
		req.Content,
		req.Slug,
		req.Description,
		req.Top,
		req.BelongTo,
		req.UserID,
		req.CategoryID,
		req.TagsID,
	)
}

func (s *articleService) UpdateArticle(req dto.UpdateArticleRequest) error {
	if req.Updates == nil || len(req.Updates) == 0 {
		return fmt.Errorf("没有要更新的字段")
	}
	return s.repo.UpdateArticle(req.ID, req.Updates, req.TagsID)
}

func (s *articleService) DeleteArticle(req dto.DeleteArticleRequest) error {
	if req.ID == 0 {
		return fmt.Errorf("无效的 ID")
	}
	return s.repo.DeleteArticle(req.ID)
}

func (s *articleService) GetArticleByID(id uint) (*dto.ArticleResponse, error) {
	article, err := s.repo.GetArticleByID(id)
	if err != nil {
		return nil, err
	}

	// 转换为 DTO
	tags := make([]dto.TagResponse, len(article.Tags))
	for i, tag := range article.Tags {
		tags[i] = dto.TagResponse{
			ID:   tag.ID,
			Name: tag.Name,
		}
	}

	return &dto.ArticleResponse{
		ID:          article.ID,
		Title:       article.Title,
		Content:     article.Content,
		Slug:        article.Slug,
		Description: article.Description,
		Top:         article.Top,
		BelongTo:    article.BelongTo,
		UserID:      article.UserID,
		CategoryID:  article.CategoryID,
		Tags:        tags,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
	}, nil
}

func (s *articleService) GetArticlesByPage(req dto.GetArticlesByPageRequest) (*dto.ArticleListResponse, error) {
	limit := req.PageSize
	offset := (req.Page - 1) * req.PageSize

	articles, err := s.repo.GetArticlesByPage(limit, offset)
	if err != nil {
		return nil, err
	}

	// 转换为 DTO
	var articleResponses []dto.ArticleResponse
	for _, article := range articles {
		tags := make([]dto.TagResponse, len(article.Tags))
		for i, tag := range article.Tags {
			tags[i] = dto.TagResponse{
				ID:   tag.ID,
				Name: tag.Name,
			}
		}

		articleResponses = append(articleResponses, dto.ArticleResponse{
			ID:          article.ID,
			Title:       article.Title,
			Content:     article.Content,
			Slug:        article.Slug,
			Description: article.Description,
			Top:         article.Top,
			BelongTo:    article.BelongTo,
			UserID:      article.UserID,
			CategoryID:  article.CategoryID,
			Tags:        tags,
			CreatedAt:   article.CreatedAt,
			UpdatedAt:   article.UpdatedAt,
		})
	}

	total, err := s.repo.CountArticle()
	if err != nil {
		return nil, err
	}

	return &dto.ArticleListResponse{
		Articles: articleResponses,
		Total:    int(total),
	}, nil
}

func (s *articleService) CountArticles() (int64, error) {
	return s.repo.CountArticle()
}
