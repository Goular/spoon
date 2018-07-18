package wechat

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/silenceper/wechat/menu"
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
	mu := Wechat.GetMenu()
	buttons := make([]*menu.Button, 1)
	btn := new(menu.Button)
	//创建click类型菜单
	btn.SetClickButton("name", "key123")
	buttons[0] = btn
	//设置btn为二级菜单
	btn2 := new(menu.Button)
	btn2.SetSubButton("subButton", buttons)
	buttons2 := make([]*menu.Button, 1)
	buttons2[0] = btn2
	//发送请求
	err := mu.SetMenu(buttons2)
	if err != nil {
		fmt.Printf("err= %v", err)
		return
	}
}
