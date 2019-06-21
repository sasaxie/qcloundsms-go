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

const pullStatus4MobileUrl = "https://yun.tim.qq.com/v5/tlssmssvr/pullstatus4mobile?sdkappid=%d&random=%d"

type MobileStatusPuller struct {
	Base
}

func NewMobileStatusPuller(appID int, appKey string) *MobileStatusPuller {
	return &MobileStatusPuller{
		Base{
			AppID:  appID,
			AppKey: appKey,
		},
	}
}

func (m *MobileStatusPuller) pull(smsType int, nationCode, mobile string, beginTime, endTime int64, max int) ([]byte, error) {
	random := util.GetRandom()
	now := util.GetCurrentTime()

	type Body struct {
		Sig        string `json:"sig"`
		Type       int    `json:"type"`
		Time       int64  `json:"time"`
		Max        int    `json:"max"`
		BeginTime  int64  `json:"begin_time"`
		EndTime    int64  `json:"end_time"`
		NationCode string `json:"nationcode"`
		Mobile     string `json:"mobile"`
	}

	body := new(Body)
	body.Sig = util.CalculateSignature(m.AppKey, random, now)
	body.Type = smsType
	body.Time = now
	body.Max = max
	body.BeginTime = beginTime
	body.EndTime = endTime
	body.NationCode = nationCode
	body.Mobile = mobile

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Timeout: 60 * time.Second,
	}

	req, err := http.NewRequest("POST", fmt.Sprintf(pullStatus4MobileUrl, m.AppID, random), bytes.NewBuffer(bodyJson))
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

	return b, nil
}

func (m *MobileStatusPuller) PullCallback(nationCode, mobile string, beginTime, endTime int64, max int) (*StatusPullCallbackResult, error) {
	b, err := m.pull(0, nationCode, mobile, beginTime, endTime, max)
	if err != nil {
		return nil, err
	}

	result := new(StatusPullCallbackResult)
	err = result.ParseFromHTTPResponseBody(b)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *MobileStatusPuller) PullReply(nationCode, mobile string, beginTime, endTime int64, max int) (*StatusPullReplyResult, error) {
	b, err := m.pull(1, nationCode, mobile, beginTime, endTime, max)
	if err != nil {
		return nil, err
	}

	result := new(StatusPullReplyResult)
	err = result.ParseFromHTTPResponseBody(b)
	if err != nil {
		return nil, err
	}

	return result, nil
}
