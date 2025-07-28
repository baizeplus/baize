
## 平台简介

Baize 是一个功能齐全的 Web 应用框架，专注于提供基于角色的访问控制（RBAC）功能。它适用于构建企业级管理系统，支持用户权限管理、系统监控、任务调度、数据字典、代码生成等模块，具备完整的前后端分离架构设计。
，毫无保留给个人及企业免费使用。

## 前端代码地址
[baize-react](https://gitee.com/baizeplus/baize-react) https://gitee.com/baizeplus/baize-react  , https://github.com/baizeplus/baize-react
<br>
[baize-vue](https://gitee.com/baizeplus/baize-vue) https://gitee.com/baizeplus/baize-vue  , https://github.com/baizeplus/baize-vue
<br>

```markdown
# Baize：基于角色的访问控制综合Web应用框架

## 项目简介

## 主要功能
- **用户权限管理**：支持用户注册、登录、角色分配、权限控制。
- **角色管理**：支持创建、编辑、删除角色，并为角色分配权限。
- **权限控制**：基于 RBAC 模型，支持细粒度权限控制。
- **系统监控**：包括定时任务管理、在线用户管理、操作日志、登录日志等。
- **数据字典**：支持字典类型和字典数据的管理。
- **代码生成**：支持从数据库表自动生成代码模板(待修改)。
- **文件上传**：支持多种文件存储方式（本地、S3 等）。
- **通知系统**：支持系统消息推送和 SSE（Server-Sent Events）实时通知。

## 技术栈
- **后端**：Go + Gin 框架
- **数据库**：MySQL
- **缓存**：Redis、本地缓存
- **权限控制**：基于 RBAC 模型
- **前端**：Vue和react（配套前端）
- **部署**：Docker 支持
- **日志**：Zap 日志库
```
## 目录结构
```
.
├── app/                            # 主要应用代码
│   ├── baize/                     # 核心框架实用程序和基本类型
│   ├── business/                  # 业务逻辑模块
│   │   ├── monitor/              # 监控功能(jobs, logs, etc.)
│   │   ├── system/               # 核心系统功能 (users, roles.)
│   │   └── tool/                 # 开发工具和生成器
│   ├── constant/                 # 应用程序常量
│   ├── datasource/              # 数据源实现 (MySQL, Redis, S3)
│   ├── docs/                    # API文档 (Swagger)
│   ├── middlewares/             # 中间件（权限、日志、会话）
│   ├── routes/                  # 路由定义
│   ├── setting/                 # 应用程序配置
│   └── utils/                   # 工具类（加密、文件处理、响应封装等）
├── config/                      # 配置文件
├── template/                    # 代码生成模板
└── sql/                        # 数据库脚本

```


## 数据流

该应用程序遵循清晰的架构模式，数据流清晰：

1. HTTP 请求 → 中间件（身份验证/日志记录）→ 控制器
2. 控制器 → 服务层（业务逻辑）
3. 服务层 → 数据访问层 (DAO)
4. DAO → 数据库/缓存/存储
```ascii
Request → [Middleware] → [Controller] → [Service] → [DAO] → [Database]
                                                         ↓
                                                    [Cache/Redis]
                                                         ↓
                                                  [Storage (S3/Local)]
```


关键组件交互：
- 控制器处理 HTTP 请求和响应格式
- 服务实现业务逻辑和事务管理
- DAO 处理数据持久化和检索
- 中间件提供横切关注点（身份验证、日志记录）
- 缓存层提升频繁访问数据的性能
- 存储服务处理文件上传和下载




- 后端采用Gin、Zap、sqly(sqlx升级)。
- 权限认证使用共享Session(单机支持本地缓存,集群需要redis)，支持多终端认证系统。
- 支持加载动态权限菜单，多方式轻松权限控制。
- 高效率开发，使用代码生成器可以一键生成后端代码。(正在完善)
- 特别鸣谢：[ruoyi-vue](https://gitee.com/y_project/RuoYi-Vue?_from=gitee_search )，

## <p>随手 star ⭐是一种美德。 你们的star就是我的动力</p>

## 也期待小伙伴一起加入白泽   

## 内置功能

1. 用户管理：用户是系统操作者，该功能主要完成系统用户配置。
2. 部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
3. 岗位管理：配置系统用户所属担任职务。
4. 菜单管理：配置系统菜单，操作权限，按钮权限标识等。
5. 角色管理：角色菜单权限分配、设置角色按机构进行数据范围权限划分。
6. 字典管理：对系统中经常使用的一些较为固定的数据进行维护。
7. 参数管理：对系统动态配置常用参数。
8. 通知公告：系统通知公告信息发布维护。
9. 登录日志：系统登录日志记录查询包含登录异常。
10. 在线用户：当前系统中活跃用户状态监控。
11. 服务监控：监视当前系统CPU、内存、磁盘等相关信息。
12. 操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
13. 系统接口：根据业务代码注释自动生成相关的api接口文档。
14. 代码生成：前后端代码的生成（Go、sql）支持CRUD下载 。
15. 定时任务：在线（添加、修改、删除)任务调度。
## 版本规则
v5.6.7<br>
1位为主版本号（5）：当功能模块有较大的变动，比如增加多个模块或者整体架构发生变化,数据库结构发生变化。此版本号由项目决定是否修改。
<br>
2为次版本号（6）：当功能有一定的增加或变化，比如修改了API接口。此版本号由项目决定是否修改。
<br>
3为阶段版本号(7)：一般是 Bug 修复或是一些小的变动，要经常发布修订版，时间间隔不限。
<br>
主版本号升级请参考更新说明更新修改或添加相应的数据表。
次版本号升级请参考更新说明查看API接口修改情况。
阶段版本号不会影响数据库与api接口，除修复重大bug不更新说明文档



## 在线体验

- admin/admin123

react演示地址：http://react.ibaize.vip
<br>
vue3演示地址：http://vue.ibaize.vip
<br>
文档地址：https://doc.ibaize.vip
<br>

## 演示图

<table>
    <tr>
        <td><img src="https://gitee.com/smell2/BaiZe/raw/imgs/202110241805797.jpg"/></td>
        <td><img src="https://gitee.com/smell2/BaiZe/raw/imgs/202110241806256.jpg"/></td>
    </tr>
    <tr>
        <td><img src="https://gitee.com/smell2/BaiZe/raw/imgs/202110242322137.png"/></td>
        <td><img src="https://gitee.com/smell2/BaiZe/raw/imgs/202110242323820.png"/></td>
    </tr>  
    <tr>
        <td><img src="https://gitee.com/smell2/BaiZe/raw/imgs/202112082243214.png"/></td>
        <td><img src="https://gitee.com/smell2/BaiZe/raw/imgs/202112082242154.png"/></td>
    </tr>

</table>



## 白泽管理系统交流群


QQ群： [![加入QQ群](https://img.shields.io/badge/83064682-blue.svg)](https://qm.qq.com/cgi-bin/qm/qr?k=rAIw_VQ_blbSQu0J6fApnm5RbAc2CHbp&jump_from=webapi) 点击按钮入群。
##欢迎加入
