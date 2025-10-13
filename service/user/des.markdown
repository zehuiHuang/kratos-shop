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
```

```