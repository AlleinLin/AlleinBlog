package com.blog.service;

import com.baomidou.mybatisplus.extension.service.IService;
import com.blog.entity.ArticleTag;

import java.util.List;

public interface ArticleTagService extends IService<ArticleTag> {

    void saveArticleTags(Long articleId, List<Long> tagIds);

    void deleteArticleTags(Long articleId);

    List<Long> getTagIdsByArticleId(Long articleId);
}
