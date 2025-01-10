package controller

import (
	"github.com/gin-gonic/gin"
	"mygo/internal/dto"
	"mygo/internal/service"
	"mygo/internal/util"
	"mygo/pkg/bizerrors"
	"mygo/pkg/response" // 统一响应包
)

type ArticleController struct {
	articleService service.ArticleService
}

func NewArticleController(articleService service.ArticleService) *ArticleController {
	return &ArticleController{articleService: articleService}
}

func (c *ArticleController) CreateArticle(ctx *gin.Context) {
	var req dto.CreateArticleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// 参数绑定失败 => 可以直接返回 BizError(CodeInvalidParams)
		response.NewResponse(ctx, bizerrors.NewBizError(bizerrors.CodeInvalidParams, "请求参数错误: "+err.Error()), nil)
		return
	}

	err := c.articleService.CreateArticle(req)
	response.NewResponse(ctx, err, gin.H{"msg": "文章创建成功"})
}

func (c *ArticleController) UpdateArticle(ctx *gin.Context) {
	idStr := ctx.Param("id")
	articleID, parseErr := util.StringToUint(idStr)
	if parseErr != nil {
		response.NewResponse(ctx, bizerrors.NewBizError(bizerrors.CodeInvalidParams, "无效的文章ID"), nil)
		return
	}

	var req dto.UpdateArticleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewResponse(ctx, bizerrors.NewBizError(bizerrors.CodeInvalidParams, "请求参数错误: "+err.Error()), nil)
		return
	}
	req.ID = articleID

	err := c.articleService.UpdateArticle(req)
	response.NewResponse(ctx, err, gin.H{"msg": "文章更新成功"})
}

func (c *ArticleController) DeleteArticle(ctx *gin.Context) {
	idStr := ctx.Param("id")
	articleID, parseErr := util.StringToUint(idStr)
	if parseErr != nil {
		response.NewResponse(ctx, bizerrors.NewBizError(bizerrors.CodeInvalidParams, "无效的文章ID"), nil)
		return
	}

	req := dto.DeleteArticleRequest{ID: articleID}
	err := c.articleService.DeleteArticle(req)
	response.NewResponse(ctx, err, gin.H{"msg": "文章删除成功"})
}

func (c *ArticleController) GetArticleByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	articleID, parseErr := util.StringToUint(idStr)
	if parseErr != nil {
		response.NewResponse(ctx, bizerrors.NewBizError(bizerrors.CodeInvalidParams, "无效的文章ID"), nil)
		return
	}

	article, err := c.articleService.GetArticleByID(articleID)
	response.NewResponse(ctx, err, gin.H{
		"msg":     "文章获取成功",
		"article": article,
	})
}

func (c *ArticleController) GetArticlesByPage(ctx *gin.Context) {
	var req dto.GetArticlesByPageRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.NewResponse(ctx, bizerrors.NewBizError(bizerrors.CodeInvalidParams, "请求参数错误: "+err.Error()), nil)
		return
	}

	articles, err := c.articleService.GetArticlesByPage(req)
	response.NewResponse(ctx, err, gin.H{
		"msg":      "获取文章列表成功",
		"articles": articles,
	})
}
