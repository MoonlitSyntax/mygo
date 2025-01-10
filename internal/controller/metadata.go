package controller

import (
	"mygo/internal/service"
	"mygo/pkg/err_code"
	"mygo/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MetaDataController struct {
	metaDataService service.MetaDataService
}

func NewMetaDataController(metaDataService service.MetaDataService) *MetaDataController {
	return &MetaDataController{metaDataService: metaDataService}
}

// GetAllArticleMetadata 获取所有文章元数据
func (c *MetaDataController) GetAllArticleMetadata(ctx *gin.Context) {
	data, err := c.metaDataService.GetAllArticleMetadata()
	if err != nil {
		response.NewResponse(ctx, err_code.MetaDataFetchFailed.WithDetails(err.Error()), nil)
		return
	}

	response.NewResponse(ctx, err_code.Success, data)
}

// GetArticleMetadataByCategory 根据分类获取文章元数据
func (c *MetaDataController) GetArticleMetadataByCategory(ctx *gin.Context) {
	categoryID, err := strconv.ParseUint(ctx.Param("category_id"), 10, 64)
	if err != nil {
		response.NewResponse(ctx, err_code.InvalidParams.WithDetails("Invalid category ID"), nil)
		return
	}

	data, err := c.metaDataService.GetArticleMetadataByCategory(uint(categoryID))
	if err != nil {
		response.NewResponse(ctx, err_code.MetaDataFetchFailed.WithDetails(err.Error()), nil)
		return
	}

	if len(data.Metadata) == 0 {
		response.NewResponse(ctx, err_code.MetaDataNotFound, nil)
		return
	}

	response.NewResponse(ctx, err_code.Success, data)
}

// GetArticleMetadataByTag 根据标签获取文章元数据
func (c *MetaDataController) GetArticleMetadataByTag(ctx *gin.Context) {
	tagID, err := strconv.ParseUint(ctx.Param("tag_id"), 10, 64)
	if err != nil {
		response.NewResponse(ctx, err_code.InvalidParams.WithDetails("Invalid tag ID"), nil)
		return
	}

	data, err := c.metaDataService.GetArticleMetadataByTag(uint(tagID))
	if err != nil {
		response.NewResponse(ctx, err_code.MetaDataFetchFailed.WithDetails(err.Error()), nil)
		return
	}

	if len(data.Metadata) == 0 {
		response.NewResponse(ctx, err_code.MetaDataNotFound, nil)
		return
	}

	response.NewResponse(ctx, err_code.Success, data)
}
