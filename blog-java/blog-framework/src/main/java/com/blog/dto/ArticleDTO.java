package com.blog.dto;

import jakarta.validation.constraints.NotBlank;
import lombok.Data;

import java.util.List;

@Data
public class ArticleDTO {

    private Long id;

    @NotBlank(message = "文章标题不能为空")
    private String title;

    @NotBlank(message = "文章内容不能为空")
    private String content;

    private String summary;

    private String category;

    private String thumbnail;

    private String isTop;

    private Boolean isDraft;

    private String isComment;

    private List<String> tags;
}
