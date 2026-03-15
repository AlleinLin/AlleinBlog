package com.blog.entity;

import com.baomidou.mybatisplus.annotation.*;
import lombok.Data;

import java.io.Serial;
import java.io.Serializable;
import java.time.LocalDateTime;

@Data
@TableName("comment")
public class Comment implements Serializable {

    @Serial
    private static final long serialVersionUID = 1L;

    @TableId(type = IdType.AUTO)
    private Long id;

    private Long articleId;

    private Long parentId;

    private String content;

    private Long toUserId;

    @TableField(fill = FieldFill.INSERT)
    private Long createBy;

    @TableField(fill = FieldFill.INSERT)
    private LocalDateTime createTime;

    @TableField(fill = FieldFill.INSERT_UPDATE)
    private Long updateBy;

    @TableField(fill = FieldFill.INSERT_UPDATE)
    private LocalDateTime updateTime;

    @TableLogic
    private Integer delFlag;

    @TableField(exist = false)
    private String userName;

    @TableField(exist = false)
    private String nickName;

    @TableField(exist = false)
    private String avatar;

    @TableField(exist = false)
    private String toUserName;

    @TableField(exist = false)
    private String toNickName;

    @TableField(exist = false)
    private java.util.List<Comment> children;
}
