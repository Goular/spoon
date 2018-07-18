package wechat

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"spoon/util/wechat/menu"
	"spoon/handler"
	"github.com/pkg/errors"
)

// 获取自定义菜单
func MenuGet(c *gin.Context) {
	mu := Wechat.GetMenu()
	resMenu, err := mu.GetMenu()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resMenu)
	}
}

// 删除自定义菜单
func MenuDelete(c *gin.Context) {
	mu := Wechat.GetMenu()
	err := mu.DeleteMenu()
	if err != nil {
		fmt.Println(err)
	}
}

// 创建自定义菜单
func MenuCreate(c *gin.Context) {
	var menu_str = `{"button":[{"type":"click","name":"今日歌曲22","key":"V1001_TODAY_MUSIC"},{"type":"click","name":"歌手简介","key":"V1001_TODAY_SINGER"},{"name":"菜单","sub_button":[{"type":"view","name":"搜索","url":"http://www.soso.com/"},{"type":"view","name":"视频","url":"http://v.qq.com/"},{"type":"click","name":"赞一下我们","key":"V1001_GOOD"}]}]}`
	err := menu.CreateMenu(Wechat, menu_str)
	if err != nil {
		handler.SendResponse(c, err, nil)
	} else {
		handler.SendResponse(c, errors.New("菜单成功创建"), nil)
	}
}
