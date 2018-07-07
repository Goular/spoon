package qrcode

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
	"encoding/base64"
	"spoon/handler"
)

// 提交POST请求获取二维码
func Obtain(c *gin.Context) {
	url := c.PostForm("url")
	imgWidth := c.PostForm("img_width")
	imgHeight := c.PostForm("img_height")
	area := imgWidth + "x" + imgHeight
	rUrl := "https://api.qrserver.com/v1/create-qr-code/?size=" + area + "&data=" + url
	resp, err := http.Get(rUrl)
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	encodeString := base64.StdEncoding.EncodeToString(body)
	maps := make(map[string]string)
	maps["imgUrl"] = "data:image/png;base64," + encodeString
	handler.SendResponse(c, nil, maps)
}
