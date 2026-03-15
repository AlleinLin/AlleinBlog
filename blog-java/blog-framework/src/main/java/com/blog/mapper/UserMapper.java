package com.blog.mapper;

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import com.blog.entity.User;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;
import org.apache.ibatis.annotations.Select;

import java.util.List;

@Mapper
public interface UserMapper extends BaseMapper<User> {

    @Select("SELECT DISTINCT a.permission FROM user_role ur " +
            "LEFT JOIN role r ON ur.role_id = r.id " +
            "LEFT JOIN role_access ra ON ur.role_id = ra.role_id " +
            "LEFT JOIN access a ON a.id = ra.access_id " +
            "WHERE ur.user_id = #{userId} AND r.status = '0' AND (a.status = '0' OR a.status IS NULL)")
    List<String> selectPermissionsByUserId(@Param("userId") Long userId);
}
