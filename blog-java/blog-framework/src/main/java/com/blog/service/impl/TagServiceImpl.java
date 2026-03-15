package com.blog.service.impl;

import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.blog.common.Constants;
import com.blog.entity.ArticleTag;
import com.blog.entity.Tag;
import com.blog.mapper.ArticleTagMapper;
import com.blog.mapper.TagMapper;
import com.blog.service.ArticleTagService;
import com.blog.service.TagService;
import com.blog.util.BeanCopyUtil;
import com.blog.vo.TagVO;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;
import java.util.stream.Collectors;

@Service
public class TagServiceImpl extends ServiceImpl<TagMapper, Tag> implements TagService {

    @Autowired
    private ArticleTagMapper articleTagMapper;

    @Override
    public List<TagVO> getTagList() {
        LambdaQueryWrapper<Tag> wrapper = new LambdaQueryWrapper<>();
        List<Tag> tags = list(wrapper);
        return BeanCopyUtil.copyBeanList(tags, TagVO.class);
    }

    @Override
    public TagVO getTagById(Long id) {
        Tag tag = getById(id);
        if (tag == null) {
            return null;
        }
        return BeanCopyUtil.copyBean(tag, TagVO.class);
    }

    @Override
    public Tag getOrAddTagByName(String name) {
        if (name == null || name.trim().isEmpty()) {
            return null;
        }
        LambdaQueryWrapper<Tag> wrapper = new LambdaQueryWrapper<>();
        wrapper.eq(Tag::getName, name);
        Tag tag = getOne(wrapper);
        if (tag == null) {
            tag = new Tag();
            tag.setName(name);
            save(tag);
        }
        return tag;
    }

    @Override
    public List<TagVO> getTagsByArticleId(Long articleId) {
        LambdaQueryWrapper<ArticleTag> wrapper = new LambdaQueryWrapper<>();
        wrapper.eq(ArticleTag::getArticleId, articleId);
        List<ArticleTag> articleTags = articleTagMapper.selectList(wrapper);
        if (articleTags.isEmpty()) {
            return new ArrayList<>();
        }
        List<Long> tagIds = articleTags.stream()
                .map(ArticleTag::getTagId)
                .collect(Collectors.toList());
        List<Tag> tags = listByIds(tagIds);
        return BeanCopyUtil.copyBeanList(tags, TagVO.class);
    }
}
