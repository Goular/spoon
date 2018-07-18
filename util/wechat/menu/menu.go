package menu

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"bytes"
	"github.com/silenceper/wechat"
	"encoding/json"
	"strconv"
	"github.com/pkg/errors"
)

// 由于github.com/silenceper/wechat创建菜单方式比较难用，所以在这里在封装一个创建方法

const (
	menuCreateURL = "https://api.weixin.qq.com/cgi-bin/menu/create"
	//menuGetURL               = "https://api.weixin.qq.com/cgi-bin/menu/get"
	//menuDeleteURL            = "https://api.weixin.qq.com/cgi-bin/menu/delete"
	//menuAddConditionalURL    = "https://api.weixin.qq.com/cgi-bin/menu/addconditional"
	//menuDeleteConditionalURL = "https://api.weixin.qq.com/cgi-bin/menu/delconditional"
	//menuTryMatchURL          = "https://api.weixin.qq.com/cgi-bin/menu/trymatch"
	//menuSelfMenuInfoURL      = "https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info"
)

// 直接将前端返回的json参数POST到微信服务器进行创建菜单
func CreateMenu(wechat *wechat.Wechat, str string) error {
	// 获取access_token
	access_token, err := wechat.GetAccessToken()
	if err != nil {
		return err
	}
	// 拼接访问的网址
	uri := fmt.Sprintf("%s?access_token=%s", menuCreateURL, access_token)
	resp, err := http.Post(uri, "application/json", bytes.NewBuffer([]byte(str)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	result := make(map[string]string)
	json.Unmarshal(body, &result)
	errCodeStr, ok := result["errcode"]
	if ok {
		errCode, err := strconv.Atoi(errCodeStr)
		if err != nil {
			return err
		} else {
			if errCode != 0 {
				return errors.New(string(body))
			}
		}
	}
	return nil
}
