package controller

import (
	"strconv"

	"blog-go/middleware"
	"blog-go/model"
	"blog-go/response"
	"blog-go/service"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	articleService *service.ArticleService
}

func NewArticleController() *ArticleController {
	return &ArticleController{
		articleService: &service.ArticleService{},
	}
}

func (c *ArticleController) GetHotArticleList(ctx *gin.Context) {
	result, err := c.articleService.GetHotArticleList()
	if err != nil {
		response.Error(ctx, 500, "获取热门文章失败")
		return
	}
	response.Success(ctx, result)
}

func (c *ArticleController) GetArticleList(ctx *gin.Context) {
	var query model.ArticleQueryDTO
	if err := ctx.ShouldBindQuery(&query); err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	if query.PageNum == 0 {
		query.PageNum = 1
	}
	if query.PageSize == 0 {
		query.PageSize = 10
	}

	result, err := c.articleService.GetArticleList(&query)
	if err != nil {
		response.Error(ctx, 500, "获取文章列表失败")
		return
	}
	response.Success(ctx, result)
}

func (c *ArticleController) GetArticleDetail(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	c.articleService.UpdateViewCount(id)

	result, err := c.articleService.GetArticleDetail(id)
	if err != nil {
		response.Error(ctx, 500, err.Error())
		return
	}
	response.Success(ctx, result)
}

func (c *ArticleController) AddArticle(ctx *gin.Context) {
	var dto model.ArticleDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	userID := middleware.GetUserID(ctx)
	id, err := c.articleService.AddArticle(&dto, userID)
	if err != nil {
		response.Error(ctx, 500, "添加文章失败")
		return
	}
	response.Success(ctx, id)
}

func (c *ArticleController) UpdateArticle(ctx *gin.Context) {
	var dto model.ArticleDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	userID := middleware.GetUserID(ctx)
	if err := c.articleService.UpdateArticle(&dto, userID); err != nil {
		response.Error(ctx, 500, err.Error())
		return
	}
	response.Success(ctx, nil)
}

func (c *ArticleController) DeleteArticle(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.BadRequest(ctx, "参数错误")
		return
	}

	if err := c.articleService.DeleteArticle(id); err != nil {
		response.Error(ctx, 500, "删除文章失败")
		return
	}
	response.Success(ctx, nil)
}

func (c *ArticleController) GetArticleCount(ctx *gin.Context) {
	count, err := c.articleService.GetArticleCount()
	if err != nil {
		response.Error(ctx, 500, "获取文章数量失败")
		return
	}
	response.Success(ctx, count)
}

func (c *ArticleController) GetArchiveList(ctx *gin.Context) {
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	result, err := c.articleService.GetArchiveList(pageNum, pageSize)
	if err != nil {
		response.Error(ctx, 500, "获取归档列表失败")
		return
	}
	response.Success(ctx, result)
}
