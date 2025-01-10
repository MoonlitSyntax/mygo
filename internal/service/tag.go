package service

import (
	"mygo/internal/dto"
	"mygo/internal/repository"
	"mygo/pkg/bizerrors"
)

type TagService interface {
	CreateTag(req dto.CreateTagRequest) error
	UpdateTag(req dto.UpdateTagRequest, id uint) error
	DeleteTag(req dto.DeleteTagRequest) error
	GetTagByID(id uint) (*dto.TagResponse, error)
	GetAllTags() ([]dto.TagResponse, error)
	CountTags() (int64, error)
}

type tagService struct {
	repo repository.TagRepository
}

func NewTagService(repo repository.TagRepository) TagService {
	return &tagService{repo: repo}
}

func (s *tagService) CreateTag(req dto.CreateTagRequest) error {
	err := s.repo.CreateTag(req.Name)
	if err != nil {
		return bizerrors.NewBizError(
			bizerrors.CodeTagCreateFailed,
			"创建标签失败: "+err.Error(),
		)
	}
	return nil
}

func (s *tagService) UpdateTag(req dto.UpdateTagRequest, id uint) error {
	if req.Name == "" {
		// 参数错误
		return bizerrors.NewBizError(
			bizerrors.CodeInvalidParams,
			"标签名称不能为空",
		)
	}
	err := s.repo.UpdateTag(id, req.Name)
	if err != nil {
		return bizerrors.NewBizError(
			bizerrors.CodeTagUpdateFailed,
			"更新标签失败: "+err.Error(),
		)
	}
	return nil
}
func (s *tagService) DeleteTag(req dto.DeleteTagRequest) error {
	if req.ID == 0 {
		// 参数错误
		return bizerrors.NewBizError(
			bizerrors.CodeInvalidParams,
			"无效的标签 ID",
		)
	}
	err := s.repo.DeleteTag(req.ID)
	if err != nil {
		return bizerrors.NewBizError(
			bizerrors.CodeTagDeleteFailed,
			"删除标签失败: "+err.Error(),
		)
	}
	return nil
}

func (s *tagService) GetTagByID(id uint) (*dto.TagResponse, error) {
	tag, err := s.repo.GetTagById(id)
	if err != nil {
		// 可能是 Tag 不存在
		return nil, bizerrors.NewBizError(
			bizerrors.CodeTagNotFound,
			"标签不存在: "+err.Error(),
		)
	}

	return &dto.TagResponse{
		ID:   tag.ID,
		Name: tag.Name,
	}, nil
}

func (s *tagService) GetAllTags() ([]dto.TagResponse, error) {
	tags, err := s.repo.GetAllTags()
	if err != nil {
		return nil, bizerrors.NewBizError(
			bizerrors.CodeServerError,
			"获取所有标签失败: "+err.Error(),
		)
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

func (s *tagService) CountTags() (int64, error) {
	count, err := s.repo.CountTag()
	if err != nil {
		return 0, bizerrors.NewBizError(
			bizerrors.CodeServerError,
			"统计标签失败: "+err.Error(),
		)
	}
	return count, nil
}
