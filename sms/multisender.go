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

const sendMultiSmsUrl = "https://yun.tim.qq.com/v5/tlssmssvr/sendmultisms2?sdkappid=%d&random=%d"

type MultiSender struct {
	Base
}

func NewMultiSender(appID int, appKey string) *MultiSender {
	return &MultiSender{
		Base{
			AppID:  appID,
			AppKey: appKey,
		},
	}
}

/*
 * 普通群发：明确指定内容，如果有多个签名，请在内容中以【】的方式添加到信息内容中，否则系统将使用默认签名
 * smsType 短信类型，0 为普通短信，1 为营销短信
 * nationCode 国家码，如 86 为中国
 * phoneNumbers 不带国家码的手机号列表
 * msg 信息内容，必须与申请的目标格式一致，否则将返回错误
 * extend 扩展码，可填空
 * ext 服务端原样返回的参数，可填空
 */
func (m *MultiSender) Send(smsType int, nationCode string, phoneNumbers []string, msg, extend, ext string) (*MultiSenderResult, error) {
	random := util.GetRandom()
	now := util.GetCurrentTime()

	type Tel struct {
		NationCode string `json:"nationcode"`
		Mobile     string `json:"mobile"`
	}

	type Body struct {
		Tel    []*Tel `json:"tel"`
		Type   int    `json:"type"`
		Msg    string `json:"msg"`
		Sig    string `json:"sig"`
		Time   int64  `json:"time"`
		Extend string `json:"extend,omitempty"`
		Ext    string `json:"ext,omitempty"`
	}

	body := new(Body)
	body.Tel = make([]*Tel, 0)
	for _, phoneNumber := range phoneNumbers {
		tel := new(Tel)
		tel.NationCode = nationCode
		tel.Mobile = phoneNumber
		body.Tel = append(body.Tel, tel)
	}

	body.Type = smsType
	body.Msg = msg
	body.Sig = util.CalculateSignatureWithPhoneNumbers(m.AppKey, random, now, phoneNumbers)
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

	req, err := http.NewRequest("POST", fmt.Sprintf(sendMultiSmsUrl, m.AppID, random), bytes.NewBuffer(bodyJson))
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

	result := new(MultiSenderResult)
	err = result.ParseFromHTTPResponseBody(b)
	if err != nil {
		return nil, err
	}

	return result, nil
}

/*
 * 指定模版群发
 * nationCode 国家码，如 86 为中国
 * phoneNumbers 不带国家码的手机号列表
 * templateId 信息内容
 * params 模版参数列表，如模版 {1}...{2}...{3}，那么需要带三个参数
 * sign 签名，如果填空，系统会使用默认签名
 * extend 扩展码，可填空
 * ext 服务端原样返回的参数，可填空
 */
func (m *MultiSender) SendWithParam(nationCode string, phoneNumbers []string, templateId int, params []string, sign, extend, ext string) (*MultiSenderResult, error) {
	random := util.GetRandom()
	now := util.GetCurrentTime()

	type Tel struct {
		NationCode string `json:"nationcode"`
		Mobile     string `json:"mobile"`
	}

	type Body struct {
		Tel    []*Tel   `json:"tel"`
		Sign   string   `json:"sign,omitempty"`
		TplId  int      `json:"tpl_id"`
		Params []string `json:"params"`
		Sig    string   `json:"sig"`
		Time   int64    `json:"time"`
		Extend string   `json:"extend,omitempty"`
		Ext    string   `json:"ext,omitempty"`
	}

	body := new(Body)
	body.Tel = make([]*Tel, 0)
	for _, phoneNumber := range phoneNumbers {
		tel := new(Tel)
		tel.NationCode = nationCode
		tel.Mobile = phoneNumber
		body.Tel = append(body.Tel, tel)
	}

	body.Sign = sign
	body.TplId = templateId
	body.Params = params
	body.Sig = util.CalculateSignatureWithPhoneNumbers(m.AppKey, random, now, phoneNumbers)
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

	req, err := http.NewRequest("POST", fmt.Sprintf(sendMultiSmsUrl, m.AppID, random), bytes.NewBuffer(bodyJson))
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

	result := new(MultiSenderResult)
	err = result.ParseFromHTTPResponseBody(b)
	if err != nil {
		return nil, err
	}

	return result, nil
}
