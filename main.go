package main

import (
	"errors"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	v "spoon/pkg/version"
	"spoon/config"
	"spoon/model"
	"spoon/router"
	"spoon/router/middleware"
	"fmt"
	"os"
	"encoding/json"
	_ "spoon/handler/wechat" // 添加wechat公众号的初始方法
)

var (
	cfg     = pflag.StringP("config", "c", "", "apiserver config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	pflag.Parse()

	// 添加版本信息的显示
	if *version {
		v := v.Get()
		marshalled, err := json.MarshalIndent(&v, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(marshalled))
		return
	}

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// init db
	model.DB.Init()
	defer model.DB.Close()

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	// Routes.
	router.Load(
		// Cores.
		g,

		// Middlwares.
		middleware.Logging(),
		middleware.RequestId(),
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	// 开启HTTPS服务器功能
	//cert := viper.GetString("tls.cert")
	//key := viper.GetString("tls.key")
	//if cert != "" && key != "" {
	//	go func() {
	//		log.Infof("Start to listening the incoming requests on https address: %s", viper.GetString("tls.addr"))
	//		log.Info(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
	//	}()
	//}

	// 开启普通HTTP服务器的功能
	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	// log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error()) // 使用普通的Golang自带的服务器进行HTTP服务启动
	g.Run(viper.GetString("addr")) // 使用Gin服务器进行HTTP服务启动
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
