package com.blog.service;

import com.baomidou.mybatisplus.extension.plugins.pagination.Page;
import com.baomidou.mybatisplus.extension.service.IService;
import com.blog.dto.CommentDTO;
import com.blog.entity.Comment;
import com.blog.vo.CommentVO;

public interface CommentService extends IService<Comment> {

    Page<CommentVO> getCommentList(Long articleId, Integer pageNum, Integer pageSize);

    void addComment(CommentDTO commentDTO);

    void updateComment(CommentDTO commentDTO);

    void deleteComment(Long id);
}
