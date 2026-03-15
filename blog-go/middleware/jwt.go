package middleware

import (
	"context"
	"strconv"
	"strings"

	"blog-go/constants"
	"blog-go/database"
	"blog-go/model"
	"blog-go/response"
	"blog-go/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(c, constants.Unauthorized.Msg)
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Unauthorized(c, constants.Unauthorized.Msg)
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(parts[1])
		if err != nil {
			response.Unauthorized(c, constants.Unauthorized.Msg)
			c.Abort()
			return
		}

		ctx := context.Background()
		userKey := constants.RedisUserKey + strconv.FormatUint(claims.UserID, 10)
		userData, err := database.RedisClient.Get(ctx, userKey).Result()
		if err != nil || userData == "" {
			response.Unauthorized(c, constants.Unauthorized.Msg)
			c.Abort()
			return
		}

		c.Set("userId", claims.UserID)
		c.Next()
	}
}

func GetUserID(c *gin.Context) uint64 {
	if userID, exists := c.Get("userId"); exists {
		return userID.(uint64)
	}
	return 0
}

func GetUserInfo(c *gin.Context) *model.UserInfoVO {
	if userInfo, exists := c.Get("userInfo"); exists {
		return userInfo.(*model.UserInfoVO)
	}
	return nil
}
