package service

import (
	"errors"

	"blog-go/constants"
	"blog-go/database"
	"blog-go/model"

	"gorm.io/gorm"
)

type TagService struct{}

func (s *TagService) GetTagList() ([]model.TagVO, error) {
	var tags []model.Tag
	err := database.DB.Where("del_flag = ?", 0).Find(&tags).Error
	if err != nil {
		return nil, err
	}

	var vos []model.TagVO
	for _, tag := range tags {
		vos = append(vos, model.TagVO{
			ID:   tag.ID,
			Name: tag.Name,
		})
	}
	return vos, nil
}

func (s *TagService) GetTagByID(id uint64) (*model.TagVO, error) {
	var tag model.Tag
	if err := database.DB.First(&tag, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.TagNotExist.Msg)
		}
		return nil, err
	}

	return &model.TagVO{
		ID:   tag.ID,
		Name: tag.Name,
	}, nil
}

func (s *TagService) GetOrAddByName(name string) (*model.Tag, error) {
	if name == "" {
		return nil, nil
	}

	var tag model.Tag
	err := database.DB.Where("name = ?", name).First(&tag).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tag = model.Tag{
				Name: name,
			}
			if err := database.DB.Create(&tag).Error; err != nil {
				return nil, err
			}
			return &tag, nil
		}
		return nil, err
	}
	return &tag, nil
}

func (s *TagService) GetTagsByArticleID(articleID uint64) ([]model.TagVO, error) {
	var articleTags []model.ArticleTag
	database.DB.Where("article_id = ?", articleID).Find(&articleTags)

	if len(articleTags) == 0 {
		return []model.TagVO{}, nil
	}

	tagIDs := make([]uint64, len(articleTags))
	for i, at := range articleTags {
		tagIDs[i] = at.TagID
	}

	var tags []model.Tag
	database.DB.Where("id IN ?", tagIDs).Find(&tags)

	var vos []model.TagVO
	for _, tag := range tags {
		vos = append(vos, model.TagVO{
			ID:   tag.ID,
			Name: tag.Name,
		})
	}
	return vos, nil
}
