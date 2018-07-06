package email

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/url"
	"bytes"
	"net/http"
	"io/ioutil"
)

func Send(c *gin.Context) {
	c.JSON(http.StatusOK, SendMail())
}

func SendMail() string {
	RequestURI := "http://api.sendcloud.net/apiv2/mail/send"
	PostParams := url.Values{
		"apiUser":  {"goular_test_9aYMZ6"},
		"apiKey":   {"9FQ9e96HnPHF0adL"},
		"from":     {"postman@jiagongwu.com"},
		"fromName": {"goggo"},
		"to":       {"zhaojt_exam@126.com"},
		"subject":  {"吔屎啦，梁非凡"},
		"html":     {"哈哈梁非凡"},
	}
	PostBody := bytes.NewBufferString(PostParams.Encode())
	ResponseHandler, err := http.Post(RequestURI, "application/x-www-form-urlencoded", PostBody)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ResponseHandler.Body.Close()
	BodyByte, err := ioutil.ReadAll(ResponseHandler.Body)
	if err != nil {
		panic(err)
	}
	return string(BodyByte)
}
