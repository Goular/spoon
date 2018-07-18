package wechat

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"spoon/util/wechat/menu"
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

const menu_str = `
{
    "button": [
        {
            "name": "扫码", 
            "sub_button": [
                {
                    "type": "scancode_waitmsg", 
                    "name": "扫码带提示", 
                    "key": "rselfmenu_0_0", 
                    "sub_button": [ ]
                }, 
                {
                    "type": "scancode_push", 
                    "name": "扫码推事件", 
                    "key": "rselfmenu_0_1", 
                    "sub_button": [ ]
                }
            ]
        }, 
        {
            "name": "发图", 
            "sub_button": [
                {
                    "type": "pic_sysphoto", 
                    "name": "系统拍照发图", 
                    "key": "rselfmenu_1_0", 
                   "sub_button": [ ]
                 }, 
                {
                    "type": "pic_photo_or_album", 
                    "name": "拍照或者相册发图", 
                    "key": "rselfmenu_1_1", 
                    "sub_button": [ ]
                }, 
                {
                    "type": "pic_weixin", 
                    "name": "微信相册发图", 
                    "key": "rselfmenu_1_2", 
                    "sub_button": [ ]
                }
            ]
        }, 
        {
            "name": "发送位置", 
            "type": "location_select", 
            "key": "rselfmenu_2_0"
        },
        {
           "type": "media_id", 
           "name": "图片", 
           "media_id": "MEDIA_ID1"
        }, 
        {
           "type": "view_limited", 
           "name": "图文消息", 
           "media_id": "MEDIA_ID2"
        }
    ]
}
`

// 创建自定义菜单
func MenuCreate(c *gin.Context) {
	menu.CreateMenu(Wechat, menu_str)
}
