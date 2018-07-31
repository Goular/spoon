package turingapi

import (
	"sync"
	"net/http"
	"io/ioutil"
	"github.com/spf13/viper"
	"encoding/json"
	"bytes"
)

const RobotURL = "http://openapi.tuling123.com/openapi/api/v2"

// 2018-07-31 添加图灵机器人库的使用
// 添加文本对话
func ChatRobotWithText(query string, location *Location) (error, []byte) {
	requestBody := &Body{
		ReqType:    0,
		Perception: structureTextPerception(query, location),
		UserInfo:   structureUserInfo(),
	}
	requestJson, err1 := json.Marshal(requestBody)
	if err1 != nil {
		return err1, nil
	}
	return chatRobot(requestJson)
}

// 处理返回的JSON字符串，同时输出内容
func GetResponseTxt(result []byte) string {
	var text interface{}
	err2 := json.Unmarshal(result, &text)
	if err2 != nil {
		return "返回数据异常"
	}
	var resultStr string
	if v, ok := text.(map[string]interface{}); ok {
		if results, ok := v["results"]; ok {
			if v2, ok2 := results.([]interface{}); ok2 {
				if v3, ok3 := v2[0].(map[string]interface{}); ok3 {
					if v4, ok4 := v3["values"].(map[string]interface{}); ok4 {
						resultStr = v4["text"].(string)
					}
				}
			}
		}
	}
	return resultStr
}

// 添加图片对话 -- 说真的还是不太灵光
func ChatRobotWithImage(text, imgUrl string, location *Location) (error, []byte) {
	requestBody := &Body{
		ReqType:    0,
		Perception: structureImagePerception(text, imgUrl, location),
		UserInfo:   structureUserInfo(),
	}
	requestJson, err1 := json.Marshal(requestBody)
	if err1 != nil {
		return err1, nil
	}
	return chatRobot(requestJson)
}

// 添加音频对话
func ChatRobotWithAudio(text, audioUrl string, location *Location) (error, []byte) {
	requestBody := &Body{
		ReqType:    0,
		Perception: structureAudioPerception(text, audioUrl, location),
		UserInfo:   structureUserInfo(),
	}
	requestJson, err1 := json.Marshal(requestBody)
	if err1 != nil {
		return err1, nil
	}
	return chatRobot(requestJson)
}

// 统一提交URL请求的类
func chatRobot(requestJson []byte) (error, []byte) {
	// 用于阻塞
	wg := &sync.WaitGroup{}
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)
	var result []byte
	wg.Add(1)
	go func() {
		defer wg.Done()
		resp, err := http.Post(RobotURL, "application/json", bytes.NewReader(requestJson))
		if err != nil {
			errChan <- err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			errChan <- err
		}
		result = body
	}()
	go func() {
		wg.Wait()
		close(finished)
	}()
	select {
	case <-finished:
	case err := <-errChan:
		return err, nil
	}
	return nil, result
}

// 构建用户信息
func structureUserInfo() *UserInfo {
	return &UserInfo{
		ApiKey: viper.GetString("turingapi.apiKey"),
		UserId: viper.GetString("turingapi.userId"),
	}
}

// 构建文本信息
func structureTextPerception(text string, location *Location) *Query {
	return &Query{
		InputText: struct{ Text string `json:"text"` }{Text: text},
		SelfInfo:  struct{ Location *Location `json:"location"` }{Location: location},
	}
}

// 构建图片信息
func structureImagePerception(text, url string, location *Location) *Query {
	return &Query{
		InputText:  struct{ Text string `json:"text"` }{Text: text},
		InputImage: struct{ Url string `json:"url"` }{Url: url},
		SelfInfo:   struct{ Location *Location `json:"location"` }{Location: location},
	}
}

// 构建音频信息
func structureAudioPerception(text, url string, location *Location) *Query {
	return &Query{
		InputText:  struct{ Text string `json:"text"` }{Text: text},
		InputMedia: struct{ Url string `json:"url"` }{Url: url},
		SelfInfo:   struct{ Location *Location `json:"location"` }{Location: location},
	}
}
