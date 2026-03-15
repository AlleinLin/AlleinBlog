package com.blog.vo;

import lombok.Data;

import java.time.LocalDateTime;

@Data
public class ArticleVO {

    private Long id;

    private String title;

    private String content;

    private String summary;

    private Long categoryId;

    private String categoryName;

    private String thumbnail;

    private String isTop;

    private String status;

    private Long viewCount;

    private String isComment;

    private LocalDateTime createTime;

    private LocalDateTime updateTime;

    private String userName;

    private String nickName;

    private java.util.List<TagVO> tags;
}
