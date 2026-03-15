<div align="center">

# 🌸 Blog System - Java Version

[![Spring Boot](https://img.shields.io/badge/Spring%20Boot-3.2.0-brightgreen.svg)](https://spring.io/projects/spring-boot)
[![MyBatis-Plus](https://img.shields.io/badge/MyBatis%20Plus-3.5.5-blue.svg)](https://baomidou.com/)
[![JWT](https://img.shields.io/badge/JWT-0.12.3-orange.svg)](https://jwt.io/)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

一个基于 Spring Boot 3 的现代化博客系统后端

[功能特性](#-功能特性) • [快速开始](#-快速开始) • [API文档](#-api文档) • [项目结构](#-项目结构)

</div>

---

## ✨ 功能特性

- 🔐 **用户认证** - JWT Token 认证，Redis 缓存用户会话
- 📝 **文章管理** - 支持草稿/发布状态、Markdown 内容、文章摘要
- 🏷️ **分类标签** - 文章分类与标签管理，支持多标签关联
- 💬 **评论系统** - 嵌套评论，支持回复功能
- 🔒 **权限控制** - RBAC 权限模型，细粒度权限控制
- 📊 **访问统计** - 文章阅读量统计，Redis 缓存优化
- 📅 **归档功能** - 按日期归档文章
- 📖 **API文档** - 集成 Knife4j，在线 API 文档

## 🛠️ 技术栈

| 技术 | 版本 | 说明 |
|------|------|------|
| Spring Boot | 3.2.0 | 基础框架 |
| MyBatis-Plus | 3.5.5 | ORM 框架 |
| Spring Security | - | 安全框架 |
| JWT | 0.12.3 | Token 认证 |
| Redis | - | 缓存数据库 |
| MySQL | 8.0+ | 关系型数据库 |
| Knife4j | 4.4.0 | API 文档 |
| Lombok | - | 简化代码 |
| Hutool | 5.8.24 | 工具库 |

## 📦 项目结构

```
blog-java/
├── pom.xml                              # 父 POM，管理依赖版本
├── blog-framework/                      # 核心框架模块
│   ├── pom.xml
│   └── src/main/java/com/blog/
│       ├── entity/                      # 实体类
│       │   ├── Article.java             # 文章实体
│       │   ├── Category.java            # 分类实体
│       │   ├── Tag.java                 # 标签实体
│       │   ├── Comment.java             # 评论实体
│       │   ├── User.java                # 用户实体
│       │   ├── Role.java                # 角色实体
│       │   └── Access.java              # 权限实体
│       ├── dto/                         # 数据传输对象
│       │   ├── ArticleDTO.java          # 文章 DTO
│       │   ├── LoginDTO.java            # 登录 DTO
│       │   └── CommentDTO.java          # 评论 DTO
│       ├── vo/                          # 视图对象
│       │   ├── ArticleVO.java           # 文章 VO
│       │   ├── UserInfoVO.java          # 用户信息 VO
│       │   └── CommentVO.java           # 评论 VO
│       ├── mapper/                      # MyBatis Mapper 接口
│       ├── service/                     # 业务服务接口
│       │   └── impl/                    # 业务服务实现
│       ├── config/                      # 配置类
│       │   ├── SecurityConfig.java      # 安全配置
│       │   ├── RedisConfig.java         # Redis 配置
│       │   └── MybatisPlusConfig.java   # MyBatis-Plus 配置
│       ├── security/                    # 安全认证
│       │   ├── LoginUser.java           # 登录用户信息
│       │   ├── JwtAuthenticationFilter.java  # JWT 过滤器
│       │   └── AuthenticationEntryPointImpl.java
│       ├── exception/                   # 异常处理
│       │   ├── BusinessException.java   # 业务异常
│       │   └── GlobalExceptionHandler.java  # 全局异常处理
│       ├── common/                      # 公共类
│       │   ├── Result.java              # 统一响应
│       │   ├── ErrorCode.java           # 错误码
│       │   └── Constants.java           # 常量
│       └── util/                        # 工具类
│           ├── JwtUtil.java             # JWT 工具
│           ├── RedisCache.java          # Redis 缓存工具
│           └── BeanCopyUtil.java        # Bean 拷贝工具
└── blog-web/                            # Web 应用模块
    ├── pom.xml
    └── src/main/
        ├── java/com/blog/
        │   ├── controller/              # 控制器
        │   │   ├── ArticleController.java
        │   │   ├── CategoryController.java
        │   │   ├── TagController.java
        │   │   ├── CommentController.java
        │   │   └── UserController.java
        │   └── BlogApplication.java     # 启动类
        └── resources/
            ├── application.yml          # 配置文件
            └── schema.sql               # 数据库脚本
```

## 🚀 快速开始

### 环境要求

- JDK 17+
- Maven 3.6+
- MySQL 8.0+
- Redis 6.0+

### 安装步骤

1. **克隆项目**
   ```bash
   git clone <repository-url>
   cd blog-java
   ```

2. **创建数据库**
   ```sql
   -- 执行 blog-web/src/main/resources/schema.sql
   ```

3. **修改配置**
   
   编辑 `blog-web/src/main/resources/application.yml`：
   ```yaml
   spring:
     datasource:
       url: jdbc:mysql://localhost:3306/blog_system
       username: your_username
       password: your_password
     data:
       redis:
         host: localhost
         port: 6379
         password: your_redis_password
   ```

4. **启动项目**
   ```bash
   mvn clean install
   mvn spring-boot:run -pl blog-web
   ```

5. **访问应用**
   - 应用地址: http://localhost:8081
   - API 文档: http://localhost:8081/doc.html

### 默认账号

| 用户名 | 密码 | 角色 |
|--------|------|------|
| admin | 123456 | 管理员 |

## 📖 API文档

### 认证相关

| 接口 | 方法 | 描述 | 认证 |
|------|------|------|------|
| `/login` | POST | 用户登录 | ❌ |
| `/register` | POST | 用户注册 | ❌ |
| `/logout` | POST | 用户登出 | ✅ |

### 文章相关

| 接口 | 方法 | 描述 | 认证 |
|------|------|------|------|
| `/article/list` | GET | 文章列表 | ❌ |
| `/article/hot` | GET | 热门文章 | ❌ |
| `/article/{id}` | GET | 文章详情 | ❌ |
| `/article` | POST | 新增文章 | ✅ |
| `/article` | PUT | 更新文章 | ✅ |
| `/article/{id}` | DELETE | 删除文章 | ✅ |
| `/article/archive` | GET | 归档列表 | ❌ |

### 分类标签

| 接口 | 方法 | 描述 | 认证 |
|------|------|------|------|
| `/category/list` | GET | 分类列表 | ❌ |
| `/tag/list` | GET | 标签列表 | ❌ |

### 评论相关

| 接口 | 方法 | 描述 | 认证 |
|------|------|------|------|
| `/comment/list` | GET | 评论列表 | ❌ |
| `/comment` | POST | 发表评论 | ✅ |
| `/comment/{id}` | DELETE | 删除评论 | ✅ |

## 📊 数据库设计

### ER 图

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│    User     │────<│  UserRole   │>────│    Role     │
└─────────────┘     └─────────────┘     └─────────────┘
       │                                       │
       │                                 ┌─────────────┐
       │                                 │ RoleAccess  │
       │                                 └─────────────┘
       │                                       │
       │                                 ┌─────────────┐
       │                                 │   Access    │
       │                                 └─────────────┘

┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│   Article   │────<│ ArticleTag  │>────│    Tag      │
└─────────────┘     └─────────────┘     └─────────────┘
       │
       │           ┌─────────────┐
       └──────────>│  Category   │
                   └─────────────┘
       │
       │           ┌─────────────┐
       └──────────>│  Comment    │
                   └─────────────┘
```

## 🔒 权限设计

采用 RBAC 权限模型：

```
用户 ── 用户角色关联 ──> 角色 ── 角色权限关联 ──> 权限
```

### 权限标识

| 权限 | 标识 | 描述 |
|------|------|------|
| 发表博客 | `article:add` | 发表新文章 |
| 编辑博客 | `article:edit` | 编辑文章 |
| 删除博客 | `article:delete` | 删除文章 |
| 查看博客 | `article:view` | 查看文章 |

## 📝 开发说明

### 代码规范

- 遵循阿里巴巴 Java 开发手册
- 使用 Lombok 简化代码
- 统一使用 `Result` 封装响应
- 异常使用 `BusinessException` 抛出

### 日志规范

```java
@Slf4j
@Service
public class ArticleServiceImpl {
    public void method() {
        log.info("操作描述: {}", param);
        log.error("错误描述: ", exception);
    }
}
```

## 📄 License

[MIT License](LICENSE)

---

<div align="center">

Made with ❤️ by Blog Team

</div>
