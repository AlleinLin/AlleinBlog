package com.blog.controller;

import com.blog.common.Result;
import com.blog.service.TagService;
import com.blog.vo.TagVO;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@Tag(name = "标签管理")
@RestController
@RequestMapping("/tag")
public class TagController {

    @Autowired
    private TagService tagService;

    @Operation(summary = "获取标签列表")
    @GetMapping("/list")
    public Result<List<TagVO>> getTagList() {
        return Result.ok(tagService.getTagList());
    }

    @Operation(summary = "获取标签详情")
    @GetMapping("/{id}")
    public Result<TagVO> getTagById(@PathVariable Long id) {
        return Result.ok(tagService.getTagById(id));
    }
}
