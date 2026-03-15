package com.blog.common;

import lombok.AllArgsConstructor;
import lombok.Getter;

@Getter
@AllArgsConstructor
public enum ErrorCode {

    SUCCESS(200, "操作成功"),
    FAILED(500, "操作失败"),
    VALIDATE_FAILED(400, "参数校验失败"),
    UNAUTHORIZED(401, "未登录或token已过期"),
    FORBIDDEN(403, "没有相关权限"),
    NOT_FOUND(404, "资源不存在"),
    USER_NOT_EXIST(1001, "用户不存在"),
    USER_PASSWORD_ERROR(1002, "用户密码错误"),
    USER_EXIST(1003, "用户名已存在"),
    ARTICLE_NOT_EXIST(2001, "文章不存在"),
    CATEGORY_NOT_EXIST(2002, "分类不存在"),
    TAG_NOT_EXIST(2003, "标签不存在"),
    COMMENT_NOT_EXIST(3001, "评论不存在");

    private final Integer code;
    private final String msg;
}
