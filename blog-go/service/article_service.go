package service

import (
	"context"
	"errors"
	"strconv"

	"blog-go/constants"
	"blog-go/database"
	"blog-go/model"
	"blog-go/utils"

	"gorm.io/gorm"
)

type ArticleService struct{}

func (s *ArticleService) GetHotArticleList() ([]model.HotArticleVO, error) {
	var articles []model.Article
	err := database.DB.Where("status = ? AND del_flag = ?", constants.ArticleStatusNormal, 0).
		Order("view_count DESC").
		Limit(5).
		Find(&articles).Error
	if err != nil {
		return nil, err
	}

	var vos []model.HotArticleVO
	for _, article := range articles {
		vos = append(vos, model.HotArticleVO{
			ID:        article.ID,
			Title:     article.Title,
			ViewCount: article.ViewCount,
		})
	}
	return vos, nil
}

func (s *ArticleService) GetArticleList(query *model.ArticleQueryDTO) (*model.PageResult, error) {
	var articles []model.Article
	var total int64

	db := database.DB.Model(&model.Article{}).Where("status = ? AND del_flag = ?", constants.ArticleStatusNormal, 0)

	if query.CategoryID > 0 {
		db = db.Where("category_id = ?", query.CategoryID)
	}

	if query.TagID > 0 {
		var articleTags []model.ArticleTag
		database.DB.Where("tag_id = ?", query.TagID).Find(&articleTags)
		if len(articleTags) == 0 {
			return model.NewPageResult([]model.ArticleVO{}, 0, query.PageNum, query.PageSize), nil
		}
		articleIDs := make([]uint64, len(articleTags))
		for i, at := range articleTags {
			articleIDs[i] = at.ArticleID
		}
		db = db.Where("id IN ?", articleIDs)
	}

	if query.YearMonth != "" {
		db = db.Where("DATE_FORMAT(create_time, '%Y/%c') = ?", query.YearMonth)
	}

	if query.Title != "" {
		db = db.Where("title LIKE ?", "%"+query.Title+"%")
	}

	db.Count(&total)

	offset := (query.PageNum - 1) * query.PageSize
	err := db.Order("is_top DESC, create_time DESC").
		Offset(offset).
		Limit(query.PageSize).
		Find(&articles).Error
	if err != nil {
		return nil, err
	}

	vos := s.toVOList(articles)
	return model.NewPageResult(vos, total, query.PageNum, query.PageSize), nil
}

func (s *ArticleService) GetArticleDetail(id uint64) (*model.ArticleVO, error) {
	var article model.Article
	if err := database.DB.First(&article, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.ArticleNotExist.Msg)
		}
		return nil, err
	}

	vo := s.toVO(&article)

	var category model.Category
	if err := database.DB.First(&category, article.CategoryID).Error; err == nil {
		vo.CategoryName = category.Name
	}

	var articleTags []model.ArticleTag
	database.DB.Where("article_id = ?", id).Find(&articleTags)
	if len(articleTags) > 0 {
		tagIDs := make([]uint64, len(articleTags))
		for i, at := range articleTags {
			tagIDs[i] = at.TagID
		}
		var tags []model.Tag
		database.DB.Where("id IN ?", tagIDs).Find(&tags)
		for _, tag := range tags {
			vo.Tags = append(vo.Tags, model.TagVO{ID: tag.ID, Name: tag.Name})
		}
	}

	return vo, nil
}

func (s *ArticleService) AddArticle(dto *model.ArticleDTO, userID uint64) (uint64, error) {
	categoryService := &CategoryService{}
	category, err := categoryService.GetOrAddByName(dto.Category)
	if err != nil {
		return 0, err
	}

	status := constants.ArticleStatusNormal
	if dto.IsDraft {
		status = constants.ArticleStatusDraft
	}

	article := &model.Article{
		Title:      dto.Title,
		Content:    dto.Content,
		Summary:    dto.Summary,
		CategoryID: category.ID,
		Thumbnail:  dto.Thumbnail,
		IsTop:      dto.IsTop,
		Status:     status,
		IsComment:  dto.IsComment,
		CreateBy:   userID,
	}

	if err := database.DB.Create(article).Error; err != nil {
		return 0, err
	}

	if len(dto.Tags) > 0 {
		tagService := &TagService{}
		for _, tagName := range dto.Tags {
			tag, err := tagService.GetOrAddByName(tagName)
			if err != nil {
				continue
			}
			articleTag := &model.ArticleTag{
				ArticleID: article.ID,
				TagID:     tag.ID,
			}
			database.DB.Create(articleTag)
		}
	}

	return article.ID, nil
}

func (s *ArticleService) UpdateArticle(dto *model.ArticleDTO, userID uint64) error {
	var article model.Article
	if err := database.DB.First(&article, dto.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(constants.ArticleNotExist.Msg)
		}
		return err
	}

	categoryService := &CategoryService{}
	category, err := categoryService.GetOrAddByName(dto.Category)
	if err != nil {
		return err
	}

	status := constants.ArticleStatusNormal
	if dto.IsDraft {
		status = constants.ArticleStatusDraft
	}

	updates := map[string]interface{}{
		"title":       dto.Title,
		"content":     dto.Content,
		"summary":     dto.Summary,
		"category_id": category.ID,
		"thumbnail":   dto.Thumbnail,
		"is_top":      dto.IsTop,
		"status":      status,
		"is_comment":  dto.IsComment,
		"update_by":   userID,
	}

	if err := database.DB.Model(&article).Updates(updates).Error; err != nil {
		return err
	}

	database.DB.Where("article_id = ?", dto.ID).Delete(&model.ArticleTag{})

	if len(dto.Tags) > 0 {
		tagService := &TagService{}
		for _, tagName := range dto.Tags {
			tag, err := tagService.GetOrAddByName(tagName)
			if err != nil {
				continue
			}
			articleTag := &model.ArticleTag{
				ArticleID: dto.ID,
				TagID:     tag.ID,
			}
			database.DB.Create(articleTag)
		}
	}

	return nil
}

func (s *ArticleService) DeleteArticle(id uint64) error {
	if err := database.DB.Model(&model.Article{}).Where("id = ?", id).Update("del_flag", 1).Error; err != nil {
		return err
	}
	database.DB.Where("article_id = ?", id).Delete(&model.ArticleTag{})
	return nil
}

func (s *ArticleService) GetArticleCount() (int64, error) {
	var count int64
	err := database.DB.Model(&model.Article{}).
		Where("status = ? AND del_flag = ?", constants.ArticleStatusNormal, 0).
		Count(&count).Error
	return count, err
}

func (s *ArticleService) UpdateViewCount(id uint64) error {
	ctx := context.Background()
	return utils.IncreaseCacheMapValue(ctx, constants.RedisArticleViewKey, strconv.FormatUint(id, 10), 1)
}

func (s *ArticleService) GetArchiveList(pageNum, pageSize int) ([]model.ArchiveVO, error) {
	var results []model.ArchiveVO
	offset := (pageNum - 1) * pageSize
	err := database.DB.Model(&model.Article{}).
		Select("DATE_FORMAT(create_time, '%Y/%c') as date, COUNT(*) as count").
		Where("status = ? AND del_flag = ?", constants.ArticleStatusNormal, 0).
		Group("date").
		Order("date DESC").
		Offset(offset).
		Limit(pageSize).
		Scan(&results).Error
	return results, err
}

func (s *ArticleService) toVO(article *model.Article) *model.ArticleVO {
	return &model.ArticleVO{
		ID:         article.ID,
		Title:      article.Title,
		Content:    article.Content,
		Summary:    article.Summary,
		CategoryID: article.CategoryID,
		Thumbnail:  article.Thumbnail,
		IsTop:      article.IsTop,
		Status:     article.Status,
		ViewCount:  article.ViewCount,
		IsComment:  article.IsComment,
		CreateTime: article.CreateTime,
		UpdateTime: article.UpdateTime,
	}
}

func (s *ArticleService) toVOList(articles []model.Article) []model.ArticleVO {
	var vos []model.ArticleVO
	for _, article := range articles {
		vo := s.toVO(&article)
		var category model.Category
		if err := database.DB.First(&category, article.CategoryID).Error; err == nil {
			vo.CategoryName = category.Name
		}
		vos = append(vos, *vo)
	}
	return vos
}
