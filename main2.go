package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
)

func main() {
	router := gin.Default()

	router.Any("/", hello)
	router.Run(":8001")
}

func hello(c *gin.Context) {

	//配置微信参数
	config := &wechat.Config{
		AppID:          "wx5991d31418d74a4d",
		AppSecret:      "9f3a0acb95255d428846898e258e120c",
		Token:          "dazhen",
		EncodingAESKey: "AxKX1vmaaVgi1GyNxOeFn5OFoYN0l16wmmzE1GKHh5c",
	}
	wc := wechat.NewWechat(config)

	// 传入request和responseWriter
	server := wc.GetServer(c.Request, c.Writer)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {

		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}
