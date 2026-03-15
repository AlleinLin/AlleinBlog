package service

import (
	"errors"

	"blog-go/constants"
	"blog-go/database"
	"blog-go/model"

	"gorm.io/gorm"
)

type CommentService struct{}

func (s *CommentService) GetCommentList(articleID uint64, pageNum, pageSize int) (*model.PageResult, error) {
	var comments []model.Comment
	var total int64

	db := database.DB.Model(&model.Comment{}).
		Where("article_id = ? AND parent_id = 0 AND del_flag = ?", articleID, 0)

	db.Count(&total)

	offset := (pageNum - 1) * pageSize
	err := db.Order("create_time DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&comments).Error
	if err != nil {
		return nil, err
	}

	vos := s.toVOList(comments)

	for i, vo := range vos {
		var children []model.Comment
		database.DB.Where("parent_id = ?", vo.ID).Order("create_time ASC").Find(&children)
		vos[i].Children = s.toVOList(children)
	}

	return model.NewPageResult(vos, total, pageNum, pageSize), nil
}

func (s *CommentService) AddComment(dto *model.CommentDTO, userID uint64) error {
	comment := &model.Comment{
		ArticleID: dto.ArticleID,
		ParentID:  dto.ParentID,
		Content:   dto.Content,
		ToUserID:  dto.ToUserID,
		CreateBy:  userID,
	}
	return database.DB.Create(comment).Error
}

func (s *CommentService) UpdateComment(dto *model.CommentDTO, userID uint64) error {
	var comment model.Comment
	if err := database.DB.First(&comment, dto.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(constants.CommentNotExist.Msg)
		}
		return err
	}

	if comment.CreateBy != userID {
		return errors.New("无权修改他人评论")
	}

	return database.DB.Model(&comment).Update("content", dto.Content).Error
}

func (s *CommentService) DeleteComment(id, userID uint64) error {
	var comment model.Comment
	if err := database.DB.First(&comment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(constants.CommentNotExist.Msg)
		}
		return err
	}

	if comment.CreateBy != userID {
		return errors.New("无权删除他人评论")
	}

	return database.DB.Model(&comment).Update("del_flag", 1).Error
}

func (s *CommentService) toVO(comment *model.Comment) model.CommentVO {
	vo := model.CommentVO{
		ID:         comment.ID,
		ArticleID:  comment.ArticleID,
		ParentID:   comment.ParentID,
		Content:    comment.Content,
		CreateBy:   comment.CreateBy,
		ToUserID:   comment.ToUserID,
		CreateTime: comment.CreateTime,
	}

	userService := &UserService{}
	if user, err := userService.GetUserByID(comment.CreateBy); err == nil {
		vo.UserName = user.UserName
		vo.NickName = user.NickName
		vo.Avatar = user.Avatar
	}

	if comment.ToUserID > 0 {
		if user, err := userService.GetUserByID(comment.ToUserID); err == nil {
			vo.ToUserName = user.UserName
			vo.ToNickName = user.NickName
		}
	}

	return vo
}

func (s *CommentService) toVOList(comments []model.Comment) []model.CommentVO {
	if len(comments) == 0 {
		return []model.CommentVO{}
	}

	userIDs := make([]uint64, 0)
	for _, c := range comments {
		userIDs = append(userIDs, c.CreateBy)
		if c.ToUserID > 0 {
			userIDs = append(userIDs, c.ToUserID)
		}
	}

	userService := &UserService{}
	users, _ := userService.GetUsersByIDs(userIDs)
	userMap := make(map[uint64]*model.User)
	for i := range users {
		userMap[users[i].ID] = &users[i]
	}

	var vos []model.CommentVO
	for _, comment := range comments {
		vo := model.CommentVO{
			ID:         comment.ID,
			ArticleID:  comment.ArticleID,
			ParentID:   comment.ParentID,
			Content:    comment.Content,
			CreateBy:   comment.CreateBy,
			ToUserID:   comment.ToUserID,
			CreateTime: comment.CreateTime,
		}

		if user, ok := userMap[comment.CreateBy]; ok {
			vo.UserName = user.UserName
			vo.NickName = user.NickName
			vo.Avatar = user.Avatar
		}

		if comment.ToUserID > 0 {
			if user, ok := userMap[comment.ToUserID]; ok {
				vo.ToUserName = user.UserName
				vo.ToNickName = user.NickName
			}
		}

		vos = append(vos, vo)
	}

	return vos
}
