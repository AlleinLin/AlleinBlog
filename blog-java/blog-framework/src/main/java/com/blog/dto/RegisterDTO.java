package com.blog.dto;

import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotBlank;
import lombok.Data;

@Data
public class RegisterDTO {

    @NotBlank(message = "用户名不能为空")
    private String userName;

    @NotBlank(message = "昵称不能为空")
    private String nickName;

    @NotBlank(message = "密码不能为空")
    private String password;

    @Email(message = "邮箱格式错误")
    private String email;
}
