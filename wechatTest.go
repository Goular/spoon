package main

import (
	"github.com/gin-gonic/gin"
	_ "spoon/handler/wechat"
	"spoon/handler/wechat"
)

// 微信公众号测试
func main() {
	router := gin.Default()
	wechats := router.Group("/wechat")
	{
		wechats.Any("/reply", wechat.Reply)
		wechats.GET("/access_token", wechat.AccessToken)
		wechats.GET("/menu_get", wechat.MenuGet)
		wechats.GET("/menu_delete", wechat.MenuDelete)
		wechats.GET("/menu_create", wechat.MenuCreate)
	}
	router.Run(":8001")
}
