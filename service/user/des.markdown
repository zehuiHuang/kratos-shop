## kratos 相关
```
kratos new user

cd user

kratos proto add api/user/v1/user.proto

kratos proto server api/user/v1/user.proto -t internal/service
```

## 项目结构
1、api---- api的接口定义
2、configs -----配置文件
3、cmd-----------启动main
4、internal------conf-----结构体
5、internal------data-----数据客户端初始化(比如redis、mysql等)
6、internal------biz-----业务数据结构体(类似于DDD的核心业务对象)
7、internal------service-----


# 环境依赖
```shell
##consul
docker run -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600/udp consul:1.15.4 consul agent -dev -client=0.0.0.0

## Jaege
docker run --rm --name jaeger -p14268:14268 -p16686:16686 jaegertracing/all-in-one
```


# API 调用验证
```shell

```