package wechat

import (
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/message"
	"spoon/util/wechat/reply"
	"fmt"
	"spoon/util/turingapi"
)

func Reply(c *gin.Context) {
	// 传入request和responseWriter
	server := Wechat.GetServer(c.Request, c.Writer)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		fmt.Println(msg)
		switch msg.MsgType {
		case message.MsgTypeText: // 文本消息
			return reply.ReplyText(replyText(msg))
		case message.MsgTypeImage: // 图片消息
			return reply.ReplyNil()
		case message.MsgTypeVoice: // 语音消息
			return reply.ReplyNil()
		case message.MsgTypeVideo: // 视频消息
			return reply.ReplyNil()
		case message.MsgTypeShortVideo: // 小视频消息
			return reply.ReplyNil()
		case message.MsgTypeLocation: // 地理位置消息
			return reply.ReplyNil()
		case message.MsgTypeLink: // 链接消息
			return reply.ReplyNil()
		case message.MsgTypeEvent: // 事件推送消息
			switch msg.Event {
			case message.EventSubscribe: // EventSubscribe 订阅
				return reply.ReplyNil()
			case message.EventUnsubscribe: // 取消订阅
				return reply.ReplyNil()
			case message.EventScan: // 用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者
				return reply.ReplyNil()
			case message.EventLocation: // 上报地理位置事件
				return reply.ReplyNil()
			case message.EventClick: // 点击菜单拉取消息时的事件推送
				return reply.ReplyNil()
			case message.EventView: // 点击菜单跳转链接时的事件推送
				return reply.ReplyNil()
			case message.EventScancodePush: // 扫码推事件的事件推送
				return reply.ReplyNil()
			case message.EventScancodeWaitmsg: // 扫码推事件且弹出“消息接收中”提示框的事件推送
				return reply.ReplyNil()
			case message.EventPicSysphoto: // 弹出系统拍照发图的事件推送
				return reply.ReplyNil()
			case message.EventPicPhotoOrAlbum: // 弹出拍照或者相册发图的事件推送
				return reply.ReplyNil()
			case message.EventPicWeixin: // 弹出微信相册发图器的事件推送
				return reply.ReplyNil()
			case message.EventLocationSelect: // 弹出地理位置选择器的事件推送
				return reply.ReplyNil()
			default: // 默认返回空消息
				return reply.ReplyNil()
			}
		default: // 默认返回空消息
			return reply.ReplyNil()
		}
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

// 回复文本的内容
func replyText(msg message.MixMessage) string {
	err, result := turingapi.ChatRobotWithText(msg.Content, nil)
	if err != nil {
		return "机器人回复存在Bug,请谅解"
	} else {
		return turingapi.GetResponseTxt(result)
	}
}
