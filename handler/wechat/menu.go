package wechat

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

// 自定义菜单的选项
func MenuQuery(c *gin.Context) {
	mu := Wechat.GetMenu()
	resMenu, err := mu.GetMenu()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resMenu)
	}
}
