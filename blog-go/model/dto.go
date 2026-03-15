package model

type ArticleDTO struct {
	ID        uint64   `json:"id"`
	Title     string   `json:"title" binding:"required"`
	Content   string   `json:"content" binding:"required"`
	Summary   string   `json:"summary"`
	Category  string   `json:"category"`
	Thumbnail string   `json:"thumbnail"`
	IsTop     string   `json:"isTop"`
	IsDraft   bool     `json:"isDraft"`
	IsComment string   `json:"isComment"`
	Tags      []string `json:"tags"`
}

type ArticleQueryDTO struct {
	PageNum    int    `form:"pageNum" json:"pageNum"`
	PageSize   int    `form:"pageSize" json:"pageSize"`
	CategoryID uint64 `form:"categoryId" json:"categoryId"`
	TagID      uint64 `form:"tagId" json:"tagId"`
	YearMonth  string `form:"yearMonth" json:"yearMonth"`
	Title      string `form:"title" json:"title"`
}

type LoginDTO struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterDTO struct {
	UserName string `json:"userName" binding:"required"`
	NickName string `json:"nickName" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"omitempty,email"`
}

type CommentDTO struct {
	ID        uint64 `json:"id"`
	ArticleID uint64 `json:"articleId" binding:"required"`
	ParentID  uint64 `json:"parentId"`
	Content   string `json:"content" binding:"required"`
	ToUserID  uint64 `json:"toUserId"`
}

type UserInfoDTO struct {
	NickName    string `json:"nickName"`
	Signature   string `json:"signature"`
	Email       string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	Sex         string `json:"sex"`
	Avatar      string `json:"avatar"`
}
