package com.blog.service;

import com.baomidou.mybatisplus.extension.service.IService;
import com.blog.dto.ArticleDTO;
import com.blog.dto.ArticleQueryDTO;
import com.blog.entity.Article;
import com.blog.vo.ArchiveVO;
import com.blog.vo.ArticleVO;
import com.blog.vo.HotArticleVO;

import java.util.List;

public interface ArticleService extends IService<Article> {

    List<HotArticleVO> getHotArticleList();

    com.baomidou.mybatisplus.extension.plugins.pagination.Page<ArticleVO> getArticleList(ArticleQueryDTO queryDTO);

    ArticleVO getArticleDetail(Long id);

    Long addArticle(ArticleDTO articleDTO);

    void updateArticle(ArticleDTO articleDTO);

    void deleteArticle(Long id);

    Long getArticleCount();

    void updateViewCount(Long id);

    List<ArchiveVO> getArchiveList(Integer pageNum, Integer pageSize);

    ArticleVO getPreviousNextArticle(Long id);
}
