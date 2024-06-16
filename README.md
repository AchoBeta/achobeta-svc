# achobeta-svc
achobeta 团队 go 服务核心代码, 使用大仓的管理模式, 简化项目依赖
> 小仓也是很好的管理模式, 但暂时没有拆出去的规划

# 架构
api层使用http协议通信, api 与 service 之间使用grpc通信, 具体接口定义在proto中

# 本地启动方式
#### 前提
api\authz 目录下的config文件记得写, 已经给了模板, 参与团队项目开发时请联系管理员
#### 编译及启动
项目依赖proto, 所以必须先执行编译后才能启动, 太长不看可以只看全局方式
```bash
# (全局) 在 achobeta-svc 目录下执行以下命令
make install  # 用于下载依赖, 一般只需要下载一次
# 本地编译必须先编译proto, 可以直接
make build
# 单独编译proto时, 也可以执行
make proto
# 启动提供两种方式, 推荐docker
# 1. docker, make build 之后使用docker compose
docker-compose up
# 后台启动
docker-compose up -d 
# 也可以加上--build, 构建启动即
docker-compose up -d --build
# 2. 进程方式启动(一般用于调试), 在make build 之后
# 以启动achobeta-svc-api为例 
make run target=api # 如果是authz则target=authz, 详情见Makefile

# (子项目编译, 以achobeta-svc-authz为例)
# authz 项目依赖grpc, 所以需要先编译proto
# 1. 可以先在proto的目录下执行
make install build # install 只需要执行一次
# 2. 然后在authz 目录下执行
make build
```

# 检查
修改代码后需要检查代码是否规范, 项目提供了全局检查和子项目检查
```bash
# 1. 全局检查, 在项目根目录下执行
make lint
# 2. 子项目检查, 子项目下执行
make lint

# tips: 检查插件使用的是 buf(proto)/golang-lint(go)
# 可能需要先下载依赖, 既
make install
```



# 其他
如有其他问题请联系 Achobeta