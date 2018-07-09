# Spoon

### 概述
<pre>
    使用第三方框架
        Gin     --      Restful API框架
        GORM    --      数据库ORM框架
</pre>

### 更新进程
<pre>
    2018/07/05 23:00    添加云片短信发送平台
    2018/07/06 14:00    添加sendcloud邮件发送功能
    2018/07/07 17:00    添加二维码功能和captcha的api功能
    2018/07/09 10:00    添加HTTPS访问功能(证书Site: api.jiagongwu.com)
    2018/07/09 11:00    用Makefile管理API.  Makefile学习:https://www.cnblogs.com/wang_yb/p/3990952.html
    2018/07/09 11:30    添加API命令添加版本功能
</pre>

### 目录结构
<pre>
    ├── admin.sh                     # 进程的start|stop|status|restart控制文件
    ├── conf                         # 配置文件统一存放目录
    │   ├── config.yaml              # 配置文件
    │   ├── server.crt               # TLS配置文件
    │   └── server.key
    ├── config                       # 专门用来处理配置和配置文件的Go package
    │   └── config.go
    ├── db.sql                       # 在部署新环境时，可以登录MySQL客户端，执行source db.sql创建数据库和表
    ├── docs                         # swagger文档，执行 swag init 生成的
    │   ├── docs.go
    │   └── swagger
    │       ├── swagger.json
    │       └── swagger.yaml
    ├── handler                      # 类似MVC架构中的C，用来读取输入，并将处理流程转发给实际的处理函数，最后返回结果
    │   ├── handler.go
    │   ├── sd                       # 健康检查handler
    │   │   └── check.go
    │   └── user                     # 核心：用户业务逻辑handler
    │       ├── create.go            # 新增用户
    │       ├── delete.go            # 删除用户
    │       ├── get.go               # 获取指定的用户信息
    │       ├── list.go              # 查询用户列表
    │       ├── login.go             # 用户登录
    │       ├── update.go            # 更新用户
    │       └── user.go              # 返回json的结构体
</pre>
