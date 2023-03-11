# 目录结构
## configs:
配置文件  
## docs:
文档集合  
## global:
全局变量  
## internal:
内部模块  

    dao:数据访问层（Database Access Object),所有与数据相关的操作都会在dao层进行，例如mysql，Elasticsearch等。
    middleware:HTTP中间层。
    model:模型层,用于存放model对象。
    routers:路由层相关的逻辑。
    service:项目核心业务逻辑。
## pkg:
项目相关的模块包。
## storage:
项目生成的临时文件。
## scripts:
各类构建、安装、分析等操作的脚本。
## third_party:
第三方的资源工具,如Swagger UI。


# RESTful API

## GET:  
读取和检索动作。
## POST:  
新增和新建动作
## PUT:
更新动作，用于更新一个完整的资源，要求为幂等。
## PATCH:  
更新动作，用于更新某一个资源的一个组成部分，也就是说，当只需要更新改资源的某一项时，应该使用PATCH而不是PUT，可以不幂等。
## DELETE:  
删除动作