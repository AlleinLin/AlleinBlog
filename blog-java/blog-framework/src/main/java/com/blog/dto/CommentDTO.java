package com.blog.dto;

import jakarta.validation.constraints.NotBlank;
import lombok.Data;

@Data
public class CommentDTO {

    private Long id;

    private Long articleId;

    private Long parentId;

    private String content;

    private Long toUserId;
}
