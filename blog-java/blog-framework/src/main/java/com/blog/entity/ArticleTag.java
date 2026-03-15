package com.blog.entity;

import com.baomidou.mybatisplus.annotation.*;
import lombok.Data;

import java.io.Serial;
import java.io.Serializable;

@Data
@TableName("article_tag")
public class ArticleTag implements Serializable {

    @Serial
    private static final long serialVersionUID = 1L;

    private Long articleId;

    private Long tagId;
}
