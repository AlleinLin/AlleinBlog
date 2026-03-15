<div align="center">

# 🚀 Blog System - Go Version

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)](https://golang.org/)
[![Gin](https://img.shields.io/badge/Gin-1.9.1-brightgreen.svg)](https://gin-gonic.com/)
[![GORM](https://img.shields.io/badge/GORM-1.25.5-blue.svg)](https://gorm.io/)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

一个基于 Gin 框架的高性能博客系统后端

[功能特性](#-功能特性) • [快速开始](#-快速开始) • [API文档](#-api文档) • [项目结构](#-项目结构)

</div>

---

## ✨ 功能特性

- 🔐 **用户认证** - JWT Token 认证，Redis 缓存用户会话
- 📝 **文章管理** - 支持草稿/发布状态、Markdown 内容、文章摘要
- 🏷️ **分类标签** - 文章分类与标签管理，支持多标签关联
- 💬 **评论系统** - 嵌套评论，支持回复功能
- 🔒 **权限控制** - 基于中间件的认证授权
- 📊 **访问统计** - 文章阅读量统计，Redis 缓存优化
- 📅 **归档功能** - 按日期归档文章
- ⚡ **高性能** - Go 原生并发，轻量级协程

## 🛠️ 技术栈

| 技术 | 版本 | 说明 |
|------|------|------|
| Go | 1.21+ | 编程语言 |
| Gin | 1.9.1 | Web 框架 |
| GORM | 1.25.5 | ORM 框架 |
| JWT | 5.2.0 | Token 认证 |
| Redis | 9.3.0 | 缓存数据库 |
| MySQL | 8.0+ | 关系型数据库 |
| Viper | 1.18.2 | 配置管理 |
| Bcrypt | - | 密码加密 |

## 📦 项目结构

```
blog-go/
├── go.mod                              # Go 模块定义
├── go.sum                              # 依赖锁定
├── config.yaml                         # 配置文件
├── main.go                             # 程序入口
│
├── config/                             # 配置加载
│   └── config.go                       # 配置结构体与加载逻辑
│
├── model/                              # 数据模型
│   ├── entity.go                       # 数据库实体
│   ├── dto.go                          # 数据传输对象
│   ├── vo.go                           # 视图对象
│   └── page.go                         # 分页模型
│
├── database/                           # 数据库连接
│   ├── mysql.go                        # MySQL 连接
│   └── redis.go                        # Redis 连接
│
├── middleware/                         # 中间件
│   ├── jwt.go                          # JWT 认证中间件
│   └── common.go                       # CORS、日志中间件
│
├── controller/                         # 控制器
│   ├── user_controller.go              # 用户控制器
│   ├── article_controller.go           # 文章控制器
│   ├── category_controller.go          # 分类控制器
│   ├── tag_controller.go               # 标签控制器
│   └── comment_controller.go           # 评论控制器
│
├── service/                            # 业务服务
│   ├── user_service.go                 # 用户服务
│   ├── article_service.go              # 文章服务
│   ├── category_service.go             # 分类服务
│   ├── tag_service.go                  # 标签服务
│   └── comment_service.go              # 评论服务
│
├── router/                             # 路由配置
│   └── router.go                       # 路由注册
│
├── response/                           # 响应封装
│   └── response.go                     # 统一响应格式
│
├── constants/                          # 常量定义
│   └── constants.go                    # 系统常量、错误码
│
└── utils/                              # 工具函数
    ├── jwt.go                          # JWT 工具
    ├── redis.go                        # Redis 工具
    └── password.go                     # 密码加密工具
```

## 🚀 快速开始

### 环境要求

- Go 1.21+
- MySQL 8.0+
- Redis 6.0+

### 安装步骤

1. **克隆项目**
   ```bash
   git clone <repository-url>
   cd blog-go
   ```

2. **创建数据库**
   ```sql
   -- 使用 blog-java 的 schema.sql 或手动创建
   CREATE DATABASE blog_system DEFAULT CHARACTER SET utf8mb4;
   ```

3. **修改配置**
   
   编辑 `config.yaml`：
   ```yaml
   server:
     port: 8082
     mode: debug

   database:
     host: localhost
     port: 3306
     username: root
     password: your_password
     dbname: blog_system
     charset: utf8mb4

   redis:
     host: localhost
     port: 6379
     password: ""
     db: 0

   jwt:
     secret: your-jwt-secret-key
     expire: 24h
   ```

4. **安装依赖**
   ```bash
   go mod tidy
   ```

5. **启动项目**
   ```bash
   go run main.go
   ```

6. **访问应用**
   - 应用地址: http://localhost:8082

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

**登录示例：**
```bash
curl -X POST http://localhost:8082/login \
  -H "Content-Type: application/json" \
  -d '{"userName":"admin","password":"123456"}'
```

**响应示例：**
```json
{
  "code": 200,
  "msg": "操作成功",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "userInfo": {
      "id": 1,
      "userName": "admin",
      "nickName": "管理员",
      "avatar": "https://..."
    }
  }
}
```

### 文章相关

| 接口 | 方法 | 描述 | 认证 |
|------|------|------|------|
| `/article/list` | GET | 文章列表 | ❌ |
| `/article/hot` | GET | 热门文章 | ❌ |
| `/article/:id` | GET | 文章详情 | ❌ |
| `/article` | POST | 新增文章 | ✅ |
| `/article` | PUT | 更新文章 | ✅ |
| `/article/:id` | DELETE | 删除文章 | ✅ |
| `/article/archive` | GET | 归档列表 | ❌ |

**获取文章列表示例：**
```bash
curl "http://localhost:8082/article/list?pageNum=1&pageSize=10"
```

**新增文章示例：**
```bash
curl -X POST http://localhost:8082/article \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{
    "title": "我的第一篇博客",
    "content": "# Hello World\n\n这是我的第一篇博客...",
    "category": "Go",
    "tags": ["Go", "Gin"],
    "isDraft": false
  }'
```

### 分类标签

| 接口 | 方法 | 描述 | 认证 |
|------|------|------|------|
| `/category/list` | GET | 分类列表 | ❌ |
| `/category/:id` | GET | 分类详情 | ❌ |
| `/tag/list` | GET | 标签列表 | ❌ |
| `/tag/:id` | GET | 标签详情 | ❌ |

### 评论相关

| 接口 | 方法 | 描述 | 认证 |
|------|------|------|------|
| `/comment/list` | GET | 评论列表 | ❌ |
| `/comment` | POST | 发表评论 | ✅ |
| `/comment` | PUT | 更新评论 | ✅ |
| `/comment/:id` | DELETE | 删除评论 | ✅ |

## 📊 数据模型

### 核心实体

```go
type User struct {
    ID         uint64    `gorm:"primaryKey"`
    UserName   string    `gorm:"uniqueIndex;size:64"`
    NickName   string    `gorm:"size:64"`
    Password   string    `gorm:"size:128"`
    Email      string    `gorm:"size:64"`
    Avatar     string    `gorm:"size:256"`
    Type       string    `gorm:"size:1;default:0"`  // 0:普通用户 1:管理员
    Status     string    `gorm:"size:1;default:0"`  // 0:正常 1:停用
    CreateTime time.Time `gorm:"autoCreateTime"`
}

type Article struct {
    ID         uint64    `gorm:"primaryKey"`
    Title      string    `gorm:"size:256"`
    Content    string    `gorm:"type:longtext"`
    Summary    string    `gorm:"size:1024"`
    CategoryID uint64
    Thumbnail  string    `gorm:"size:256"`
    IsTop      string    `gorm:"size:1;default:0"`  // 是否置顶
    Status     string    `gorm:"size:1;default:0"`  // 0:发布 1:草稿
    ViewCount  uint64    `gorm:"default:0"`
    IsComment  string    `gorm:"size:1;default:1"`  // 是否允许评论
    CreateTime time.Time `gorm:"autoCreateTime"`
}
```

## 🔒 认证机制

### JWT Token

```go
type Claims struct {
    UserID uint64 `json:"userId"`
    jwt.RegisteredClaims
}
```

### 使用方式

```bash
# 在请求头中携带 Token
curl -H "Authorization: Bearer <token>" http://localhost:8082/user/info
```

## 📝 开发说明

### 添加新接口

1. 在 `model/` 中定义实体、DTO、VO
2. 在 `service/` 中实现业务逻辑
3. 在 `controller/` 中定义控制器
4. 在 `router/router.go` 中注册路由

### 代码规范

- 遵循 Go 官方代码规范
- 使用 `gofmt` 格式化代码
- 错误处理使用 `error` 返回
- 统一使用 `response` 包封装响应

### 日志规范

```go
log.Printf("[INFO] 操作描述: %v", param)
log.Printf("[ERROR] 错误描述: %v", err)
```

## 🔧 配置说明

```yaml
server:
  port: 8082          # 服务端口
  mode: debug         # 运行模式: debug/release

database:
  host: localhost     # 数据库地址
  port: 3306          # 数据库端口
  username: root      # 用户名
  password: root      # 密码
  dbname: blog_system # 数据库名
  charset: utf8mb4    # 字符集

redis:
  host: localhost     # Redis 地址
  port: 6379          # Redis 端口
  password: ""        # Redis 密码
  db: 0               # 数据库索引

jwt:
  secret: your-secret-key  # JWT 密钥
  expire: 24h              # Token 过期时间
```

## 🚀 部署

### 编译

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o blog-go

# Windows
GOOS=windows GOARCH=amd64 go build -o blog-go.exe

# macOS
GOOS=darwin GOARCH=amd64 go build -o blog-go
```

### Docker

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o blog-go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/blog-go .
COPY --from=builder /app/config.yaml .
EXPOSE 8082
CMD ["./blog-go"]
```

```bash
docker build -t blog-go .
docker run -d -p 8082:8082 blog-go
```

## 📄 License

[MIT License](LICENSE)

---

<div align="center">

Made with ❤️ by Blog Team

</div>
