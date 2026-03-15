package model

import (
	"time"
)

type User struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserName    string    `gorm:"uniqueIndex;size:64;not null" json:"userName"`
	NickName    string    `gorm:"size:64;not null" json:"nickName"`
	Signature   string    `gorm:"size:128" json:"signature"`
	Password    string    `gorm:"size:128;not null" json:"-"`
	Type        string    `gorm:"size:1;default:0" json:"type"`
	Status      string    `gorm:"size:1;default:0" json:"status"`
	Email       string    `gorm:"size:64" json:"email"`
	Phonenumber string    `gorm:"size:32" json:"phonenumber"`
	Sex         string    `gorm:"size:1;default:0" json:"sex"`
	Avatar      string    `gorm:"size:256" json:"avatar"`
	CreateBy    uint64    `gorm:"default:null" json:"createBy"`
	CreateTime  time.Time `gorm:"autoCreateTime" json:"createTime"`
	UpdateBy    uint64    `gorm:"default:null" json:"updateBy"`
	UpdateTime  time.Time `gorm:"autoUpdateTime" json:"updateTime"`
	DelFlag     int       `gorm:"default:0" json:"-"`
}

func (User) TableName() string {
	return "user"
}

type Role struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string    `gorm:"size:128" json:"name"`
	RoleKey    string    `gorm:"size:100" json:"roleKey"`
	Status     string    `gorm:"size:1;default:0" json:"status"`
	DelFlag    int       `gorm:"default:0" json:"-"`
	CreateBy   uint64    `gorm:"default:null" json:"createBy"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"createTime"`
	UpdateBy   uint64    `gorm:"default:null" json:"updateBy"`
	UpdateTime time.Time `gorm:"autoUpdateTime" json:"updateTime"`
	Remark     string    `gorm:"size:500" json:"remark"`
}

func (Role) TableName() string {
	return "role"
}

type Access struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	AccessName string    `gorm:"size:64;not null" json:"accessName"`
	Permission string    `gorm:"size:100;not null" json:"permission"`
	Status     string    `gorm:"size:1;default:0" json:"status"`
	CreateBy   uint64    `gorm:"default:null" json:"createBy"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"createTime"`
	UpdateBy   uint64    `gorm:"default:null" json:"updateBy"`
	UpdateTime time.Time `gorm:"autoUpdateTime" json:"updateTime"`
	DelFlag    int       `gorm:"default:0" json:"-"`
	Remark     string    `gorm:"size:500" json:"remark"`
}

func (Access) TableName() string {
	return "access"
}

type UserRole struct {
	UserID uint64 `gorm:"primaryKey" json:"userId"`
	RoleID uint64 `gorm:"primaryKey" json:"roleId"`
}

func (UserRole) TableName() string {
	return "user_role"
}

type RoleAccess struct {
	RoleID   uint64 `gorm:"primaryKey" json:"roleId"`
	AccessID uint64 `gorm:"primaryKey" json:"accessId"`
}

func (RoleAccess) TableName() string {
	return "role_access"
}

type Category struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"size:128" json:"name"`
	PID         int64     `gorm:"column:pid;default:-1" json:"pid"`
	Description string    `gorm:"size:512" json:"description"`
	Status      string    `gorm:"size:1;default:0" json:"status"`
	CreateBy    uint64    `gorm:"default:null" json:"createBy"`
	CreateTime  time.Time `gorm:"autoCreateTime" json:"createTime"`
	UpdateBy    uint64    `gorm:"default:null" json:"updateBy"`
	UpdateTime  time.Time `gorm:"autoUpdateTime" json:"updateTime"`
	DelFlag     int       `gorm:"default:0" json:"-"`
}

func (Category) TableName() string {
	return "category"
}

type Tag struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string    `gorm:"size:128" json:"name"`
	CreateBy   uint64    `gorm:"default:null" json:"createBy"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"createTime"`
	UpdateBy   uint64    `gorm:"default:null" json:"updateBy"`
	UpdateTime time.Time `gorm:"autoUpdateTime" json:"updateTime"`
	DelFlag    int       `gorm:"default:0" json:"-"`
}

func (Tag) TableName() string {
	return "tag"
}

type Article struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title      string    `gorm:"size:256" json:"title"`
	Content    string    `gorm:"type:longtext" json:"content"`
	Summary    string    `gorm:"size:1024" json:"summary"`
	CategoryID uint64    `gorm:"column:category_id" json:"categoryId"`
	Thumbnail  string    `gorm:"size:256" json:"thumbnail"`
	IsTop      string    `gorm:"column:is_top;size:1;default:0" json:"isTop"`
	Status     string    `gorm:"size:1;default:0" json:"status"`
	ViewCount  uint64    `gorm:"column:view_count;default:0" json:"viewCount"`
	IsComment  string    `gorm:"column:is_comment;size:1;default:1" json:"isComment"`
	CreateBy   uint64    `gorm:"default:null" json:"createBy"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"createTime"`
	UpdateBy   uint64    `gorm:"default:null" json:"updateBy"`
	UpdateTime time.Time `gorm:"autoUpdateTime" json:"updateTime"`
	DelFlag    int       `gorm:"default:0" json:"-"`
}

func (Article) TableName() string {
	return "article"
}

type ArticleTag struct {
	ArticleID uint64 `gorm:"primaryKey" json:"articleId"`
	TagID     uint64 `gorm:"primaryKey" json:"tagId"`
}

func (ArticleTag) TableName() string {
	return "article_tag"
}

type Comment struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	ArticleID  uint64    `gorm:"column:article_id" json:"articleId"`
	ParentID   uint64    `gorm:"column:parent_id;default:0" json:"parentId"`
	Content    string    `gorm:"size:1024" json:"content"`
	ToUserID   uint64    `gorm:"column:to_user_id" json:"toUserId"`
	CreateBy   uint64    `gorm:"default:null" json:"createBy"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"createTime"`
	UpdateBy   uint64    `gorm:"default:null" json:"updateBy"`
	UpdateTime time.Time `gorm:"autoUpdateTime" json:"updateTime"`
	DelFlag    int       `gorm:"default:0" json:"-"`
}

func (Comment) TableName() string {
	return "comment"
}
