package controller

import (
	"strconv"

	"blog-go/response"
	"blog-go/service"

	"github.com/gin-gonic/gin"
)

type TagController struct {
	tagService *service.TagService
}

func NewTagController() *TagController {
	return &TagController{
		tagService: &service.TagService{},
	}
}

func (c *TagController) GetTagList(ctx *gin.Context) {
	result, err := c.tagService.GetTagList()
	if err != nil {
		response.Error(ctx, 500, "获取标签列表失败")
		return
	}
	response.Success(ctx, result)
}

func (c *TagController) GetTagByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	result, err := c.tagService.GetTagByID(id)
	if err != nil {
		response.Error(ctx, 500, err.Error())
		return
	}
	response.Success(ctx, result)
}
