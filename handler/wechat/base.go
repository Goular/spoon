package wechat

import (
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
)

// 全局的wechat的实例
var Wechat *wechat.Wechat

func init() {
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
	Wechat = wechat.NewWechat(config)
}
