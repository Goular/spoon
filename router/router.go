package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"spoon/handler/user"
	"spoon/router/middleware"
	"spoon/handler/sd"
	"spoon/handler/captcha"
	"spoon/handler/email"
	"spoon/handler/qrcode"
	"github.com/gin-contrib/pprof"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// pprof router
	pprof.Register(g)

	// api for authentication functionalities
	g.POST("/login", user.Login)

	// The user handlers, requiring authentication
	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
		u.GET("/:username", user.Get)
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	// 手机验证码发送
	captchas := g.Group("/captcha")
	{
		// todo:后面需要添加短信账号的信息，即返回账号名称与余额
		captchas.GET("/mobile_obtain", captcha.MobileObtain)
		captchas.GET("/img_obtain", captcha.ImgObtain)

	}

	// 发送电子邮件
	emails := g.Group("/email")
	{
		emails.GET("/send", email.Send)
	}

	// 获取指定字符串的二维码图片
	qrcodes := g.Group("/qrcode")
	{
		qrcodes.POST("/obtain", qrcode.Obtain)
	}

	return g
}
