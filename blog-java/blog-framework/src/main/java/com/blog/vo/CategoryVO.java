package com.blog.vo;

import lombok.Data;

import java.time.LocalDateTime;

@Data
public class CategoryVO {

    private Long id;

    private String name;

    private Long pid;

    private String description;

    private String status;

    private Long articleCount;
}
