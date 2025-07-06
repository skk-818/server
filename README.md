# 📁 项目结构说明文档

> 项目语言：Go  
> 架构风格：Clean Architecture + Gin + Wire + Usecase 分层结构

---

## 🔧 根目录文件说明

| 文件/目录     | 描述 |
|---------------|------|
| `go.mod`      | Go module 配置文件，定义模块名和依赖项。 |
| `Makefile`    | 构建、运行、格式化、测试等常用命令的封装脚本。 |
| `cmd/`        | 启动入口，通常一个子目录对应一个服务。 |
| `etc/`        | 配置文件目录，例如 `config.yaml`、`.env` 等。 |
| `router/`     | HTTP 路由定义模块，定义路由与中间件挂载。 |

---

## 🧩 internal（核心业务代码）

> 按模块划分，每个模块遵循清晰的分层规范：`api` → `usecase` → `repo/service` → `model`

### 📁 internal/core

- 系统通用组件初始化，例如日志、配置、数据库、Redis 等。

### 📁 internal/di

- 依赖注入配置（使用 [Google Wire](https://github.com/google/wire)）
- 定义项目结构图及模块依赖关系的汇总。

### 📁 internal/middleware

- 所有自定义中间件：
    - JWT 校验
    - CORS 处理
    - 错误统一处理等

---

## 📁 internal/module 模块划分

### 📂 module/common

| 子目录      | 描述 |
|-------------|------|
| `api/`      | 通用的公共接口，例如上传、健康检查等。 |
| `model/`    | 通用的数据结构定义（DTO/VO/Entity）。 |
| `usecase/`  | 公共业务逻辑实现（可供 system 等模块调用）。 |

### 📂 module/system

| 子目录      | 描述 |
|-------------|------|
| `api/`      | 控制器层，处理路由请求（如用户、权限、角色）。 |
| `model/`    | 定义与数据库结构相关的模型，以及业务实体。 |
| `repo/`     | Repository 层，数据访问封装（接口 + 实现）。 |
| `service/`  | 业务服务类，处理非持久化逻辑（如邮件、加密、缓存）。 |
| `usecase/`  | 核心业务逻辑层，协调 `repo` 和 `service` 执行完整流程。 |

---

## 📁 pkg（通用工具包）

| 子目录       | 描述 |
|--------------|------|
| `jwt/`       | JWT Token 封装，支持生成、校验、刷新等操作。 |
| `response/`  | 响应封装，提供统一结构和便捷函数：`Success`、`Error` 等。 |
| `xerror/`    | 自定义错误类型及全局错误码（六位编码规范）。 |

---

## 📊 模块依赖方向建议

```text
api -> usecase -> service/repo -> model
                    ↑
                 common/usecase 可复用调用

```

## 📊 usecase 同层调用建议 

```text
非必要 同层不推荐调用，必要时调用请在提供方 usecase 产出调用接口，使用方借助接口调用，并使用 wire 注入绑定关系 
usecase ->  usecase interface{}
                ↑
        ->  usecase
```