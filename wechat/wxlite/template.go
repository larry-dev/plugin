package wxlite

import (
	"net/http"
	"fmt"
	"bytes"
	"io/ioutil"
	"encoding/json"
)

func (wx *WXLite) SendTemplate(token string, data map[string]interface{}) error {
	bs, _ := json.Marshal(data)
	req := bytes.NewBuffer(bs)
	bodyType := "application/json;charset=utf-8"
	resp, _ := http.Post(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=%s", token), bodyType, req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return nil
}
