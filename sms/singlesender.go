package sms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sasaxie/qcloundsms-go/util"
	"io/ioutil"
	"net/http"
	"time"
)

const singleSmsUrl = "https://yun.tim.qq.com/v5/tlssmssvr/sendsms?sdkappid=%d&random=%d"

type SingleSender struct {
	Base
}

func NewSingleSender(appID int, appKey string) *SingleSender {
	return &SingleSender{
		Base{
			AppID:  appID,
			AppKey: appKey,
		},
	}
}

/*
 * 普通单发
 * smsType 短信类型，0 为普通短信，1 为营销短信
 * nationCode 国家码，如 86 为中国
 * phoneNumber 不带国家码的手机号
 * msg 信息内容，必须与申请的模版格式一致，否则将返回错误
 * extend 扩展码，可填空
 * ext 服务端原样返回的参数，可填空
 */
func (s *SingleSender) Send(smsType int, nationCode, phoneNumber, msg, extend, ext string) (*SingleSenderResult, error) {
	random := util.GetRandom()
	now := util.GetCurrentTime()

	type Tel struct {
		NationCode string `json:"nationcode"`
		Mobile     string `json:"mobile"`
	}

	type Body struct {
		Tel    *Tel   `json:"tel"`
		Type   int    `json:"type"`
		Msg    string `json:"msg"`
		Sig    string `json:"sig"`
		Time   int64  `json:"time"`
		Extend string `json:"extend,omitempty"`
		Ext    string `json:"ext,omitempty"`
	}

	body := new(Body)
	body.Tel = &Tel{
		NationCode: nationCode,
		Mobile:     phoneNumber,
	}
	body.Type = smsType
	body.Msg = msg
	body.Sig = util.CalculateSignatureWithPhoneNumber(s.AppKey, random, now, phoneNumber)
	body.Time = now
	body.Extend = extend
	body.Ext = ext

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Timeout: 60 * time.Second,
	}

	req, err := http.NewRequest("POST", fmt.Sprintf(singleSmsUrl, s.AppID, random), bytes.NewBuffer(bodyJson))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := new(SingleSenderResult)
	err = result.ParseFromHTTPResponseBody(b)
	if err != nil {
		return nil, err
	}

	return result, nil
}

/*
 * 指定模版单发
 * nationCode 国家码，如 86 为中国
 * phoneNumber 不带国家码的手机号
 * templateId 信息内容
 * params 模版参数列表，如模版 {1}...{2}...{3}，那么需要带三个参数
 * sign 签名，如果填空，系统会使用默认签名
 * extend 扩展码，可填空
 * ext 服务端原样返回的参数，可填空
 */
func (s *SingleSender) SendWithParam(nationCode, phoneNumber string, templateId int, params []string, sign, extend, ext string) (*SingleSenderResult, error) {
	random := util.GetRandom()
	now := util.GetCurrentTime()

	type Tel struct {
		NationCode string `json:"nationcode"`
		Mobile     string `json:"mobile"`
	}

	type Body struct {
		Tel    *Tel     `json:"tel"`
		Sig    string   `json:"sig"`
		TplId  int      `json:"tpl_id"`
		Params []string `json:"params"`
		Sign   string   `json:"sign,omitempty"`
		Time   int64    `json:"time"`
		Extend string   `json:"extend,omitempty"`
		Ext    string   `json:"ext,omitempty"`
	}

	body := new(Body)
	body.Tel = &Tel{
		NationCode: nationCode,
		Mobile:     phoneNumber,
	}
	body.Sig = util.CalculateSignatureWithPhoneNumber(s.AppKey, random, now, phoneNumber)
	body.TplId = templateId
	body.Params = params
	body.Sign = sign
	body.Time = now
	body.Extend = extend
	body.Ext = ext

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Timeout: 60 * time.Second,
	}

	req, err := http.NewRequest("POST", fmt.Sprintf(singleSmsUrl, s.AppID, random), bytes.NewBuffer(bodyJson))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := new(SingleSenderResult)
	err = result.ParseFromHTTPResponseBody(b)
	if err != nil {
		return nil, err
	}

	return result, nil
}
