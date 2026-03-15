package com.blog.controller;

import com.blog.common.Result;
import com.blog.dto.LoginDTO;
import com.blog.dto.RegisterDTO;
import com.blog.service.UserService;
import com.blog.vo.LoginVO;
import com.blog.vo.UserInfoVO;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import jakarta.validation.Valid;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@Tag(name = "用户管理")
@RestController
public class UserController {

    @Autowired
    private UserService userService;

    @Operation(summary = "用户登录")
    @PostMapping("/login")
    public Result<LoginVO> login(@Valid @RequestBody LoginDTO loginDTO) {
        return Result.ok(userService.login(loginDTO));
    }

    @Operation(summary = "用户登出")
    @PostMapping("/logout")
    public Result<Void> logout() {
        userService.logout();
        return Result.ok();
    }

    @Operation(summary = "用户注册")
    @PostMapping("/register")
    public Result<Void> register(@Valid @RequestBody RegisterDTO registerDTO) {
        userService.register(registerDTO);
        return Result.ok();
    }

    @Operation(summary = "获取当前用户信息")
    @GetMapping("/user/info")
    public Result<UserInfoVO> getUserInfo() {
        return Result.ok(userService.getUserInfo());
    }

    @Operation(summary = "更新用户信息")
    @PutMapping("/user/info")
    public Result<Void> updateUserInfo(@RequestBody UserInfoVO userInfoVO) {
        userService.updateUserInfo(userInfoVO);
        return Result.ok();
    }

    @Operation(summary = "获取管理员信息")
    @GetMapping("/user/admin")
    public Result<UserInfoVO> getAdminInfo() {
        return Result.ok(userService.getAdminInfo());
    }
}
