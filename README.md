# micro-todolist
> 该项目是一个基于Go-Micro框架实现的简单备忘录微服务应用

## 主要功能
- 用户注册登录(jwt-go 鉴权)
- 新增/删除/修改/查询 备忘录

## 项目架构
### micro-todolist
```go
micro-todolist/
├── app                   // 各个微服务
│   ├── gateway           // 网关
│   ├── task              // 任务模块微服务
│   └── user              // 用户模块微服务
├── config                // 配置文件
├── consts                // 定义的常量
├── idl                   // protoc文件
│   └── pb                // 放置生成的pb文件
├── pkg                   // 各种包
│   ├── ctl               // 用户操作
│   ├── e                 // 统一错误状态码
│   └── util              // 各种工具、JWT等等..
└── types                 // 定义各种结构体
```
### gateway 组件
```go
gateway/
├── cmd                   // 启动入口
├── http                  // HTTP请求头
├── middleware            // 中间件
├── router                // http 路由模块
├── rpc                   // rpc 调用
└── wrappers              // 熔断
```
### 服务组件
```go
task/
├── cmd                   // 启动入口
├── service               // 业务服务
├── repository            // 持久层
│    ├── db               // 视图层
│    │    ├── dao         // 对数据库进行操作
│    │    ├── redis       // 缓存操作
│    │    └── model       // 定义数据库的模型
│    └── mq               // 放置 mq
├── script                // 监听 mq 的脚本
```
## 项目主要依赖

**Golang V1.18**

- Gin
- Gorm
- go-micro
- protoc
- grpc
- ini
- hystrix
- jwt-go
- crypto 
- go-redis

## 运行命令

1. 启动环境
```shell
make env-up
```
2. 安装依赖
```shell
go mod tidy
```
3. 运行服务
```shell
make run
```

## 导入 Postman 接口文档
> 直接导入即可