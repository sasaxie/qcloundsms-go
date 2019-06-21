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

const pullStatusUrl = "https://yun.tim.qq.com/v5/tlssmssvr/pullstatus?sdkappid=%d&random=%d"

type StatusPuller struct {
	Base
}

func NewStatusPuller(appID int, appKey string) *StatusPuller {
	return &StatusPuller{
		Base{
			AppID:  appID,
			AppKey: appKey,
		},
	}
}

func (s *StatusPuller) pull(smsType, max int) ([]byte, error) {
	random := util.GetRandom()
	now := util.GetCurrentTime()

	type Body struct {
		Sig  string `json:"sig"`
		Time int64  `json:"time"`
		Type int    `json:"type"`
		Max  int    `json:"max"`
	}

	body := new(Body)
	body.Sig = util.CalculateSignature(s.AppKey, random, now)
	body.Time = now
	body.Type = smsType
	body.Max = max

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Timeout: 60 * time.Second,
	}

	req, err := http.NewRequest("POST", fmt.Sprintf(pullStatusUrl, s.AppID, random), bytes.NewBuffer(bodyJson))
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

func (s *StatusPuller) PullCallback(max int) (*StatusPullCallbackResult, error) {
	b, err := s.pull(0, max)
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

func (s *StatusPuller) PullReply(max int) (*StatusPullReplyResult, error) {
	b, err := s.pull(1, max)
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
