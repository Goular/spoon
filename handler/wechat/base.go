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
		AppID:          "wxa90a30f78693c0c2",
		AppSecret:      "db543fccecfbd20922643e0f826a8ab0",
		Token:          "xingyunshulian",
		EncodingAESKey: "tByjObyhPtNneGdNn4MNVLbV2m8kNzsxjPJWQ70OFip",
		Cache:          redis,
	}
	Wechat = wechat.NewWechat(config)
}
