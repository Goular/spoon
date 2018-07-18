package main

import (
	"github.com/gin-gonic/gin"
	_ "spoon/handler/wechat"
	"spoon/handler/wechat"
	"github.com/spf13/viper"
)

// 微信公众号测试
func main() {
	router := gin.Default()
	router.Any("/wechat/reply", wechat.Reply)
	router.Run(viper.GetString("addr"))
}