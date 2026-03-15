package com.blog.entity;

import com.baomidou.mybatisplus.annotation.*;
import lombok.Data;

import java.io.Serial;
import java.io.Serializable;

@Data
@TableName("role_access")
public class RoleAccess implements Serializable {

    @Serial
    private static final long serialVersionUID = 1L;

    private Long roleId;

    private Long accessId;
}
