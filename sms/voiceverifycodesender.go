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

const sendVoiceUrl = "https://cloud.tim.qq.com/v5/tlsvoicesvr/sendcvoice?sdkappid=%d&random=%d"

type VoiceVerifyCodeSender struct {
	Base
}

func NewVoiceVerifyCodeSender(appID int, appKey string) *VoiceVerifyCodeSender {
	return &VoiceVerifyCodeSender{
		Base{
			AppID:  appID,
			AppKey: appKey,
		},
	}
}

/*
 * 发送语音短信
 * nationCode 国家码，如 86 为中国
 * phoneNumber 不带国家码的手机号
 * msg 信息类型
 * playTimes 播放次数
 * ext 服务端原样返回的参数，可填空
 */
func (v *VoiceVerifyCodeSender) Send(nationCode, phoneNumber, msg string, playTimes int, ext string) (*VoiceVerifyCodeSenderResult, error) {
	random := util.GetRandom()
	now := util.GetCurrentTime()

	type Tel struct {
		NationCode string `json:"nationcode"`
		Mobile     string `json:"mobile"`
	}

	type Body struct {
		Tel       *Tel   `json:"tel"`
		Msg       string `json:"msg"`
		PlayTimes int    `json:"playtimes"`
		Sig       string `json:"sig"`
		Time      int64  `json:"time"`
		Ext       string `json:"ext,omitempty"`
	}

	body := new(Body)
	body.Tel = &Tel{
		NationCode: nationCode,
		Mobile:     phoneNumber,
	}
	body.Msg = msg
	body.PlayTimes = playTimes
	body.Sig = util.CalculateSignatureWithPhoneNumber(v.AppKey, random, now, phoneNumber)
	body.Time = now
	body.Ext = ext

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Timeout: 60 * time.Second,
	}

	req, err := http.NewRequest("POST", fmt.Sprintf(sendVoiceUrl, v.AppID, random), bytes.NewBuffer(bodyJson))
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

	result := new(VoiceVerifyCodeSenderResult)
	err = result.ParseFromHTTPResponseBody(b)
	if err != nil {
		return nil, err
	}

	return result, nil
}
