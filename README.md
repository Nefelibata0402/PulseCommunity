### Commit 类型
add: 新功能
fix: 修复bug
refactor: 代码重构

├── cmd
│   ├── api
│   │   ├── news
│   │   ├── rpc
│   │   └── user
│   ├── config
│   ├── middlewares
│   ├── model
│   │   ├── news_model
│   │   └── user_model
│   └── router
├── grpc
│   ├── news
│   └── user
├── logs
│   ├── debug
│   ├── error
│   └── info
├── news
│   ├── api
│   ├── └── gen
│   ├── config
│   ├── internal
│   │   ├── dao
│   │   ├── data
│   │   │   └── database
│   │   ├── interceptor
│   │   └── repo
│   ├── pkg
│   └── router
├── user
│   ├── api
│   ├── └── gen
│   ├── config
│   ├── internal
│   │   ├── dao
│   │   ├── data
│   │   │   └── database
│   │   ├── interceptor
│   │   └── repo
│   ├── pkg
│   └── router