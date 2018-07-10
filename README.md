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
    2018/07/09 15:00    API 高可用方案 (HA) 使用keepalived为多个Nginx添加高可用的结构
    2018/07/10 10:00    Go Test 单元测试,压力/性能测试,性能分析//Test、Benchmark、Example 开头的测试函数
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


# Golang测试
<pre>
    性能测试: https://github.com/hyper0x/go_command_tutorial/blob/master/0.12.md

    go test 执行测试用例时，是以 go 包为单位进行测试的。执行时需要指定包名，比如：go test 包名，如果没有指定包名，默认会选择执行命令时所在的包。go test 在执行时会遍历以 _test.go 结尾的源码文件，执行其中以 Test、Benchmark、Example 开头的测试函数。其中源码文件需要满足以下规范：

    文件名必须是 _test.go 结尾，跟源文件在同一个包。
    测试用例函数必须以 Test、Benchmark、Example 开头
    执行测试用例时的顺序，会按照源码中的顺序依次执行
    单元测试函数 TestXxx() 的参数是 testing.T，可以使用该类型来记录错误或测试状态
    性能测试函数 BenchmarkXxx() 的参数是 testing.B，函数内以 b.N 作为循环次数，其中 N 会动态变化
    示例函数 ExampleXxx() 没有参数，执行完会将输出与注释 // Output: 进行对比
    测试函数原型：func TestXxx(t *testing.T)，Xxx 部分为任意字母数字组合，首字母大写，例如： TestgenShortId 是错误的函数名，TestGenShortId 是正确的函数名

    通过调用 testing.T 的 Error、Errorf、FailNow、Fatal、FatalIf 方法来说明测试不通过，通过调用 Log、Logf 方法来记录测试信息：
        t.Log t.Logf     # 正常信息
        t.Error t.Errorf # 测试失败信息
        t.Fatal t.Fatalf # 致命错误，测试程序退出的信息
        t.Fail     # 当前测试标记为失败
        t.Failed   # 查看失败标记
        t.FailNow  # 标记失败，并终止当前测试函数的执行，需要注意的是，我们只能在运行测试函数的 Goroutine 中调用 t.FailNow 方法，而不能在我们在测试代码创建出的 Goroutine 中调用它
        t.Skip     # 调用 t.Skip 方法相当于先后对 t.Log 和 t.SkipNow 方法进行调用，而调用 t.Skipf 方法则相当于先后对 t.Logf 和 t.SkipNow 方法进行调用。方法 t.Skipped 的结果值会告知我们当前的测试是否已被忽略
        t.Parallel # 标记为可并行运算

    测试文件与源文件最好同一个包


    压力测试
        在 util 目录下执行命令 go test -test.bench=".*"：

        $ go test -test.bench=".*"
        goos: linux
        goarch: amd64
        pkg: apiserver/util
        BenchmarkGenShortId-2                	  500000	      2291 ns/op
        BenchmarkGenShortIdTimeConsuming-2   	  500000	      2333 ns/op
        PASS
        ok  	apiserver/util	2.373s
        复制代码
        上面的结果显示，我们没有执行任何 TestXXX 的单元测试函数，只执行了压力测试函数
        第一条显示了 BenchmarkGenShortId 执行了 500000 次，每次的执行平均时间是 2291 纳秒
        第二条显示了 BenchmarkGenShortIdTimeConsuming 执行了 500000，每次的平均执行时间是 2333 纳秒
        最后一条显示总执行时间
        BenchmarkGenShortIdTimeConsuming 比 BenchmarkGenShortId 多了两个调用 b.StopTimer() 和 b.StartTimer()。

        b.StopTimer()：调用该函数停止压力测试的时间计数
        b.StartTimer()：重新开始时间
        在 b.StopTimer() 和 b.StartTimer() 之间可以做一些准备工作，这样这些时间不影响我们测试函数本身的性能。

    查看性能并生成函数调用图
            执行命令：
            $ go test -bench=".*" -cpuprofile=cpu.profile ./util
            复制代码
            上述命令会在当前目录下生成 cpu.profile 和 util.test 文件。

            执行 go tool pprof util.test cpu.profile 查看性能（进入交互界面后执行 top 指令）

    小总结
        在实际的开发中，要养成编写单元测试代码的好习惯，在项目上线前，最好对一些业务逻辑比较复杂的函数做一些性能测试，提前发现性能问题。

        至于怎么去分析性能，比如查找耗时最久的函数等，笔者链接了郝林大神专业的分析方法（go tool pprof），更深的分析技巧需要读者在实际开发中自己去探索。

</pre>
