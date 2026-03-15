package com.blog.service.impl;

import com.baomidou.mybatisplus.core.conditions.query.LambdaQueryWrapper;
import com.baomidou.mybatisplus.extension.plugins.pagination.Page;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.blog.dto.CommentDTO;
import com.blog.entity.Comment;
import com.blog.entity.User;
import com.blog.exception.BusinessException;
import com.blog.mapper.CommentMapper;
import com.blog.security.LoginUser;
import com.blog.service.CommentService;
import com.blog.service.UserService;
import com.blog.util.BeanCopyUtil;
import com.blog.vo.CommentVO;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

@Service
public class CommentServiceImpl extends ServiceImpl<CommentMapper, Comment> implements CommentService {

    @Autowired
    private UserService userService;

    @Override
    public Page<CommentVO> getCommentList(Long articleId, Integer pageNum, Integer pageSize) {
        LambdaQueryWrapper<Comment> wrapper = new LambdaQueryWrapper<>();
        wrapper.eq(Comment::getArticleId, articleId)
                .isNull(Comment::getParentId)
                .orderByDesc(Comment::getCreateTime);
        Page<Comment> page = new Page<>(pageNum, pageSize);
        page(page, wrapper);
        List<Comment> comments = page.getRecords();
        List<CommentVO> voList = convertToVOList(comments);
        voList.forEach(vo -> {
            LambdaQueryWrapper<Comment> childWrapper = new LambdaQueryWrapper<>();
            childWrapper.eq(Comment::getParentId, vo.getId())
                    .orderByAsc(Comment::getCreateTime);
            List<Comment> children = list(childWrapper);
            vo.setChildren(convertToVOList(children));
        });
        Page<CommentVO> voPage = new Page<>(page.getCurrent(), page.getSize(), page.getTotal());
        voPage.setRecords(voList);
        return voPage;
    }

    @Override
    public void addComment(CommentDTO commentDTO) {
        Comment comment = BeanCopyUtil.copyBean(commentDTO, Comment.class);
        save(comment);
    }

    @Override
    public void updateComment(CommentDTO commentDTO) {
        Comment comment = getById(commentDTO.getId());
        if (comment == null) {
            throw new BusinessException("评论不存在");
        }
        LoginUser loginUser = (LoginUser) SecurityContextHolder.getContext().getAuthentication().getPrincipal();
        if (!comment.getCreateBy().equals(loginUser.getUser().getId())) {
            throw new BusinessException("无权修改他人评论");
        }
        comment.setContent(commentDTO.getContent());
        updateById(comment);
    }

    @Override
    public void deleteComment(Long id) {
        Comment comment = getById(id);
        if (comment == null) {
            throw new BusinessException("评论不存在");
        }
        LoginUser loginUser = (LoginUser) SecurityContextHolder.getContext().getAuthentication().getPrincipal();
        if (!comment.getCreateBy().equals(loginUser.getUser().getId())) {
            throw new BusinessException("无权删除他人评论");
        }
        removeById(id);
    }

    private List<CommentVO> convertToVOList(List<Comment> comments) {
        if (comments.isEmpty()) {
            return new ArrayList<>();
        }
        List<Long> userIds = comments.stream()
                .map(Comment::getCreateBy)
                .distinct()
                .collect(Collectors.toList());
        Map<Long, User> userMap = userService.listByIds(userIds).stream()
                .collect(Collectors.toMap(User::getId, user -> user));
        return comments.stream().map(comment -> {
            CommentVO vo = BeanCopyUtil.copyBean(comment, CommentVO.class);
            User user = userMap.get(comment.getCreateBy());
            if (user != null) {
                vo.setUserName(user.getUserName());
                vo.setNickName(user.getNickName());
                vo.setAvatar(user.getAvatar());
            }
            if (comment.getToUserId() != null) {
                User toUser = userMap.get(comment.getToUserId());
                if (toUser != null) {
                    vo.setToUserName(toUser.getUserName());
                    vo.setToNickName(toUser.getNickName());
                }
            }
            return vo;
        }).collect(Collectors.toList());
    }
}
