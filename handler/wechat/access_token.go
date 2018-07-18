package wechat

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 返回access_token
func AccessToken(c *gin.Context) {
	access_token, err := Wechat.GetAccessToken()
	if err != nil {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusOK, access_token)
	}
}
