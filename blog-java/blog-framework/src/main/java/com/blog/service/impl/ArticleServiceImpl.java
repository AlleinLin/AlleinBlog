package com.blog.service.impl;

import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import com.baomidou.mybatisplus.extension.plugins.pagination.Page;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.blog.common.Constants;
import com.blog.dto.ArticleDTO;
import com.blog.dto.ArticleQueryDTO;
import com.blog.entity.Article;
import com.blog.entity.ArticleTag;
import com.blog.entity.Category;
import com.blog.entity.Tag;
import com.blog.exception.BusinessException;
import com.blog.mapper.ArticleMapper;
import com.blog.service.ArticleService;
import com.blog.service.ArticleTagService;
import com.blog.service.CategoryService;
import com.blog.service.TagService;
import com.blog.util.BeanCopyUtil;
import com.blog.util.RedisCache;
import com.blog.vo.ArchiveVO;
import com.blog.vo.ArticleVO;
import com.blog.vo.HotArticleVO;
import com.blog.vo.TagVO;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.util.StringUtils;

import java.util.List;
import java.util.stream.Collectors;

@Service
public class ArticleServiceImpl extends ServiceImpl<ArticleMapper, Article> implements ArticleService {

    @Autowired
    private CategoryService categoryService;

    @Autowired
    private TagService tagService;

    @Autowired
    private ArticleTagService articleTagService;

    @Autowired
    private RedisCache redisCache;

    @Override
    public List<HotArticleVO> getHotArticleList() {
        LambdaQueryWrapper<Article> wrapper = new LambdaQueryWrapper<>();
        wrapper.eq(Article::getStatus, Constants.ARTICLE_STATUS_NORMAL)
                .orderByDesc(Article::getViewCount)
                .last("LIMIT 5");
        List<Article> articles = list(wrapper);
        return BeanCopyUtil.copyBeanList(articles, HotArticleVO.class);
    }

    @Override
    public Page<ArticleVO> getArticleList(ArticleQueryDTO queryDTO) {
        Page<Article> page = new Page<>(queryDTO.getPageNum(), queryDTO.getPageSize());
        LambdaQueryWrapper<Article> wrapper = new LambdaQueryWrapper<>();
        wrapper.eq(Article::getStatus, Constants.ARTICLE_STATUS_NORMAL);
        if (queryDTO.getCategoryId() != null) {
            wrapper.eq(Article::getCategoryId, queryDTO.getCategoryId());
        }
        if (queryDTO.getTagId() != null) {
            List<Long> articleIds = articleTagService.getTagIdsByArticleId(queryDTO.getTagId());
            if (articleIds.isEmpty()) {
                return new Page<>();
            }
            wrapper.in(Article::getId, articleIds);
        }
        if (StringUtils.hasText(queryDTO.getYearMonth())) {
            wrapper.apply("DATE_FORMAT(create_time, '%Y/%c') = {0}", queryDTO.getYearMonth());
        }
        if (StringUtils.hasText(queryDTO.getTitle())) {
            wrapper.like(Article::getTitle, queryDTO.getTitle());
        }
        wrapper.orderByDesc(Article::getIsTop).orderByDesc(Article::getCreateTime);
        page(page, wrapper);
        Page<ArticleVO> voPage = new Page<>(page.getCurrent(), page.getSize(), page.getTotal());
        voPage.setRecords(convertToVOList(page.getRecords()));
        return voPage;
    }

    @Override
    public ArticleVO getArticleDetail(Long id) {
        Article article = getById(id);
        if (article == null) {
            throw new BusinessException("文章不存在");
        }
        ArticleVO vo = BeanCopyUtil.copyBean(article, ArticleVO.class);
        Category category = categoryService.getById(article.getCategoryId());
        if (category != null) {
            vo.setCategoryName(category.getName());
        }
        List<TagVO> tags = tagService.getTagsByArticleId(id);
        vo.setTags(tags);
        return vo;
    }

    @Override
    @Transactional
    public Long addArticle(ArticleDTO articleDTO) {
        Article article = BeanCopyUtil.copyBean(articleDTO, Article.class);
        Category category = categoryService.getOrAddCategoryByName(articleDTO.getCategory());
        article.setCategoryId(category.getId());
        article.setStatus(Boolean.TRUE.equals(articleDTO.getIsDraft()) ?
                Constants.ARTICLE_STATUS_DRAFT : Constants.ARTICLE_STATUS_NORMAL);
        save(article);
        if (articleDTO.getTags() != null && !articleDTO.getTags().isEmpty()) {
            List<Long> tagIds = articleDTO.getTags().stream()
                    .map(name -> tagService.getOrAddTagByName(name).getId())
                    .collect(Collectors.toList());
            articleTagService.saveArticleTags(article.getId(), tagIds);
        }
        return article.getId();
    }

    @Override
    @Transactional
    public void updateArticle(ArticleDTO articleDTO) {
        Article article = getById(articleDTO.getId());
        if (article == null) {
            throw new BusinessException("文章不存在");
        }
        Article updateArticle = BeanCopyUtil.copyBean(articleDTO, Article.class);
        Category category = categoryService.getOrAddCategoryByName(articleDTO.getCategory());
        updateArticle.setCategoryId(category.getId());
        updateArticle.setStatus(Boolean.TRUE.equals(articleDTO.getIsDraft()) ?
                Constants.ARTICLE_STATUS_DRAFT : Constants.ARTICLE_STATUS_NORMAL);
        updateById(updateArticle);
        articleTagService.deleteArticleTags(articleDTO.getId());
        if (articleDTO.getTags() != null && !articleDTO.getTags().isEmpty()) {
            List<Long> tagIds = articleDTO.getTags().stream()
                    .map(name -> tagService.getOrAddTagByName(name).getId())
                    .collect(Collectors.toList());
            articleTagService.saveArticleTags(articleDTO.getId(), tagIds);
        }
    }

    @Override
    public void deleteArticle(Long id) {
        removeById(id);
        articleTagService.deleteArticleTags(id);
    }

    @Override
    public Long getArticleCount() {
        return baseMapper.selectArticleCount();
    }

    @Override
    public void updateViewCount(Long id) {
        redisCache.increaseCacheMapValue(Constants.REDIS_ARTICLE_VIEW_KEY, id.toString(), 1);
    }

    @Override
    public List<ArchiveVO> getArchiveList(Integer pageNum, Integer pageSize) {
        int offset = (pageNum - 1) * pageSize;
        return baseMapper.selectArchiveList(offset, pageSize);
    }

    @Override
    public ArticleVO getPreviousNextArticle(Long id) {
        return null;
    }

    private List<ArticleVO> convertToVOList(List<Article> articles) {
        return articles.stream().map(article -> {
            ArticleVO vo = BeanCopyUtil.copyBean(article, ArticleVO.class);
            Category category = categoryService.getById(article.getCategoryId());
            if (category != null) {
                vo.setCategoryName(category.getName());
            }
            return vo;
        }).collect(Collectors.toList());
    }
}
