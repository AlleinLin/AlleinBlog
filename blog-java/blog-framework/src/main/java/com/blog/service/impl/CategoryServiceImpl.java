package com.blog.service.impl;

import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.blog.common.Constants;
import com.blog.entity.Category;
import com.blog.mapper.ArticleMapper;
import com.blog.mapper.CategoryMapper;
import com.blog.service.CategoryService;
import com.blog.util.BeanCopyUtil;
import com.blog.vo.CategoryVO;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.stream.Collectors;

@Service
public class CategoryServiceImpl extends ServiceImpl<CategoryMapper, Category> implements CategoryService {

    @Autowired
    private ArticleMapper articleMapper;

    @Override
    public List<CategoryVO> getCategoryList() {
        LambdaQueryWrapper<Category> wrapper = new LambdaQueryWrapper<>();
        wrapper.eq(Category::getStatus, Constants.STATUS_NORMAL);
        List<Category> categories = list(wrapper);
        return categories.stream().map(category -> {
            CategoryVO vo = BeanCopyUtil.copyBean(category, CategoryVO.class);
            LambdaQueryWrapper<com.blog.entity.Article> articleWrapper = new LambdaQueryWrapper<>();
            articleWrapper.eq(com.blog.entity.Article::getCategoryId, category.getId())
                    .eq(com.blog.entity.Article::getStatus, Constants.ARTICLE_STATUS_NORMAL);
            vo.setArticleCount(count(articleWrapper));
            return vo;
        }).collect(Collectors.toList());
    }

    @Override
    public CategoryVO getCategoryById(Long id) {
        Category category = getById(id);
        if (category == null) {
            return null;
        }
        return BeanCopyUtil.copyBean(category, CategoryVO.class);
    }

    @Override
    public Category getOrAddCategoryByName(String name) {
        if (name == null || name.trim().isEmpty()) {
            return null;
        }
        LambdaQueryWrapper<Category> wrapper = new LambdaQueryWrapper<>();
        wrapper.eq(Category::getName, name);
        Category category = getOne(wrapper);
        if (category == null) {
            category = new Category();
            category.setName(name);
            category.setPid(-1L);
            category.setStatus(Constants.STATUS_NORMAL);
            save(category);
        }
        return category;
    }
}
