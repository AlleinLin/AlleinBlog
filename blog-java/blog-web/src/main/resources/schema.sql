-- 创建数据库
CREATE DATABASE IF NOT EXISTS blog_system DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE blog_system;

-- 用户表
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键',
    `user_name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '用户名',
    `nick_name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '昵称',
    `signature` VARCHAR(128) DEFAULT '这个人很懒，什么都没写' COMMENT '个性签名',
    `password` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '密码',
    `type` CHAR(1) DEFAULT '0' COMMENT '用户类型：0普通用户，1管理员',
    `status` CHAR(1) DEFAULT '0' COMMENT '账号状态（0正常 1停用）',
    `email` VARCHAR(64) DEFAULT NULL COMMENT '邮箱',
    `phonenumber` VARCHAR(32) DEFAULT NULL COMMENT '手机号',
    `sex` CHAR(1) DEFAULT '0' COMMENT '性别（0男，1女，2未知）',
    `avatar` VARCHAR(256) DEFAULT NULL COMMENT '头像',
    `create_by` BIGINT DEFAULT NULL COMMENT '创建人',
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_by` BIGINT DEFAULT NULL COMMENT '更新人',
    `update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `del_flag` INT DEFAULT 0 COMMENT '删除标志（0未删除 1已删除）',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_name` (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 角色表
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(128) DEFAULT NULL COMMENT '角色名称',
    `role_key` VARCHAR(100) DEFAULT NULL COMMENT '角色权限字符串',
    `status` CHAR(1) DEFAULT '0' COMMENT '角色状态（0正常 1停用）',
    `del_flag` INT DEFAULT 0 COMMENT '删除标志',
    `create_by` BIGINT DEFAULT NULL,
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `update_by` BIGINT DEFAULT NULL,
    `update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `remark` VARCHAR(500) DEFAULT NULL COMMENT '备注',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';

-- 权限表
DROP TABLE IF EXISTS `access`;
CREATE TABLE `access` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `access_name` VARCHAR(64) NOT NULL COMMENT '权限名',
    `permission` VARCHAR(100) NOT NULL COMMENT '权限标识',
    `status` CHAR(1) DEFAULT '0' COMMENT '权限状态（0正常 1停用）',
    `create_by` BIGINT DEFAULT NULL,
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `update_by` BIGINT DEFAULT NULL,
    `update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `del_flag` INT DEFAULT 0 COMMENT '是否删除（0未删除 1已删除）',
    `remark` VARCHAR(500) DEFAULT NULL COMMENT '备注',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='权限表';

-- 用户角色关联表
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role` (
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `role_id` BIGINT NOT NULL COMMENT '角色ID',
    PRIMARY KEY (`user_id`, `role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户角色关联表';

-- 角色权限关联表
DROP TABLE IF EXISTS `role_access`;
CREATE TABLE `role_access` (
    `role_id` BIGINT NOT NULL COMMENT '角色ID',
    `access_id` BIGINT NOT NULL COMMENT '权限ID',
    PRIMARY KEY (`role_id`, `access_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色权限关联表';

-- 分类表
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(128) DEFAULT NULL COMMENT '分类名',
    `pid` BIGINT DEFAULT -1 COMMENT '父分类ID',
    `description` VARCHAR(512) DEFAULT NULL COMMENT '描述',
    `status` CHAR(1) DEFAULT '0' COMMENT '状态（0正常 1禁用）',
    `create_by` BIGINT DEFAULT NULL,
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `update_by` BIGINT DEFAULT NULL,
    `update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `del_flag` INT DEFAULT 0 COMMENT '删除标志',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='分类表';

-- 标签表
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(128) DEFAULT NULL COMMENT '标签名',
    `create_by` BIGINT DEFAULT NULL,
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `update_by` BIGINT DEFAULT NULL,
    `update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `del_flag` INT DEFAULT 0 COMMENT '删除标志',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='标签表';

-- 文章表
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(256) DEFAULT NULL COMMENT '标题',
    `content` LONGTEXT COMMENT '文章内容',
    `summary` VARCHAR(1024) DEFAULT NULL COMMENT '文章摘要',
    `category_id` BIGINT DEFAULT NULL COMMENT '所属分类ID',
    `thumbnail` VARCHAR(256) DEFAULT NULL COMMENT '缩略图',
    `is_top` CHAR(1) DEFAULT '0' COMMENT '是否置顶（0否 1是）',
    `status` CHAR(1) DEFAULT '0' COMMENT '状态（0已发布 1草稿）',
    `view_count` BIGINT DEFAULT 0 COMMENT '访问量',
    `is_comment` CHAR(1) DEFAULT '1' COMMENT '是否允许评论（0否 1是）',
    `create_by` BIGINT DEFAULT NULL,
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `update_by` BIGINT DEFAULT NULL,
    `update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `del_flag` INT DEFAULT 0 COMMENT '删除标志',
    PRIMARY KEY (`id`),
    KEY `idx_category_id` (`category_id`),
    KEY `idx_create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='文章表';

-- 文章标签关联表
DROP TABLE IF EXISTS `article_tag`;
CREATE TABLE `article_tag` (
    `article_id` BIGINT NOT NULL COMMENT '文章ID',
    `tag_id` BIGINT NOT NULL COMMENT '标签ID',
    PRIMARY KEY (`article_id`, `tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='文章标签关联表';

-- 评论表
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `article_id` BIGINT DEFAULT NULL COMMENT '文章ID',
    `parent_id` BIGINT DEFAULT NULL COMMENT '父评论ID',
    `content` VARCHAR(1024) DEFAULT NULL COMMENT '评论内容',
    `to_user_id` BIGINT DEFAULT NULL COMMENT '回复用户ID',
    `create_by` BIGINT DEFAULT NULL,
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `update_by` BIGINT DEFAULT NULL,
    `update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `del_flag` INT DEFAULT 0 COMMENT '删除标志',
    PRIMARY KEY (`id`),
    KEY `idx_article_id` (`article_id`),
    KEY `idx_parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='评论表';

-- 初始化数据
-- 插入管理员用户（密码：123456，使用BCrypt加密）
INSERT INTO `user` (`id`, `user_name`, `nick_name`, `signature`, `password`, `type`, `status`, `email`, `avatar`) VALUES
(1, 'admin', '管理员', '欢迎来到我的博客', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iAt6Z5EH', '1', '0', 'admin@blog.com', 'https://api.dicebear.com/7.x/avataaars/svg?seed=admin');

-- 插入角色
INSERT INTO `role` (`id`, `name`, `role_key`, `status`, `remark`) VALUES
(1, '管理员', 'admin', '0', '系统管理员'),
(2, '普通用户', 'normal_user', '0', '普通用户');

-- 插入权限
INSERT INTO `access` (`id`, `access_name`, `permission`, `status`, `remark`) VALUES
(1, '发表博客', 'article:add', '0', '发表博客文章'),
(2, '删除博客', 'article:delete', '0', '删除博客文章'),
(3, '编辑博客', 'article:edit', '0', '编辑博客文章'),
(4, '查看博客', 'article:view', '0', '查看博客文章');

-- 插入用户角色关联
INSERT INTO `user_role` (`user_id`, `role_id`) VALUES (1, 1);

-- 插入角色权限关联
INSERT INTO `role_access` (`role_id`, `access_id`) VALUES 
(1, 1), (1, 2), (1, 3), (1, 4),
(2, 4);

-- 插入分类
INSERT INTO `category` (`id`, `name`, `pid`, `description`, `status`) VALUES
(1, 'Java', -1, 'Java相关技术', '0'),
(2, 'Go', -1, 'Go语言相关技术', '0'),
(3, '前端', -1, '前端开发技术', '0'),
(4, '数据库', -1, '数据库相关', '0'),
(5, '随笔', -1, '生活随笔', '0');

-- 插入标签
INSERT INTO `tag` (`id`, `name`) VALUES
(1, 'Spring Boot'),
(2, 'MyBatis'),
(3, 'Gin'),
(4, 'GORM'),
(5, 'Vue'),
(6, 'React'),
(7, 'MySQL'),
(8, 'Redis');
