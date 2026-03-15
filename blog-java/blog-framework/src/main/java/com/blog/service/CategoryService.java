package com.blog.service;

import com.baomidou.mybatisplus.extension.service.IService;
import com.blog.entity.Category;
import com.blog.vo.CategoryVO;

import java.util.List;

public interface CategoryService extends IService<Category> {

    List<CategoryVO> getCategoryList();

    CategoryVO getCategoryById(Long id);

    Category getOrAddCategoryByName(String name);
}
