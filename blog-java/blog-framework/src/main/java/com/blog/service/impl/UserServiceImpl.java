package com.blog.service.impl;

import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.blog.common.Constants;
import com.blog.common.ErrorCode;
import com.blog.dto.LoginDTO;
import com.blog.dto.RegisterDTO;
import com.blog.entity.Role;
import com.blog.entity.User;
import com.blog.entity.UserRole;
import com.blog.exception.BusinessException;
import com.blog.mapper.UserMapper;
import com.blog.security.LoginUser;
import com.blog.service.RoleService;
import com.blog.service.UserRoleService;
import com.blog.service.UserService;
import com.blog.util.BeanCopyUtil;
import com.blog.util.JwtUtil;
import com.blog.util.RedisCache;
import com.blog.vo.LoginVO;
import com.blog.vo.UserInfoVO;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;
import java.util.concurrent.TimeUnit;

@Service
public class UserServiceImpl extends ServiceImpl<UserMapper, User> implements UserService {

    @Autowired
    private AuthenticationManager authenticationManager;

    @Autowired
    private RedisCache redisCache;

    @Autowired
    private PasswordEncoder passwordEncoder;

    @Autowired
    private RoleService roleService;

    @Autowired
    private UserRoleService userRoleService;

    @Override
    public LoginVO login(LoginDTO loginDTO) {
        UsernamePasswordAuthenticationToken authenticationToken =
                new UsernamePasswordAuthenticationToken(loginDTO.getUserName(), loginDTO.getPassword());
        Authentication authentication = authenticationManager.authenticate(authenticationToken);
        LoginUser loginUser = (LoginUser) authentication.getPrincipal();
        String userId = loginUser.getUser().getId().toString();
        String token = JwtUtil.createToken(Long.parseLong(userId));
        redisCache.setCacheObject(Constants.REDIS_USER_KEY + userId, loginUser, 24, TimeUnit.HOURS);
        UserInfoVO userInfo = BeanCopyUtil.copyBean(loginUser.getUser(), UserInfoVO.class);
        return new LoginVO(token, userInfo);
    }

    @Override
    public void logout() {
        Authentication authentication = SecurityContextHolder.getContext().getAuthentication();
        if (authentication != null && authentication.getPrincipal() instanceof LoginUser loginUser) {
            redisCache.deleteObject(Constants.REDIS_USER_KEY + loginUser.getUser().getId());
        }
    }

    @Override
    public UserInfoVO getUserInfo() {
        Authentication authentication = SecurityContextHolder.getContext().getAuthentication();
        if (authentication != null && authentication.getPrincipal() instanceof LoginUser loginUser) {
            return BeanCopyUtil.copyBean(loginUser.getUser(), UserInfoVO.class);
        }
        return null;
    }

    @Override
    public void updateUserInfo(UserInfoVO userInfoVO) {
        Authentication authentication = SecurityContextHolder.getContext().getAuthentication();
        if (authentication != null && authentication.getPrincipal() instanceof LoginUser loginUser) {
            User user = BeanCopyUtil.copyBean(userInfoVO, User.class);
            user.setId(loginUser.getUser().getId());
            updateById(user);
            loginUser.setUser(getById(user.getId()));
            redisCache.setCacheObject(Constants.REDIS_USER_KEY + user.getId(), loginUser);
        }
    }

    @Override
    public UserInfoVO getAdminInfo() {
        LambdaQueryWrapper<User> wrapper = new LambdaQueryWrapper<>();
        wrapper.eq(User::getType, Constants.USER_TYPE_ADMIN);
        User admin = getOne(wrapper);
        if (admin == null) {
            return null;
        }
        return BeanCopyUtil.copyBean(admin, UserInfoVO.class);
    }

    @Override
    @Transactional
    public void register(RegisterDTO registerDTO) {
        User existUser = getByUserName(registerDTO.getUserName());
        if (existUser != null) {
            throw new BusinessException(ErrorCode.USER_EXIST);
        }
        User user = BeanCopyUtil.copyBean(registerDTO, User.class);
        user.setPassword(passwordEncoder.encode(registerDTO.getPassword()));
        user.setType(Constants.USER_TYPE_NORMAL);
        user.setStatus(Constants.STATUS_NORMAL);
        save(user);
        LambdaQueryWrapper<Role> roleWrapper = new LambdaQueryWrapper<>();
        roleWrapper.eq(Role::getRoleKey, "normal_user");
        Role role = roleService.getOne(roleWrapper);
        if (role != null) {
            UserRole userRole = new UserRole();
            userRole.setUserId(user.getId());
            userRole.setRoleId(role.getId());
            userRoleService.save(userRole);
        }
    }

    @Override
    public User getByUserName(String userName) {
        LambdaQueryWrapper<User> wrapper = new LambdaQueryWrapper<>();
        wrapper.eq(User::getUserName, userName);
        return getOne(wrapper);
    }
}
