package com.blog.dto;

import lombok.Data;

@Data
public class ArticleQueryDTO {

    private Integer pageNum = 1;

    private Integer pageSize = 10;

    private Long categoryId;

    private Long tagId;

    private String yearMonth;

    private String title;
}
