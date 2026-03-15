package model

import "time"

type ArticleVO struct {
	ID           uint64    `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Summary      string    `json:"summary"`
	CategoryID   uint64    `json:"categoryId"`
	CategoryName string    `json:"categoryName"`
	Thumbnail    string    `json:"thumbnail"`
	IsTop        string    `json:"isTop"`
	Status       string    `json:"status"`
	ViewCount    uint64    `json:"viewCount"`
	IsComment    string    `json:"isComment"`
	CreateTime   time.Time `json:"createTime"`
	UpdateTime   time.Time `json:"updateTime"`
	UserName     string    `json:"userName"`
	NickName     string    `json:"nickName"`
	Tags         []TagVO   `json:"tags"`
}

type TagVO struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type CategoryVO struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	PID          int64  `json:"pid"`
	Description  string `json:"description"`
	Status       string `json:"status"`
	ArticleCount int64  `json:"articleCount"`
}

type CommentVO struct {
	ID         uint64      `json:"id"`
	ArticleID  uint64      `json:"articleId"`
	ParentID   uint64      `json:"parentId"`
	Content    string      `json:"content"`
	CreateBy   uint64      `json:"createBy"`
	UserName   string      `json:"userName"`
	NickName   string      `json:"nickName"`
	Avatar     string      `json:"avatar"`
	ToUserID   uint64      `json:"toUserId"`
	ToUserName string      `json:"toUserName"`
	ToNickName string      `json:"toNickName"`
	CreateTime time.Time   `json:"createTime"`
	Children   []CommentVO `json:"children"`
}

type UserInfoVO struct {
	ID          uint64 `json:"id"`
	UserName    string `json:"userName"`
	NickName    string `json:"nickName"`
	Signature   string `json:"signature"`
	Email       string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	Sex         string `json:"sex"`
	Avatar      string `json:"avatar"`
	Type        string `json:"type"`
}

type LoginVO struct {
	Token    string     `json:"token"`
	UserInfo UserInfoVO `json:"userInfo"`
}

type HotArticleVO struct {
	ID        uint64 `json:"id"`
	Title     string `json:"title"`
	ViewCount uint64 `json:"viewCount"`
}

type ArchiveVO struct {
	Date  string `json:"date"`
	Count int64  `json:"count"`
}
