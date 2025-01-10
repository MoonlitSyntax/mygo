package controller

import (
	"mygo/internal/dto"
	"mygo/internal/service"
	"mygo/internal/util"
	"mygo/pkg/err_code"
	"mygo/pkg/response"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryController {
	return &CategoryController{categoryService: categoryService}
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var req dto.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewResponse(ctx, err_code.InvalidParams.WithDetails(err.Error()), nil)
		return
	}

	if err := c.categoryService.CreateCategory(req); err != nil {
		response.NewResponse(ctx, err_code.CategoryCreateFailed.WithDetails(err.Error()), nil)
		return
	}

	response.NewResponse(ctx, err_code.Success, nil)
}

func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	categoryID, err := util.StringToUint(id)
	if err != nil {
		response.NewResponse(ctx, err_code.InvalidParams.WithDetails("Invalid category ID"), nil)
		return
	}

	var req dto.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewResponse(ctx, err_code.InvalidParams.WithDetails(err.Error()), nil)
		return
	}

	if err := c.categoryService.UpdateCategory(req, categoryID); err != nil {
		response.NewResponse(ctx, err_code.CategoryUpdateFailed.WithDetails(err.Error()), nil)
		return
	}

	response.NewResponse(ctx, err_code.Success, nil)
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	categoryID, err := util.StringToUint(id)
	if err != nil {
		response.NewResponse(ctx, err_code.InvalidParams.WithDetails("Invalid category ID"), nil)
		return
	}

	req := dto.DeleteCategoryRequest{ID: categoryID}
	if err := c.categoryService.DeleteCategory(req); err != nil {
		response.NewResponse(ctx, err_code.CategoryDeleteFailed.WithDetails(err.Error()), nil)
		return
	}

	response.NewResponse(ctx, err_code.Success, nil)
}

func (c *CategoryController) GetAllCategories(ctx *gin.Context) {
	categories, err := c.categoryService.GetAllCategories()
	if err != nil {
		response.NewResponse(ctx, err_code.ServerError.WithDetails(err.Error()), nil)
		return
	}

	response.NewResponse(ctx, err_code.Success, categories)
}
