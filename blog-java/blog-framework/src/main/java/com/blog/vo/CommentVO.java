package com.blog.vo;

import lombok.Data;

import java.time.LocalDateTime;
import java.util.List;

@Data
public class CommentVO {

    private Long id;

    private Long articleId;

    private Long parentId;

    private String content;

    private Long createBy;

    private String userName;

    private String nickName;

    private String avatar;

    private Long toUserId;

    private String toUserName;

    private String toNickName;

    private LocalDateTime createTime;

    private List<CommentVO> children;
}
