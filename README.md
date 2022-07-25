## lebron


```sh
# 使用goctl生成代码
go install github.com/zeromicro/go-zero/tools/goctl@latest
# goctl 生成rpc 代码,后面的rpc是文件夹名称
goctl rpc new rpc
# goctl 生成 http 代码
goctl api new admin

# goctl 生成修改过的代码
goctl api go -api api.api -dir .
```

### 本项目包含很多小项目。apps文件夹中，每个文件夹都是一个小项目
```
.
├── README.md
├── apps
│   ├── app
│   ├── cart
│   ├── order
│   ├── pay
│   ├── product
│   ├── recommend
│   ├── reply
│   └── user
└── go.mod
```

### 每个小项目中又分为不同的模块
- api - 对外的BFF服务，接受来自客户端的请求，暴露HTTP接口
- rpc - 对内的微服务，仅接受来自内部其他微服务或者BFF的请求，暴露gRPC接口
- rmq - 负责进行流式任务处理，上游一般依赖消息队列，比如kafka等
- admin - 也是对内的服务，区别于rpc，更多的是面向运营侧的且数据权限较高，通过隔离可带来更好的代码级别的安全，直接提供HTTP接口

