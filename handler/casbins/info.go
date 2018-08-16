package casbins

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Info(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}
