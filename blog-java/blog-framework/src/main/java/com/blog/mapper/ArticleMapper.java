package com.blog.mapper;

import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import com.blog.entity.Article;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;
import org.apache.ibatis.annotations.Select;

import java.util.List;

@Mapper
public interface ArticleMapper extends BaseMapper<Article> {

    @Select("SELECT DATE_FORMAT(create_time, '%Y/%c') as date, COUNT(*) as count " +
            "FROM article WHERE del_flag = 0 AND status = '0' " +
            "GROUP BY date ORDER BY date DESC LIMIT #{offset}, #{limit}")
    List<com.blog.vo.ArchiveVO> selectArchiveList(@Param("offset") int offset, @Param("limit") int limit);

    @Select("SELECT COUNT(*) FROM article WHERE del_flag = 0 AND status = '0'")
    Long selectArticleCount();
}
