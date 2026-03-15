package controller

import (
	"blog-go/middleware"
	"blog-go/model"
	"blog-go/response"
	"blog-go/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: &service.UserService{},
	}
}

func (c *UserController) Login(ctx *gin.Context) {
	var dto model.LoginDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	result, err := c.userService.Login(&dto)
	if err != nil {
		response.Error(ctx, 500, err.Error())
		return
	}

	response.Success(ctx, result)
}

func (c *UserController) Logout(ctx *gin.Context) {
	userID := middleware.GetUserID(ctx)
	if err := c.userService.Logout(userID); err != nil {
		response.Error(ctx, 500, "登出失败")
		return
	}
	response.Success(ctx, nil)
}

func (c *UserController) Register(ctx *gin.Context) {
	var dto model.RegisterDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	if err := c.userService.Register(&dto); err != nil {
		response.Error(ctx, 500, err.Error())
		return
	}

	response.Success(ctx, nil)
}

func (c *UserController) GetUserInfo(ctx *gin.Context) {
	userID := middleware.GetUserID(ctx)
	userInfo, err := c.userService.GetUserInfo(userID)
	if err != nil {
		response.Error(ctx, 500, "获取用户信息失败")
		return
	}
	response.Success(ctx, userInfo)
}

func (c *UserController) UpdateUserInfo(ctx *gin.Context) {
	userID := middleware.GetUserID(ctx)
	var dto model.UserInfoDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	if err := c.userService.UpdateUserInfo(userID, &dto); err != nil {
		response.Error(ctx, 500, "更新用户信息失败")
		return
	}

	response.Success(ctx, nil)
}

func (c *UserController) GetAdminInfo(ctx *gin.Context) {
	userInfo, err := c.userService.GetAdminInfo()
	if err != nil {
		response.Error(ctx, 500, "获取管理员信息失败")
		return
	}
	response.Success(ctx, userInfo)
}
