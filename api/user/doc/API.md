# bitbucket.org/pwq/tata/api/user

##### 项目简介
> 1. 提供监管平台的数据加解密的实现，实现明文请求加密提交
> 2. 提供监管平台的数据直接推送接口，供业务使用
> 3. 提供监管平台的提交队列接口的实现，供业务接入使用

##### 编译环境
> 1. 请使用golang v1.18.x以上版本编译执行。

##### 外部依赖包
> 1. 无

##### 编译执行
> 1. 启动接入示例（ext）
	stack.PushLoop(context.TODO(), 10) // 循环队列

	ext.ListenAndServe(log) // 接口监听

> 2. 配置参考
    -http.addr :8000
    -db.mysql {"addr":"test","dsn":"root:debug@tcp(localhost:3306)/education?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8,utf8mb4","readDSN":["root:debug@tcp(localhost:3306)/education?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8,utf8mb4"],"archive":20,"idle":10,"idleTimeout":"4h","queryTimeout":"15s","execTimeout":"5s","tranTimeout":"5s"}

##### 测试
> 1. 执行当前目录下所有测试文件，测试所有功能

##### 特别说明
> 1. 无


### sdk serviceName, operationName, tag 规范

注: serviceName 及之前的 family 字段, operationName 及之前的 title 字段

serviceName 使用 APP_ID 可以通过 caster 上 APP_ID 环境变量获取

#### 全局 Tag

| 名称     | 类型   | 备注                                                 |
|----------|--------|------------------------------------------------------|
| hostname | string | 主机名                                               |
| ip       | string | caster上使用 POD_IP 环境变量，其他环境取第一个外网IP |
| zone     | string | zone caster 使用 ZONE 环境变量 e.g. sh               |
| region   | string | region caster 使用 REGION 环境变量 e.g. region       |

#### HTTP

HTTP server && client 共同 tag

| 名称             | 类型   | 备注                                       |
|------------------|--------|--------------------------------------------|
| http.method      | string | GET、POST ...                              |
| http.url         | string | http 完整 URL，包含 query                  |
| http.status_code | int    | http 状态码                                |

HTTP server 

operationName 设置:

- 非 restful API 的应用使用 URL path 部分 例如 URL http://api.tita.co/x/internal/user/info?mid=123 operationName 为 /x/internal/user/info
- restful API 的使用路由定义，使用 {} 代替可变部分, 例如 URL http://api.tita.co/x/internal/user/info/123 其中 123 为 mid，则 operationName 设置为 /x/internal/user/info/{mid}

| 名称      | 类型   | 备注                                       |
|-----------|--------|--------------------------------------------|
| span.kind | string | 固定值 server                              |
| component | string | 组件名称 e.g. library/net/http/baldemaster |

HTTP client

operationName 设置:

- 请求内部非 restful 的应用可以直接设置为 URL 的 path 部分
- 请求的三方的服务或者 restful API operationName 可以直接设置为 HTTP:{Method} e.g. HTTP:GET

| 名称         | 类型   | 备注                                                                                 |
|--------------|--------|--------------------------------------------------------------------------------------|
| span.kind    | string | 固定值 client                                                                        |
| component    | string | 组件名称 e.g. library/net/http 或者 net/http                                         |
| peer.service | string | 请求的服务APP_ID，例如请求 account-service 则应该设置为 main.account.account-service |
| \_peer.sign  | string | URL 的 path 部分不包含 query                                                         |

注: peer.service 不知道可以不设置，_peer.sign 用于自定探测 peer.service

#### gRPC

gRPC server && client 共同 tag

gRPC server

operationName 设置:

- 使用 FullMethod https://github.com/grpc/grpc-go/blob/master/interceptor.go#L47:2	


| 名称      | 类型   | 备注          |
|-----------|--------|---------------|
| span.kind | string | 固定值 server |
| component | string | 固定值 gRPC   |

gRPC client

operationName 设置:

- 使用 FullMethod https://github.com/grpc/grpc-go/blob/master/interceptor.go#L47:2	

| 名称         | 类型   | 备注                                                                                 |
|--------------|--------|--------------------------------------------------------------------------------------|
| span.kind    | string | 固定值 client                                                                        |
| component    | string | 固定值 gRPC                                                                          |
| peer.service | string | 请求的服务APP_ID，例如请求 account-service 则应该设置为 main.account.account-service |
| \_peer.sign  | string | gRPC FullMethod                                                                      |

#### goRPC

TODO

#### Memcache

TODO

#### Redis

TODO

#### MySQL

先进行声明和初始化：
mysql.Session = mysql.New(mysql.Parse(env.Value.DB.Mysql))
fmt.Println(mysql.Session.Ping(context.TODO()))

直接调用实现实例即可完成开发：
result, err := mysql.Session.GetInfoByID(context.TODO(), id)

#### Databus

TODO

#### HBase

TODO

#### ElasticSearch

TODO
