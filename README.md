<div align="center">

# 📝 Blog System

[![Java](https://img.shields.io/badge/Java-17+-ED8B00.svg)](https://www.oracle.com/java/)
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)](https://golang.org/)
[![Spring Boot](https://img.shields.io/badge/Spring%20Boot-3.2.0-brightgreen.svg)](https://spring.io/projects/spring-boot)
[![Gin](https://img.shields.io/badge/Gin-1.9.1-brightgreen.svg)](https://gin-gonic.com/)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

一个现代化的博客系统，提供 **Java** 和 **Go** 两种技术栈实现

[功能特性](#-功能特性) • [项目结构](#-项目结构) • [快速开始](#-快速开始) • [技术对比](#-技术对比)

</div>

---

## 📖 项目简介

本项目是一个完整的博客系统后端，采用前后端分离架构，提供了 Java (Spring Boot) 和 Go (Gin) 两种实现版本。两个版本功能对等，可根据团队技术栈选择使用。

### 为什么提供两个版本？

- **技术选型灵活** - 根据团队技术栈选择合适的版本
- **性能对比参考** - 了解不同语言实现的性能差异
- **学习对比** - 对比学习两种技术栈的实现方式

## ✨ 功能特性

| 功能 | 描述 |
|------|------|
| 🔐 **用户认证** | JWT Token 认证，Redis 缓存用户会话 |
| 📝 **文章管理** | 支持草稿/发布状态、Markdown 内容、文章摘要 |
| 🏷️ **分类标签** | 文章分类与标签管理，支持多标签关联 |
| 💬 **评论系统** | 嵌套评论，支持回复功能 |
| 🔒 **权限控制** | 基于角色的访问控制 (RBAC) |
| 📊 **访问统计** | 文章阅读量统计，Redis 缓存优化 |
| 📅 **归档功能** | 按日期归档文章 |
| ⚡ **高性能** | Redis 缓存 + 数据库优化 |

## 📁 项目结构

```
blog/
├── blog-java/                    # Java 版本 (Spring Boot)
│   ├── blog-framework/           # 核心框架模块
│   │   └── src/main/java/com/blog/
│   │       ├── entity/           # 实体类
│   │       ├── dto/              # 数据传输对象
│   │       ├── vo/               # 视图对象
│   │       ├── mapper/           # MyBatis Mapper
│   │       ├── service/          # 业务服务
│   │       ├── config/           # 配置类
│   │       ├── security/         # 安全认证
│   │       ├── exception/        # 异常处理
│   │       └── util/             # 工具类
│   ├── blog-web/                 # Web 应用模块
│   │   └── src/main/
│   │       ├── java/com/blog/controller/  # 控制器
│   │       └── resources/
│   │           ├── application.yml        # 配置文件
│   │           └── schema.sql             # 数据库脚本
│   └── pom.xml
│
├── blog-go/                      # Go 版本 (Gin)
│   ├── config/                   # 配置加载
│   ├── model/                    # 数据模型
│   │   ├── entity.go             # 数据库实体
│   │   ├── dto.go                # 数据传输对象
│   │   ├── vo.go                 # 视图对象
│   │   └── page.go               # 分页模型
│   ├── database/                 # 数据库连接
│   ├── middleware/               # 中间件
│   ├── controller/               # 控制器
│   ├── service/                  # 业务服务
│   ├── router/                   # 路由配置
│   ├── response/                 # 响应封装
│   ├── constants/                # 常量定义
│   ├── utils/                    # 工具函数
│   ├── main.go                   # 程序入口
│   ├── config.yaml               # 配置文件
│   └── go.mod
│
└── README.md                     # 项目说明文档
```

## 🚀 快速开始

### 环境要求

| 组件 | 版本要求 |
|------|----------|
| MySQL | 8.0+ |
| Redis | 6.0+ |

### Java 版本启动

```bash
# 进入 Java 项目目录
cd blog-java

# 安装依赖
mvn clean install

# 启动服务
mvn spring-boot:run -pl blog-web

# 访问
# 应用地址: http://localhost:8081
# API 文档: http://localhost:8081/doc.html
```

### Go 版本启动

```bash
# 进入 Go 项目目录
cd blog-go

# 安装依赖
go mod tidy

# 启动服务
go run main.go

# 访问
# 应用地址: http://localhost:8082
```

### 数据库初始化

```sql
-- 创建数据库
CREATE DATABASE blog_system DEFAULT CHARACTER SET utf8mb4;

-- 执行建表脚本 (Java 版本)
-- blog-java/blog-web/src/main/resources/schema.sql
```

### 默认账号

| 用户名 | 密码 | 角色 |
|--------|------|------|
| admin | 123456 | 管理员 |

## ⚖️ 技术对比

### 技术栈

| 组件 | Java 版本 | Go 版本 |
|------|-----------|---------|
| 语言 | Java 17+ | Go 1.21+ |
| Web 框架 | Spring Boot 3.2.0 | Gin 1.9.1 |
| ORM | MyBatis-Plus 3.5.5 | GORM 1.25.5 |
| 认证 | Spring Security + JWT | JWT Middleware |
| 配置管理 | application.yml | Viper |
| API 文档 | Knife4j | - |
| 工具库 | Hutool | - |

### 特性对比

| 特性 | Java 版本 | Go 版本 |
|------|-----------|---------|
| RBAC 权限模型 | ✅ 完整实现 | ⚠️ 基础实现 |
| API 文档 | ✅ Knife4j | ❌ |
| 模块化设计 | ✅ 多模块 | ❌ 单模块 |
| 编译产物大小 | ~50MB | ~15MB |
| 启动速度 | ~3s | ~0.1s |
| 内存占用 | ~200MB | ~30MB |

### 性能建议

- **追求开发效率** → 选择 Java 版本，生态完善，开发效率高
- **追求极致性能** → 选择 Go 版本，资源占用低，响应速度快
- **企业级应用** → 选择 Java 版本，权限模型完善
- **个人博客/小项目** → 选择 Go 版本，部署简单

## 📖 API 文档

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

## 🗄️ 数据库设计

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

## 📦 部署

### Docker 部署

```bash
# Java 版本
cd blog-java
docker build -t blog-java .
docker run -d -p 8081:8081 blog-java

# Go 版本
cd blog-go
docker build -t blog-go .
docker run -d -p 8082:8082 blog-go
```

### Docker Compose (推荐)

```yaml
version: '3.8'
services:
  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: blog_system
    ports:
      - "3306:3306"
    volumes:
      - ./mysql-data:/var/lib/mysql

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"

  blog-java:
    build: ./blog-java
    ports:
      - "8081:8081"
    depends_on:
      - mysql
      - redis

  blog-go:
    build: ./blog-go
    ports:
      - "8082:8082"
    depends_on:
      - mysql
      - redis
```

## 📄 License

[MIT License](LICENSE)

---

<div align="center">

Made with ❤️ by Blog Team

</div>
