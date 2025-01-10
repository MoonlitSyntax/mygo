package controller

import (
	"github.com/gin-gonic/gin"
	"mygo/internal/dto"
	"mygo/internal/service"
	"mygo/internal/util"
	"mygo/pkg/bizerrors"
	"mygo/pkg/response"
)

// TagController 控制器
type TagController struct {
	tagService service.TagService
}

// NewTagController 构造函数
func NewTagController(tagService service.TagService) *TagController {
	return &TagController{tagService: tagService}
}

// CreateTag 创建标签
func (c *TagController) CreateTag(ctx *gin.Context) {
	var req dto.CreateTagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewResponse(ctx, bizerrors.NewBizError(bizerrors.CodeInvalidParams, "请求参数错误: "+err.Error()), nil)
		return
	}

	err := c.tagService.CreateTag(req)
	response.NewResponse(ctx, err, gin.H{"msg": "标签创建成功"})
}

func (c *TagController) UpdateTag(ctx *gin.Context) {
	idStr := ctx.Param("id")
	tagID, parseErr := util.StringToUint(idStr)
	if parseErr != nil {
		response.NewResponse(ctx, bizerrors.NewBizError(bizerrors.CodeInvalidParams, "无效的标签ID"), nil)
		return
	}

	var req dto.UpdateTagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewResponse(ctx, bizerrors.NewBizError(bizerrors.CodeInvalidParams, "请求参数错误: "+err.Error()), nil)
		return
	}

	err := c.tagService.UpdateTag(req, tagID)
	response.NewResponse(ctx, err, gin.H{"msg": "标签更新成功"})
}

func (c *TagController) DeleteTag(ctx *gin.Context) {
	idStr := ctx.Param("id")
	tagID, parseErr := util.StringToUint(idStr)
	if parseErr != nil {
		response.NewResponse(ctx, bizerrors.NewBizError(bizerrors.CodeInvalidParams, "无效的标签ID"), nil)
		return
	}

	req := dto.DeleteTagRequest{ID: tagID}
	// 调用 Service
	err := c.tagService.DeleteTag(req)
	response.NewResponse(ctx, err, gin.H{"msg": "标签删除成功"})
}

func (c *TagController) GetTagByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	tagID, parseErr := util.StringToUint(idStr)
	if parseErr != nil {
		response.NewResponse(ctx, bizerrors.NewBizError(bizerrors.CodeInvalidParams, "无效的标签ID"), nil)
		return
	}

	tagInfo, err := c.tagService.GetTagByID(tagID)
	response.NewResponse(ctx, err, gin.H{
		"msg": "标签获取成功",
		"tag": tagInfo,
	})
}

func (c *TagController) GetAllTags(ctx *gin.Context) {
	tags, err := c.tagService.GetAllTags()
	response.NewResponse(ctx, err, gin.H{
		"msg":  "获取所有标签成功",
		"tags": tags,
	})
}
