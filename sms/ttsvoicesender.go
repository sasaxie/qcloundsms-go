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

const sendTVoiceUrl = "https://cloud.tim.qq.com/v5/tlsvoicesvr/sendtvoice?sdkappid=%d&random=%d"

type TtsVoiceSender struct {
	Base
}

func NewTtsVoiceSender(appID int, appKey string) *TtsVoiceSender {
	return &TtsVoiceSender{
		Base{
			AppID:  appID,
			AppKey: appKey,
		},
	}
}

/*
 * 指定模版发送语音短信
 * nationCode 国家码，如 86 为中国
 * phoneNumber 不带国家码的手机号
 * templateId 信息内容
 * params 模版参数列表，如模版 {1}...{2}...{3}，那么需要带三个参数
 * playTimes 播放次数
 * ext 服务端原样返回的参数，可填空
 */
func (t *TtsVoiceSender) Send(nationCode, phoneNumber string, templateId int, params []string, playTimes int, ext string) (*TtsVoiceSenderResult, error) {
	random := util.GetRandom()
	now := util.GetCurrentTime()

	type Tel struct {
		NationCode string `json:"nationcode"`
		Mobile     string `json:"mobile"`
	}

	type Body struct {
		Tel       *Tel     `json:"tel"`
		TplId     int      `json:"tpl_id"`
		Params    []string `json:"params"`
		PlayTimes int      `json:"playtimes"`
		Sig       string   `json:"sig"`
		Time      int64    `json:"time"`
		Ext       string   `json:"ext,omitempty"`
	}

	body := new(Body)
	body.Tel = &Tel{
		NationCode: nationCode,
		Mobile:     phoneNumber,
	}
	body.TplId = templateId
	body.Params = params
	body.PlayTimes = playTimes
	body.Sig = util.CalculateSignatureWithPhoneNumber(t.AppKey, random, now, phoneNumber)
	body.Time = now
	body.Ext = ext

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Timeout: 60 * time.Second,
	}

	req, err := http.NewRequest("POST", fmt.Sprintf(sendTVoiceUrl, t.AppID, random), bytes.NewBuffer(bodyJson))
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

	result := new(TtsVoiceSenderResult)
	err = result.ParseFromHTTPResponseBody(b)
	if err != nil {
		return nil, err
	}

	return result, nil
}
