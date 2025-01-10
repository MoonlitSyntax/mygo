package controller

import (
	"mygo/internal/dto"
	"mygo/pkg/bizerrors"
	"strconv"

	"github.com/gin-gonic/gin"
	"mygo/internal/service"
	"mygo/pkg/response"
)

type MetaDataController struct {
	metaDataService service.MetaDataService
}

func NewMetaDataController(metaDataService service.MetaDataService) *MetaDataController {
	return &MetaDataController{metaDataService: metaDataService}
}

func (c *MetaDataController) GetAllArticleMetadata(ctx *gin.Context) {
	data, err := c.metaDataService.GetAllArticleMetadata()
	response.NewResponse(ctx, err, data)
}

func (c *MetaDataController) GetArticleMetadataByCategory(ctx *gin.Context) {
	categoryIDStr := ctx.Param("category_id")
	categoryID, parseErr := strconv.ParseUint(categoryIDStr, 10, 64)
	if parseErr != nil {
		response.NewResponse(ctx, bizerrors.NewBizError(bizerrors.CodeInvalidParams, "Invalid category ID"), nil)
		return
	}

	data, err := c.metaDataService.GetArticleMetadataByCategory(uint(categoryID))
	// 如果想特判“为空”也可在 Service 层返回特定错误。
	response.NewResponse(ctx, err, data)
}

func (c *MetaDataController) GetArticleMetadataByTag(ctx *gin.Context) {
	tagIDStr := ctx.Param("tag_id")
	tagID, parseErr := strconv.ParseUint(tagIDStr, 10, 64)
	if parseErr != nil {
		response.NewResponse(ctx, bizerrors.NewBizError(bizerrors.CodeInvalidParams, "Invalid tag ID"), nil)
		return
	}

	data, err := c.metaDataService.GetArticleMetadataByTag(uint(tagID))
	response.NewResponse(ctx, err, data)
}

func (c *MetaDataController) GetArticleMetadataByPage(ctx *gin.Context) {

	var req dto.GetArticleMetadataPageRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.NewResponse(ctx, bizerrors.NewBizError(bizerrors.CodeInvalidParams, "分页参数错误: "+err.Error()), nil)
		return
	}

	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}

	data, err := c.metaDataService.GetArticleMetadataByPage(req)
	response.NewResponse(ctx, err, data)
}
