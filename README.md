### Commit 类型
add: 新功能  

fix: 修复bug

refactor: 代码重构

```
pulseCommunity
├── cmd #客户端表示层：处理用户请求并将请求转发到应用层
│   ├── config #配置文件
│   ├── interfaces #用户接口层
│   │   ├── article #以article为例，还有其他例如：user、ranking、search等等
│   │   │   ├──article.go #路由处理函数
│   │   │   ├──router.go #路由注册
│   │   │   └──rpc.go #gRPC客户端
│   │   ├── main.go #初始化所有router.go
│   │   └── router.go #初始化所有个rpc.go
│   ├── middlewares #prometheus监控、tokenVerify、限流等等
│   └── model 
├── common #都有使用的 例如：服务发现
├── idl #接口描述语言
├── logs #日志文件
├── article 文章模块 (以article为例，还有ranking、search、user等等)
│   ├── application #应用层：处理应用程序的业务流程。协调领域模型和用户界面之间的交互，但不包含领域逻辑。
│   │   └── service #连接interfaces与domain层
│   ├── domain #领域层：包含核心的业务逻辑和领域模型。领域层封装了业务规则和领域逻辑
│   │   ├── entity #实体、值对象
│   │   ├── event #领域事件
│   │   ├── respository #仓库
│   │   └── service #领域层服务代码
│   ├── infrastructure # 基础设施层：实现与外部系统的交互和技术细节，包括数据持久化、消息队列、第三方服务等
│   │   ├── config #配置文件
│   │   ├── persistence#持久层
│   │   │   ├── convertor #数据库和实体之间的转换
│   │   │   ├── dao #数据库CRUD操作（Domain层respository的实现）
│   │   │   └── database #数据库连接
│   │   ├── mq #消息队列初始化
│   │   └── pkg #自己编写的包：Kafka消费者批量消费等
│   │── router.go #服务端Grpc，Etcd服务发现等等初始化
│   └── main.go #user模块入口
```
