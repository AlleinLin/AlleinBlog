package com.blog.controller;

import com.blog.common.Result;
import com.blog.service.CategoryService;
import com.blog.vo.CategoryVO;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@Tag(name = "分类管理")
@RestController
@RequestMapping("/category")
public class CategoryController {

    @Autowired
    private CategoryService categoryService;

    @Operation(summary = "获取分类列表")
    @GetMapping("/list")
    public Result<List<CategoryVO>> getCategoryList() {
        return Result.ok(categoryService.getCategoryList());
    }

    @Operation(summary = "获取分类详情")
    @GetMapping("/{id}")
    public Result<CategoryVO> getCategoryById(@PathVariable Long id) {
        return Result.ok(categoryService.getCategoryById(id));
    }
}
