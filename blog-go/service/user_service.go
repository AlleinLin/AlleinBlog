package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"blog-go/config"
	"blog-go/constants"
	"blog-go/database"
	"blog-go/model"
	"blog-go/utils"

	"gorm.io/gorm"
)

type UserService struct{}

func (s *UserService) Login(dto *model.LoginDTO) (*model.LoginVO, error) {
	var user model.User
	if err := database.DB.Where("user_name = ?", dto.UserName).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.UserNotExist.Msg)
		}
		return nil, err
	}

	if !utils.CheckPassword(dto.Password, user.Password) {
		return nil, errors.New(constants.UserPasswordError.Msg)
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	userKey := constants.RedisUserKey + strconv.FormatUint(user.ID, 10)
	userInfo := s.toUserInfoVO(&user)
	utils.SetCacheObjectWithExpire(ctx, userKey, userInfo, config.AppConfig.JWT.Expire)

	return &model.LoginVO{
		Token:    token,
		UserInfo: *userInfo,
	}, nil
}

func (s *UserService) Logout(userID uint64) error {
	ctx := context.Background()
	userKey := constants.RedisUserKey + strconv.FormatUint(userID, 10)
	return utils.DeleteCacheObject(ctx, userKey)
}

func (s *UserService) Register(dto *model.RegisterDTO) error {
	var count int64
	database.DB.Model(&model.User{}).Where("user_name = ?", dto.UserName).Count(&count)
	if count > 0 {
		return errors.New(constants.UserExist.Msg)
	}

	hashPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return err
	}

	user := &model.User{
		UserName:  dto.UserName,
		NickName:  dto.NickName,
		Password:  hashPassword,
		Email:     dto.Email,
		Type:      constants.UserTypeNormal,
		Status:    constants.StatusNormal,
		Signature: "这个人很懒，什么都没写",
	}

	if err := database.DB.Create(user).Error; err != nil {
		return err
	}

	var role model.Role
	if err := database.DB.Where("role_key = ?", "normal_user").First(&role).Error; err == nil {
		userRole := &model.UserRole{
			UserID: user.ID,
			RoleID: role.ID,
		}
		database.DB.Create(userRole)
	}

	return nil
}

func (s *UserService) GetUserInfo(userID uint64) (*model.UserInfoVO, error) {
	var user model.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return s.toUserInfoVO(&user), nil
}

func (s *UserService) UpdateUserInfo(userID uint64, dto *model.UserInfoDTO) error {
	updates := map[string]interface{}{
		"nick_name":   dto.NickName,
		"signature":   dto.Signature,
		"email":       dto.Email,
		"phonenumber": dto.Phonenumber,
		"sex":         dto.Sex,
		"avatar":      dto.Avatar,
		"update_time": time.Now(),
	}
	return database.DB.Model(&model.User{}).Where("id = ?", userID).Updates(updates).Error
}

func (s *UserService) GetAdminInfo() (*model.UserInfoVO, error) {
	var user model.User
	if err := database.DB.Where("type = ?", constants.UserTypeAdmin).First(&user).Error; err != nil {
		return nil, err
	}
	return s.toUserInfoVO(&user), nil
}

func (s *UserService) GetUserByID(userID uint64) (*model.User, error) {
	var user model.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetUsersByIDs(userIDs []uint64) ([]model.User, error) {
	var users []model.User
	if err := database.DB.Where("id IN ?", userIDs).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) toUserInfoVO(user *model.User) *model.UserInfoVO {
	return &model.UserInfoVO{
		ID:          user.ID,
		UserName:    user.UserName,
		NickName:    user.NickName,
		Signature:   user.Signature,
		Email:       user.Email,
		Phonenumber: user.Phonenumber,
		Sex:         user.Sex,
		Avatar:      user.Avatar,
		Type:        user.Type,
	}
}
