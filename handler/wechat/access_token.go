package wechat

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

// 返回access_token
func AccessToken(c *gin.Context) {
	access_token, err := Wechat.GetAccessToken()
	fmt.Println("112233")
	fmt.Println(access_token)
	fmt.Println("9966")
	if err != nil {
		c.JSON(http.StatusOK, access_token)
	} else {
		c.JSON(http.StatusOK, nil)
	}
}
