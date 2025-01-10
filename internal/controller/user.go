package controller

import (
	"mygo/internal/dto"
	"mygo/internal/service"
	"mygo/internal/util"
	"mygo/pkg/err_code"
	"mygo/pkg/response"

	"github.com/gin-gonic/gin"
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
		response.NewResponse(ctx, err_code.InvalidParams.WithDetails(err.Error()), nil)
		return
	}

	if err := c.userService.CreateUser(req); err != nil {
		response.NewResponse(ctx, err_code.UserCreateFailed.WithDetails(err.Error()), nil)
		return
	}

	response.NewResponse(ctx, err_code.Success, nil)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := util.StringToUint(id)
	if err != nil {
		response.NewResponse(ctx, err_code.InvalidParams.WithDetails("Invalid user ID"), nil)
		return
	}

	var req dto.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewResponse(ctx, err_code.InvalidParams.WithDetails(err.Error()), nil)
		return
	}

	if err := c.userService.UpdateUser(req, userID); err != nil {
		response.NewResponse(ctx, err_code.ServerError.WithDetails(err.Error()), nil)
		return
	}

	response.NewResponse(ctx, err_code.Success, nil)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := util.StringToUint(id)
	if err != nil {
		response.NewResponse(ctx, err_code.InvalidParams.WithDetails("Invalid user ID"), nil)
		return
	}

	req := dto.DeleteUserRequest{ID: userID}
	if err := c.userService.DeleteUser(req); err != nil {
		response.NewResponse(ctx, err_code.ServerError.WithDetails(err.Error()), nil)
		return
	}

	response.NewResponse(ctx, err_code.Success, nil)
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := util.StringToUint(id)
	if err != nil {
		response.NewResponse(ctx, err_code.InvalidParams.WithDetails("Invalid user ID"), nil)
		return
	}

	user, err := c.userService.GetUserByID(userID)
	if err != nil {
		response.NewResponse(ctx, err_code.NotFound.WithDetails(err.Error()), nil)
		return
	}

	response.NewResponse(ctx, err_code.Success, user)
}
