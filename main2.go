package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
	"github.com/silenceper/wechat/cache"
)

func main() {
	router := gin.Default()

	router.Any("/", hello)
	router.Run(":8001")
}

func hello(c *gin.Context) {
	redis := cache.NewRedis(&cache.RedisOpts{
		Host:     "127.0.0.1:6379",
		Password: "3071611103",
		Database: 15,
	})
	//配置微信参数
	config := &wechat.Config{
		AppID:          "wx5991d31418d74a4d",
		AppSecret:      "9f3a0acb95255d428846898e258e120c",
		Token:          "dazhen",
		EncodingAESKey: "AxKX1vmaaVgi1GyNxOeFn5OFoYN0l16wmmzE1GKHh5c",
		Cache:          redis,
	}
	wc := wechat.NewWechat(config)

	// 传入request和responseWriter
	server := wc.GetServer(c.Request, c.Writer)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		switch msg.MsgType {
		//文本消息
		case message.MsgTypeText:
			//do something

			//图片消息
		case message.MsgTypeImage:
			//do something

			//语音消息
		case message.MsgTypeVoice:
			//do something

			//视频消息
		case message.MsgTypeVideo:
			//do something

			//小视频消息
		case message.MsgTypeShortVideo:
			//do something

			//地理位置消息
		case message.MsgTypeLocation:
			//do something

			//链接消息
		case message.MsgTypeLink:
			//do something

			//事件推送消息
		case message.MsgTypeEvent:
			switch msg.Event {
			//EventSubscribe 订阅
			case message.EventSubscribe:
				//do something

				//取消订阅
			case message.EventUnsubscribe:
				//do something

				//用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者
			case message.EventScan:
				//do something

				// 上报地理位置事件
			case message.EventLocation:
				//do something

				// 点击菜单拉取消息时的事件推送
			case message.EventClick:
				//do something

				// 点击菜单跳转链接时的事件推送
			case message.EventView:
				//do something

				// 扫码推事件的事件推送
			case message.EventScancodePush:
				//do something

				// 扫码推事件且弹出“消息接收中”提示框的事件推送
			case message.EventScancodeWaitmsg:
				//do something

				// 弹出系统拍照发图的事件推送
			case message.EventPicSysphoto:
				//do something

				// 弹出拍照或者相册发图的事件推送
			case message.EventPicPhotoOrAlbum:
				//do something

				// 弹出微信相册发图器的事件推送
			case message.EventPicWeixin:
				//do something

				// 弹出地理位置选择器的事件推送
			case message.EventLocationSelect:
				//do something

			}
		}
		// todo: 需要暂时返回nil
		return nil
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
