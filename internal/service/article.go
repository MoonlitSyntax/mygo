package service

import (
	"mygo/internal/dto"
	"mygo/internal/repository"

	"mygo/pkg/bizerrors"
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
	err := s.repo.CreateArticle(
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
	if err != nil {
		// 关键：包装为 BizError，带上具体错误码 & 信息
		return bizerrors.NewBizError(bizerrors.CodeArticleCreateFailed, "文章创建失败: "+err.Error())
	}
	return nil
}

func (s *articleService) UpdateArticle(req dto.UpdateArticleRequest) error {
	if req.Updates == nil || len(req.Updates) == 0 {
		// 这里说明参数无效
		return bizerrors.NewBizError(bizerrors.CodeInvalidParams, "没有要更新的字段")
	}
	err := s.repo.UpdateArticle(req.ID, req.Updates, req.TagsID)
	if err != nil {
		return bizerrors.NewBizError(bizerrors.CodeArticleUpdateFailed, "文章更新失败: "+err.Error())
	}
	return nil
}

func (s *articleService) DeleteArticle(req dto.DeleteArticleRequest) error {
	if req.ID == 0 {
		// 参数错误
		return bizerrors.NewBizError(bizerrors.CodeInvalidParams, "无效的 ID")
	}
	err := s.repo.DeleteArticle(req.ID)
	if err != nil {
		return bizerrors.NewBizError(bizerrors.CodeArticleDeleteFailed, "文章删除失败: "+err.Error())
	}
	return nil
}

func (s *articleService) GetArticleByID(id uint) (*dto.ArticleResponse, error) {
	article, err := s.repo.GetArticleByID(id)
	if err != nil {
		// 这里可能是 gorm.ErrRecordNotFound 等
		return nil, bizerrors.NewBizError(bizerrors.CodeArticleNotFound, "文章不存在: "+err.Error())
	}

	// 转为 DTO
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
		return nil, bizerrors.NewBizError(
			bizerrors.CodeServerError,
			"分页查询文章失败: "+err.Error(),
		)
	}

	// 转换为 DTO
	var articleResponses []dto.ArticleResponse
	for _, article := range articles {
		// 构建 TagResponse 列表
		tags := make([]dto.TagResponse, len(article.Tags))
		for i, t := range article.Tags {
			tags[i] = dto.TagResponse{
				ID:   t.ID,
				Name: t.Name,
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

	// 获取总数
	total, err := s.repo.CountArticle()
	if err != nil {
		return nil, bizerrors.NewBizError(
			bizerrors.CodeServerError,
			"统计文章总数失败: "+err.Error(),
		)
	}

	// 组装返回
	return &dto.ArticleListResponse{
		Articles: articleResponses,
		Total:    int(total),
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

func (s *articleService) CountArticles() (int64, error) {
	count, err := s.repo.CountArticle()
	if err != nil {
		return 0, bizerrors.NewBizError(bizerrors.CodeServerError, "统计文章失败: "+err.Error())
	}
	return count, nil
}
