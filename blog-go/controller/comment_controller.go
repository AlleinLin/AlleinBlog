package controller

import (
	"strconv"

	"blog-go/middleware"
	"blog-go/model"
	"blog-go/response"
	"blog-go/service"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService *service.CommentService
}

func NewCommentController() *CommentController {
	return &CommentController{
		commentService: &service.CommentService{},
	}
}

func (c *CommentController) GetCommentList(ctx *gin.Context) {
	articleIDStr := ctx.Query("articleId")
	articleID, err := strconv.ParseUint(articleIDStr, 10, 64)
	if err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	result, err := c.commentService.GetCommentList(articleID, pageNum, pageSize)
	if err != nil {
		response.Error(ctx, 500, "获取评论列表失败")
		return
	}
	response.Success(ctx, result)
}

func (c *CommentController) AddComment(ctx *gin.Context) {
	var dto model.CommentDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	userID := middleware.GetUserID(ctx)
	if err := c.commentService.AddComment(&dto, userID); err != nil {
		response.Error(ctx, 500, "添加评论失败")
		return
	}
	response.Success(ctx, nil)
}

func (c *CommentController) UpdateComment(ctx *gin.Context) {
	var dto model.CommentDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	userID := middleware.GetUserID(ctx)
	if err := c.commentService.UpdateComment(&dto, userID); err != nil {
		response.Error(ctx, 500, err.Error())
		return
	}
	response.Success(ctx, nil)
}

func (c *CommentController) DeleteComment(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	userID := middleware.GetUserID(ctx)
	if err := c.commentService.DeleteComment(id, userID); err != nil {
		response.Error(ctx, 500, err.Error())
		return
	}
	response.Success(ctx, nil)
}
