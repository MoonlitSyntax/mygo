package controller

import (
	"github.com/gin-gonic/gin"
	"mygo/internal/dto"
	"mygo/internal/service"
	"mygo/internal/util"
	"mygo/pkg/bizerrors"
	"mygo/pkg/response"
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
		response.NewResponse(ctx, bizerrors.NewBizError(bizerrors.CodeInvalidParams, "请求参数错误: "+err.Error()), nil)
		return
	}

	err := c.categoryService.CreateCategory(req)
	response.NewResponse(ctx, err, nil)
}

func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	idStr := ctx.Param("id")
	categoryID, parseErr := util.StringToUint(idStr)
	if parseErr != nil {
		response.NewResponse(ctx, bizerrors.NewBizError(bizerrors.CodeInvalidParams, "Invalid category ID"), nil)
		return
	}

	var req dto.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewResponse(ctx, bizerrors.NewBizError(bizerrors.CodeInvalidParams, "请求参数错误: "+err.Error()), nil)
		return
	}

	err := c.categoryService.UpdateCategory(req, categoryID)
	response.NewResponse(ctx, err, nil)
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	idStr := ctx.Param("id")
	categoryID, parseErr := util.StringToUint(idStr)
	if parseErr != nil {
		response.NewResponse(ctx, bizerrors.NewBizError(bizerrors.CodeInvalidParams, "Invalid category ID"), nil)
		return
	}

	req := dto.DeleteCategoryRequest{ID: categoryID}
	err := c.categoryService.DeleteCategory(req)
	response.NewResponse(ctx, err, nil)
}

func (c *CategoryController) GetAllCategories(ctx *gin.Context) {
	categories, err := c.categoryService.GetAllCategories()
	response.NewResponse(ctx, err, categories)
}
