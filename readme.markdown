参考------------
https://learnku.com/articles/65095

https://github.com/aliliin/kratos-shop.git


常用命令:
```shell
##创建项目
kratos new helloworld

##添加 Proto 文件
kratos proto add api/goods/v1/goods.proto

##生成 Proto 代码  
make api

##  通过 proto 文件生产客户端代码
kratos proto client api/goods/v1/goods.proto

## 通过 proto 文件，可以直接生成对应的 Service 实现代码
kratos proto server api/goods/v1/goods.proto -t internal/service

#根据proto文件创建基础类型结构体
make config

#项目启动
kratos run

```


# 常见问题
1、使用idea中的debug运行时,配置中需要加上:
-conf ./configs
![img.png](img.png)