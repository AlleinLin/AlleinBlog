package service

import (
	"errors"

	"blog-go/constants"
	"blog-go/database"
	"blog-go/model"

	"gorm.io/gorm"
)

type CategoryService struct{}

func (s *CategoryService) GetCategoryList() ([]model.CategoryVO, error) {
	var categories []model.Category
	err := database.DB.Where("status = ? AND del_flag = ?", constants.StatusNormal, 0).Find(&categories).Error
	if err != nil {
		return nil, err
	}

	var vos []model.CategoryVO
	for _, category := range categories {
		var count int64
		database.DB.Model(&model.Article{}).
			Where("category_id = ? AND status = ? AND del_flag = ?", category.ID, constants.ArticleStatusNormal, 0).
			Count(&count)

		vos = append(vos, model.CategoryVO{
			ID:           category.ID,
			Name:         category.Name,
			PID:          category.PID,
			Description:  category.Description,
			Status:       category.Status,
			ArticleCount: count,
		})
	}
	return vos, nil
}

func (s *CategoryService) GetCategoryByID(id uint64) (*model.CategoryVO, error) {
	var category model.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.CategoryNotExist.Msg)
		}
		return nil, err
	}

	return &model.CategoryVO{
		ID:          category.ID,
		Name:        category.Name,
		PID:         category.PID,
		Description: category.Description,
		Status:      category.Status,
	}, nil
}

func (s *CategoryService) GetOrAddByName(name string) (*model.Category, error) {
	if name == "" {
		return nil, nil
	}

	var category model.Category
	err := database.DB.Where("name = ?", name).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			category = model.Category{
				Name:   name,
				PID:    -1,
				Status: constants.StatusNormal,
			}
			if err := database.DB.Create(&category).Error; err != nil {
				return nil, err
			}
			return &category, nil
		}
		return nil, err
	}
	return &category, nil
}
