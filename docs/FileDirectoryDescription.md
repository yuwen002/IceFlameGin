# 项目目录结构说明

本文件描述 Ice Falme 项目的目录结构及其职责。

```
.
├── cmd
│   └── configuration.go - 应用启动入口及配置初始化
├── config
│   ├── autoload - 配置自动加载模块
│   │   ├── template.loader.go - 模板配置加载
│   │   ├── use.loader.go - 使用配置加载
│   │   └── validator.loader.go - 验证器配置加载
│   ├── config.development.yaml - 开发环境配置文件
│   ├── config.go - 配置管理核心代码
│   └── config.yaml - 默认配置文件
├── hack
│   └── gen.yml - 代码生成配置文件（用于生成models）
├── internal
│   └── app - 核心业务逻辑
│       ├── constants - 系统常量定义
│       ├── controllers - MVC 控制器层
│       │   └── admin - 管理员相关控制器
│       ├── db - 数据库相关代码
│       │   ├── core.go - 数据库核心功能
│       │   ├── db.go - 数据库连接管理
│       │   └── query_options.go - 查询选项定义
│       ├── dto - 数据传输对象（Data Transfer Object）
│       │   ├── article.dto.go - 文章相关数据传输对象
│       │   ├── public.dto.go - 公共数据传输对象
│       │   └── ... - 其他模块数据传输对象
│       ├── middlewares - HTTP 中间件
│       │   └── auth_middleware.go - 身份验证中间件
│       ├── models - 数据模型定义
│       │   ├── association - 模型关联关系定义
│       │   ├── model - 数据库模型定义（程序自动生成）
│       │   └── query - 查询构建器（程序自动生成）
│       ├── repositories - 数据访问层（Repository Pattern）
│       │   ├── article.gen_repo.go - 文章数据访问（程序自动生成）
│       │   ├── article_category.repo.go - 文章分类数据访问
│       │   └── ... - 其他实体数据访问实现
│       ├── services - 业务逻辑层
│       │   ├── article.svc.go - 文章业务逻辑
│       │   └── ... - 其他模块业务逻辑
│       └── validators - 请求数据验证
│           ├── admin.validator.go - 管理员相关验证
│           └── ... - 其他模块验证逻辑
├── migrations
│   └── auto_migrate.go - 数据库自动迁移脚本
├── routers
│   ├── paths - 路由路径常量定义
│   │   └── admin_paths.go - 管理员相关路由路径
│   ├── admin.router.go - 管理员路由配置
│   └── router.go - 路由初始化
└── web - Web 相关资源
    ├── AdminLTE-3.2.0 - 管理后台模板
    ├── static - 静态资源文件
    └── templates - 页面模板文件
```

## 目录职责说明

- **cmd**: 应用启动入口及配置初始化
- **config**: 配置文件和加载器
- **hack**: 代码生成配置文件（如生成models）
- **internal/app**: 核心业务逻辑，采用分层架构设计
  - **constants**: 系统常量定义
  - **controllers**: MVC 控制器层，处理 HTTP 请求
  - **db**: 数据库相关功能
  - **dto**: 数据传输对象，用于层间数据传递
  - **middlewares**: HTTP 中间件，如身份验证
  - **models**: 数据模型定义
  - **repositories**: 数据访问层，实现 Repository 模式
  - **services**: 业务逻辑层
  - **validators**: 请求数据验证
- **migrations**: 数据库迁移文件
- **routers**: 路由配置
- **web**: Web 相关资源，包括模板和静态文件