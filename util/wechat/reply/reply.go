package reply

import "github.com/silenceper/wechat/message"

// 2018-07-18 12:00
// author: zjt
// 微信公众号被动消息回复方法集合

// 回复空内容
func ReplyNil() *message.Reply {
	return nil
}

// 回复文本信息
func ReplyText(str string) *message.Reply {
	text := message.NewText(str)
	return &message.Reply{message.MsgTypeText, text}
}

// 回复图片消息
func ReplyImage(mediaID string) *message.Reply {
	image :=message.NewImage(mediaID)
	return &message.Reply{message.MsgTypeImage, image}
}

// 回复视频消息
func ReplyViedo(mediaID, videoTitle, description string) *message.Reply {
	video := message.NewVideo(mediaID, videoTitle, description)
	return &message.Reply{message.MsgTypeVideo, video}
}

// 回复音乐消息
// Title:音乐标题
// Description:音乐描述
// MusicURL:音乐链接
// HQMusicUrl：高质量音乐链接，WIFI环境优先使用该链接播放音乐
// ThumbMediaId：缩略图的媒体id，通过素材管理接口上传多媒体文件，得到的id
func ReplyMusic(title, description, musicURL, hQMusicURL, thumbMediaID string) *message.Reply {
	music := message.NewMusic(title, description, musicURL, hQMusicURL, thumbMediaID)
	return &message.Reply{message.MsgTypeMusic, music}
}

// 创建单条图文信息
func CreateImageText(title, description, picURL, url string) *message.Article {
	article := new(message.Article)
	article.Title = title
	article.Description = description
	article.PicURL = picURL
	article.URL = url
	return article
}

// 回复单图文信息
// Title：图文消息标题
// Description：图文消息描述
// PicUrl	：图片链接，支持JPG、PNG格式，较好的效果为大图360200，小图200200
// Url	：点击图文消息跳转链接
func ReplySingleImageText(title, description, picURL, url string) *message.Reply {
	articles := make([]*message.Article, 1)
	article := new(message.Article)
	article.Title = title
	article.Description = description
	article.PicURL = picURL
	article.URL = url
	articles[0] = article
	news := message.NewNews(articles)
	return &message.Reply{message.MsgTypeNews, news}
}

// 回复多图文信息
// Title：图文消息标题
// Description：图文消息描述
// PicUrl	：图片链接，支持JPG、PNG格式，较好的效果为大图360200，小图200200
// Url	：点击图文消息跳转链接
func ReplyImageText(articles []*message.Article) *message.Reply {
	news := message.NewNews(articles)
	return &message.Reply{message.MsgTypeNews, news}
}
