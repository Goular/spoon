package captcha

import (
	"github.com/gin-gonic/gin"
	ypclnt "github.com/yunpian/yunpian-go-sdk/sdk"
	"net/http"
	"github.com/spf13/viper"
)

// 通过httpclient发送到第三方服务商请求手机验证码
// {
//    "msg": "发送成功",
//    "data": {
//        "code": 0,
//        "count": 1,
//        "fee": 0.05,
//        "mobile": "15800279296",
//        "msg": "发送成功",
//        "sid": 25783494258,
//        "unit": "RMB"
//    }
//}

// https://www.yunpian.com/doc/zh_CN/introduction/demos/go.html
// package main
//import (
//    "net/http"
//    "io/ioutil"
//    "net/url"
//    "fmt"
// )
//// bingone
//func main(){
//
//    // 修改为您的apikey(https://www.yunpian.com)登录官网后获取
//    apikey      := "xxxxxxxxxxxxxxxxxx"
//     // 修改为您要发送的手机号码，多个号码用逗号隔开
//    mobile      := "xxxxxxxxxxxxxxxxxx"
//    // 发送内容
//    text        := "【云片网】您的验证码是1234"
//    // 发送模板编号
//    tpl_id      := 1
//    // 语音验证码
//    code        := "1234"
//    company     := "云片网"
//    // 发送模板内容
//    tpl_value   := url.Values{"#code#":{code},"#company#":{company}}.Encode()
//
//    // 获取user信息url
//    url_get_user    := "https://sms.yunpian.com/v2/user/get.json";
//    // 智能模板发送短信url
//    url_send_sms    := "https://sms.yunpian.com/v2/sms/single_send.json";
//    // 指定模板发送短信url
//    url_tpl_sms     := "https://sms.yunpian.com/v2/sms/tpl_single_send.json";
//    // 发送语音短信url
//    url_send_voice  := "https://voice.yunpian.com/v2/voice/send.json";
//
//    data_get_user   := url.Values{"apikey": {apikey}}
//    data_send_sms   := url.Values{"apikey": {apikey}, "mobile": {mobile},"text":{text}}
//    data_tpl_sms := url.Values { "apikey": {apikey},"mobile": {mobile},
//        "tpl_id": {fmt.Sprintf("%d", tpl_id)},"tpl_value": {tpl_value}}
//    data_send_voice := url.Values{"apikey": {apikey}, "mobile": {mobile},"code":{code}}
//
//
//    httpsPostForm(url_get_user,data_get_user)
//    httpsPostForm(url_send_sms,data_send_sms)
//    httpsPostForm(url_tpl_sms,data_tpl_sms)
//    httpsPostForm(url_send_voice,data_send_voice)
//}
//
//func httpsPostForm(url string,data url.Values) {
//    resp, err := http.PostForm(url,data)
//
//    if err != nil {
//        // handle error
//    }
//
//    defer resp.Body.Close()
//    body, err := ioutil.ReadAll(resp.Body)
//    if err != nil {
//        // handle error
//    }
//
//    fmt.Println(string(body))
//
//}

// 手机验证码获取
func MobileObtain(c *gin.Context) {
	// 发送短信
	clnt := ypclnt.New(viper.GetString("yunpian_sms.apikey"))
	param := ypclnt.NewParam(2)
	param[ypclnt.MOBILE] = "15800279296"
	param[ypclnt.TEXT] = "【加工屋】您的验证码是1234"
	r := clnt.Sms().SingleSend(param)
	c.JSON(http.StatusOK, r)
}
