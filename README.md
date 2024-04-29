### Commit 类型
add: 新功能  

fix: 修复bug

refactor: 代码重构

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
│   └── userGrpc #用户模块proto文件及生成文件
├── logs #日志文件
│   ├── debug
│   ├── error
│   └── info
├── user
│   ├── application #应用层
│   │   ├── router 
│   │   └── service #连接interfaces与domain层
│   ├── domain #领域层
│   │   ├── entity #实体、值对象
│   │   ├── respository #仓库
│   │   └── service #领域层服务代码
│   ├── infrastructure # 基础设施层
│   │   ├── code
│   │   ├── config #配置文件
│   │   ├── interceptor #Grpc拦截器
│   │   ├── persistence#持久层
│   │   │   ├── dao #数据库CRUD操作（Domain层respository的实现）
│   │   │   └── database #数据库连接
│   │   ├── pkg #相关的包
│   └── main.go #user模块入口
```
