package constants

const (
	ArticleStatusNormal = "0"
	ArticleStatusDraft  = "1"

	StatusNormal  = "0"
	StatusDisable = "1"

	UserTypeAdmin  = "1"
	UserTypeNormal = "0"

	RedisUserKey        = "blog:user:"
	RedisArticleViewKey = "blog:article:view"

	DefaultPassword = "123456"
)

type ErrorCode struct {
	Code int
	Msg  string
}

var (
	Success           = ErrorCode{200, "操作成功"}
	Failed            = ErrorCode{500, "操作失败"}
	ValidateFailed    = ErrorCode{400, "参数校验失败"}
	Unauthorized      = ErrorCode{401, "未登录或token已过期"}
	Forbidden         = ErrorCode{403, "没有相关权限"}
	NotFound          = ErrorCode{404, "资源不存在"}
	UserNotExist      = ErrorCode{1001, "用户不存在"}
	UserPasswordError = ErrorCode{1002, "用户密码错误"}
	UserExist         = ErrorCode{1003, "用户名已存在"}
	ArticleNotExist   = ErrorCode{2001, "文章不存在"}
	CategoryNotExist  = ErrorCode{2002, "分类不存在"}
	TagNotExist       = ErrorCode{2003, "标签不存在"}
	CommentNotExist   = ErrorCode{3001, "评论不存在"}
)
