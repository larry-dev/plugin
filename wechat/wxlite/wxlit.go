package wxlite

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UserInfo struct {
	Nickname  string `json:"nickName"`
	AvatarUrl string `json:"avatarUrl"`
	Gender    uint    `json:"gender"`
	City      string `json:"city"`
	Province  string `json:"province"`
	Country   string `json:"country"`
	Language  string `json:"language"`
}

type Session struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	ErrCode    string `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}
type WXLite struct {
	AppId     string
	AppSecret string
}

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrCode     string `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

func (wx *WXLite) GetSession(jsCode string) (*Session, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", wx.AppId, wx.AppSecret, jsCode))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Http request is error.status code (%d)", resp.StatusCode))
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	sess := &Session{}
	if err = json.Unmarshal(result, sess); err != nil {
		return nil, err
	}
	return sess, nil
}
func (wx *WXLite) GetAccessToken() (*Token, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", wx.AppId, wx.AppSecret))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Http request is error.status code (%d)", resp.StatusCode))
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	token := &Token{}
	if err = json.Unmarshal(result, token); err != nil {
		return nil, err
	}
	return token, nil
}
