package main

import (
	"github.com/gin-gonic/gin"
	_ "spoon/handler/wechat"
	"spoon/handler/wechat"
)

// 微信公众号测试
func main() {
	router := gin.Default()
	router.Any("/wechat/reply", wechat.Reply)
	router.GET("/access_token",wechat.AccessToken)
	router.Run(":8001")
}