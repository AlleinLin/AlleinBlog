package com.blog.service;

import com.baomidou.mybatisplus.extension.service.IService;
import com.blog.entity.Tag;
import com.blog.vo.TagVO;

import java.util.List;

public interface TagService extends IService<Tag> {

    List<TagVO> getTagList();

    TagVO getTagById(Long id);

    Tag getOrAddTagByName(String name);

    List<TagVO> getTagsByArticleId(Long articleId);
}
