package com.blog.service;

import com.baomidou.mybatisplus.extension.service.IService;
import com.blog.dto.LoginDTO;
import com.blog.dto.RegisterDTO;
import com.blog.entity.User;
import com.blog.vo.LoginVO;
import com.blog.vo.UserInfoVO;

public interface UserService extends IService<User> {

    LoginVO login(LoginDTO loginDTO);

    void logout();

    UserInfoVO getUserInfo();

    void updateUserInfo(UserInfoVO userInfoVO);

    UserInfoVO getAdminInfo();

    void register(RegisterDTO registerDTO);

    User getByUserName(String userName);
}
