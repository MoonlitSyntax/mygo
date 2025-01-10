package controller

import (
	"mygo/internal/dto"
	"mygo/internal/service"
	"mygo/internal/util"
	"mygo/pkg/err_code"
	"mygo/pkg/response"

	"github.com/gin-gonic/gin"
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
		response.NewResponse(ctx, err_code.InvalidParams.WithDetails(err.Error()), nil)
		return
	}

	if err := c.articleService.CreateArticle(req); err != nil {
		response.NewResponse(ctx, err_code.ArticleCreateFailed.WithDetails(err.Error()), nil)
		return
	}

	response.NewResponse(ctx, err_code.Success, nil)
}

func (c *ArticleController) UpdateArticle(ctx *gin.Context) {
	id := ctx.Param("id")
	articleID, err := util.StringToUint(id)
	if err != nil {
		response.NewResponse(ctx, err_code.InvalidParams.WithDetails("Invalid article ID"), nil)
		return
	}

	var req dto.UpdateArticleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewResponse(ctx, err_code.InvalidParams.WithDetails(err.Error()), nil)
		return
	}

	req.ID = articleID
	if err := c.articleService.UpdateArticle(req); err != nil {
		response.NewResponse(ctx, err_code.ArticleUpdateFailed.WithDetails(err.Error()), nil)
		return
	}

	response.NewResponse(ctx, err_code.Success, nil)
}

func (c *ArticleController) DeleteArticle(ctx *gin.Context) {
	id := ctx.Param("id")
	articleID, err := util.StringToUint(id)
	if err != nil {
		response.NewResponse(ctx, err_code.InvalidParams.WithDetails("Invalid article ID"), nil)
		return
	}

	req := dto.DeleteArticleRequest{ID: articleID}
	if err := c.articleService.DeleteArticle(req); err != nil {
		response.NewResponse(ctx, err_code.ArticleDeleteFailed.WithDetails(err.Error()), nil)
		return
	}

	response.NewResponse(ctx, err_code.Success, nil)
}

func (c *ArticleController) GetArticleByID(ctx *gin.Context) {
	id := ctx.Param("id")
	articleID, err := util.StringToUint(id)
	if err != nil {
		response.NewResponse(ctx, err_code.InvalidParams.WithDetails("Invalid article ID"), nil)
		return
	}

	article, err := c.articleService.GetArticleByID(articleID)
	if err != nil {
		response.NewResponse(ctx, err_code.ArticleNotFound.WithDetails(err.Error()), nil)
		return
	}

	response.NewResponse(ctx, err_code.Success, article)
}

func (c *ArticleController) GetArticles(ctx *gin.Context) {
	var req dto.GetArticlesByPageRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.NewResponse(ctx, err_code.InvalidParams.WithDetails(err.Error()), nil)
		return
	}

	articles, err := c.articleService.GetArticlesByPage(req)
	if err != nil {
		response.NewResponse(ctx, err_code.ServerError.WithDetails(err.Error()), nil)
		return
	}

	response.NewResponse(ctx, err_code.Success, articles)
}
