package com.blog.common;

public class Constants {

    private Constants() {}

    public static final String ARTICLE_STATUS_NORMAL = "0";
    public static final String ARTICLE_STATUS_DRAFT = "1";

    public static final String STATUS_NORMAL = "0";
    public static final String STATUS_DISABLE = "1";

    public static final String USER_TYPE_ADMIN = "1";
    public static final String USER_TYPE_NORMAL = "0";

    public static final String REDIS_USER_KEY = "blog:user:";
    public static final String REDIS_ARTICLE_VIEW_KEY = "blog:article:view";

    public static final String DEFAULT_PASSWORD = "123456";
}
