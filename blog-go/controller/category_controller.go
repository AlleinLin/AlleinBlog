package controller

import (
	"strconv"

	"blog-go/response"
	"blog-go/service"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService *service.CategoryService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		categoryService: &service.CategoryService{},
	}
}

func (c *CategoryController) GetCategoryList(ctx *gin.Context) {
	result, err := c.categoryService.GetCategoryList()
	if err != nil {
		response.Error(ctx, 500, "获取分类列表失败")
		return
	}
	response.Success(ctx, result)
}

func (c *CategoryController) GetCategoryByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	result, err := c.categoryService.GetCategoryByID(id)
	if err != nil {
		response.Error(ctx, 500, err.Error())
		return
	}
	response.Success(ctx, result)
}
