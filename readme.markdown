参考------------
https://learnku.com/articles/65095

https://github.com/aliliin/kratos-shop.git


常用命令:
```shell
##创建项目
kratos new helloworld

##添加 Proto 文件
kratos proto add api/helloworld/v1/demo.proto

##生成 Proto 代码  
make api
##  通过 proto 文件生产客户端代码
kratos proto client api/helloworld/v1/demo.proto

## 通过 proto 文件，可以直接生成对应的 Service 实现代码
kratos proto server api/helloworld/v1/demo.proto -t internal/service


make config


```