package com.blog.controller;

import com.baomidou.mybatisplus.extension.plugins.pagination.Page;
import com.blog.common.Result;
import com.blog.dto.ArticleDTO;
import com.blog.dto.ArticleQueryDTO;
import com.blog.service.ArticleService;
import com.blog.vo.ArchiveVO;
import com.blog.vo.ArticleVO;
import com.blog.vo.HotArticleVO;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import jakarta.validation.Valid;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@Tag(name = "文章管理")
@RestController
@RequestMapping("/article")
public class ArticleController {

    @Autowired
    private ArticleService articleService;

    @Operation(summary = "获取热门文章列表")
    @GetMapping("/hot")
    public Result<List<HotArticleVO>> getHotArticleList() {
        return Result.ok(articleService.getHotArticleList());
    }

    @Operation(summary = "获取文章列表")
    @GetMapping("/list")
    public Result<Page<ArticleVO>> getArticleList(ArticleQueryDTO queryDTO) {
        return Result.ok(articleService.getArticleList(queryDTO));
    }

    @Operation(summary = "获取文章详情")
    @GetMapping("/{id}")
    public Result<ArticleVO> getArticleDetail(@PathVariable Long id) {
        articleService.updateViewCount(id);
        return Result.ok(articleService.getArticleDetail(id));
    }

    @Operation(summary = "添加文章")
    @PostMapping
    @PreAuthorize("hasAuthority('article:add')")
    public Result<Long> addArticle(@Valid @RequestBody ArticleDTO articleDTO) {
        return Result.ok(articleService.addArticle(articleDTO));
    }

    @Operation(summary = "更新文章")
    @PutMapping
    @PreAuthorize("hasAuthority('article:edit')")
    public Result<Void> updateArticle(@Valid @RequestBody ArticleDTO articleDTO) {
        articleService.updateArticle(articleDTO);
        return Result.ok();
    }

    @Operation(summary = "删除文章")
    @DeleteMapping("/{id}")
    @PreAuthorize("hasAuthority('article:delete')")
    public Result<Void> deleteArticle(@PathVariable Long id) {
        articleService.deleteArticle(id);
        return Result.ok();
    }

    @Operation(summary = "获取文章总数")
    @GetMapping("/count")
    public Result<Long> getArticleCount() {
        return Result.ok(articleService.getArticleCount());
    }

    @Operation(summary = "获取归档列表")
    @GetMapping("/archive")
    public Result<List<ArchiveVO>> getArchiveList(
            @RequestParam(defaultValue = "1") Integer pageNum,
            @RequestParam(defaultValue = "10") Integer pageSize) {
        return Result.ok(articleService.getArchiveList(pageNum, pageSize));
    }

    @Operation(summary = "获取上一篇/下一篇文章")
    @GetMapping("/previousNext/{id}")
    public Result<ArticleVO> getPreviousNextArticle(@PathVariable Long id) {
        return Result.ok(articleService.getPreviousNextArticle(id));
    }
}
