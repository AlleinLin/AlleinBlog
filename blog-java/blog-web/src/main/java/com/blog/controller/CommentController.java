package com.blog.controller;

import com.baomidou.mybatisplus.extension.plugins.pagination.Page;
import com.blog.common.Result;
import com.blog.dto.CommentDTO;
import com.blog.service.CommentService;
import com.blog.vo.CommentVO;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@Tag(name = "评论管理")
@RestController
@RequestMapping("/comment")
public class CommentController {

    @Autowired
    private CommentService commentService;

    @Operation(summary = "获取评论列表")
    @GetMapping("/list")
    public Result<Page<CommentVO>> getCommentList(
            @RequestParam Long articleId,
            @RequestParam(defaultValue = "1") Integer pageNum,
            @RequestParam(defaultValue = "10") Integer pageSize) {
        return Result.ok(commentService.getCommentList(articleId, pageNum, pageSize));
    }

    @Operation(summary = "添加评论")
    @PostMapping
    public Result<Void> addComment(@RequestBody CommentDTO commentDTO) {
        commentService.addComment(commentDTO);
        return Result.ok();
    }

    @Operation(summary = "更新评论")
    @PutMapping
    public Result<Void> updateComment(@RequestBody CommentDTO commentDTO) {
        commentService.updateComment(commentDTO);
        return Result.ok();
    }

    @Operation(summary = "删除评论")
    @DeleteMapping("/{id}")
    public Result<Void> deleteComment(@PathVariable Long id) {
        commentService.deleteComment(id);
        return Result.ok();
    }
}
