package controller

import (
	"github.com/gin-gonic/gin"
	"mygo/internal/dto"
	"mygo/internal/service"
	"mygo/internal/util"
	"mygo/pkg/bizerrors"
	"mygo/pkg/response"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var req dto.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewResponse(ctx, bizerrors.NewBizError(bizerrors.CodeInvalidParams, "请求参数错误: "+err.Error()), nil)
		return
	}

	err := c.userService.CreateUser(req)
	response.NewResponse(ctx, err, nil)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	userID, parseErr := util.StringToUint(idStr)
	if parseErr != nil {
		response.NewResponse(ctx, bizerrors.NewBizError(bizerrors.CodeInvalidParams, "Invalid user ID"), nil)
		return
	}

	var req dto.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewResponse(ctx, bizerrors.NewBizError(bizerrors.CodeInvalidParams, "请求参数错误: "+err.Error()), nil)
		return
	}

	err := c.userService.UpdateUser(req, userID)
	response.NewResponse(ctx, err, nil)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	userID, parseErr := util.StringToUint(idStr)
	if parseErr != nil {
		response.NewResponse(ctx, bizerrors.NewBizError(bizerrors.CodeInvalidParams, "Invalid user ID"), nil)
		return
	}

	req := dto.DeleteUserRequest{ID: userID}
	err := c.userService.DeleteUser(req)
	response.NewResponse(ctx, err, nil)
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	userID, parseErr := util.StringToUint(idStr)
	if parseErr != nil {
		response.NewResponse(ctx, bizerrors.NewBizError(bizerrors.CodeInvalidParams, "Invalid user ID"), nil)
		return
	}

	user, err := c.userService.GetUserByID(userID)
	response.NewResponse(ctx, err, user)
}
