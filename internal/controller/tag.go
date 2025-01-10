package controller

import (
	"mygo/internal/dto"
	"mygo/internal/service"
	"mygo/pkg/err_code"
	"mygo/pkg/response"

	"github.com/gin-gonic/gin"
)

type TagController struct {
	tagService service.TagService
}

func NewTagController(tagService service.TagService) *TagController {
	return &TagController{tagService: tagService}
}

func (c *TagController) CreateTag(ctx *gin.Context) {
	var req dto.CreateTagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewResponse(ctx, err_code.InvalidParams.WithDetails(err.Error()), nil)
		return
	}

	if err := c.tagService.CreateTag(req); err != nil {
		response.NewResponse(ctx, err_code.TagCreateFailed.WithDetails(err.Error()), nil)
		return
	}

	response.NewResponse(ctx, err_code.Success, nil)
}

func (c *TagController) GetAllTags(ctx *gin.Context) {
	tags, err := c.tagService.GetAllTags()
	if err != nil {
		response.NewResponse(ctx, err_code.ServerError.WithDetails(err.Error()), nil)
		return
	}

	response.NewResponse(ctx, err_code.Success, tags)
}
