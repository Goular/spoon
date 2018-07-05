package captcha

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 通过httpclient发送到第三方服务商请求手机验证码
func Gain(c *gin.Context) {
	c.JSON(http.StatusOK,"lalalal")
}
