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
    2018/07/09 11:30    添加API命令添加版本功能 -- 使用的是git的版本控制的信息 使用的是 go build -v -ldflags ${ldflags} .
    2018/07/09 12:00    给API增加启动脚本
    2018/07/09 14:38    基于Nginx的API部署方案
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

# Nginx
<pre>
    Nginx 反向代理功能
        server {
                listen      80;
                server_name  apiserver.com;
                client_max_body_size 1024M;

                location / {
                    proxy_set_header Host $http_host;
                    proxy_set_header X-Forwarded-Host $http_host;
                    proxy_set_header X-Real-IP $remote_addr;
                    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
                    proxy_pass  http://127.0.0.1:8080/;
                    client_max_body_size 100m;
                }
        }

    Nginx 负载均衡功能
        Nginx 另一个常用的功能是负载均衡，所谓的负载均衡就是指当 Nginx 收到一个 HTTP 请求后，会根据负载策略将请求转发到不同的后端服务器上。比如，apiserver 部署在两台服务器 A 和 B 上，当请求到达 Nginx 后，Nginx 会根据 A 和 B 服务器上的负载情况，将请求转发到负载较小的那台服务器上。这里要求 apiserver 是无状态的服务.

    Nginx 常用命令
        nginx -s stop       快速关闭 Nginx，可能不保存相关信息，并迅速终止 Web 服务
        nginx -s quit       平稳关闭 Nginx，保存相关信息，有安排的结束 Web 服务
        nginx -s reload     因改变了 Nginx 相关配置，需要重新加载配置而重载
        nginx -s reopen     重新打开日志文件
        nginx -c filename   为 Nginx 指定一个配置文件，来代替默认的
        nginx -t            不运行，而仅仅测试配置文件。Nginx 将检查配置文件的语法的正确性，并尝试打开配置文件中所引用到的文件
        nginx -v            显示 Nginx 的版本
        nginx -V            显示 Nginx 的版本、编译器版本和配置参数

    配置 Nginx 作为负载均衡
        负载均衡的演示需要多个后端服务，为此我们在同一个服务器上启动多个 apiserver，配置不同的端口（8080、8082），并采用 Nginx 默认的轮询转发策略（轮询：每个请求按时间顺序逐一分配到不同的后端服务器）。

</pre>
