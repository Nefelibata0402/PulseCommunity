### Commit 类型
add: 新功能  

fix: 修复bug

refactor: 代码重构

air热重载

```
├── cmd
│   ├── config
│   ├── interfaces #用户接口层
│   │   ├── news
│   │   ├── rpc
│   │   └── user
│   ├── middlewares
│   ├── model
│   │   ├── newsModel
│   │   └── userModel
│   └── router
├── idl #接口描述语言
│   ├── gen #临时proto文件
│   ├── news
│   └── user
├── logs
│   ├── debug
│   ├── error
│   └── info
├── news
│   ├── application #应用层
│   │   ├── code
│   │   ├── pkg
│   │   └──service
│   ├── config
│   ├── domain #领域层
│   │   └── respository #仓库
│   ├── infrastructure # 基础设施层
│   │   ├── interceptor
│   │   ├── persistence    # 数据库相关
│   │   │   ├── dal
│   │   │   ├── database
│   │   │   └── newsData
│   └── router
├── user
│   ├── application #应用层
│   │   ├── code
│   │   ├── pkg
│   │   └──service
│   ├── config
│   ├── domain #领域层
│   │   └── respository #仓库
│   ├── infrastructure # 基础设施层
│   │   ├── interceptor
│   │   ├── persistence    # 数据库相关
│   │   │   ├── dal
│   │   │   ├── database
│   │   │   └── userData
│   ├── pkg
│   └── router
```
