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

const statusVoiceFileUrl = "https://cloud.tim.qq.com/v5/tlsvoicesvr/statusvoicefile?sdkappid=%d&random=%d"

type StatusVoiceFile struct {
	Base
}

func NewStatusVoiceFile(appID int, appKey string) *StatusVoiceFile {
	return &StatusVoiceFile{
		Base{
			AppID:  appID,
			AppKey: appKey,
		},
	}
}

/*
 * 查询语音文件审核状态
 * fid 语音文件fid
 */
func (s *StatusVoiceFile) Get(fid string) (*VoiceFileStatusResult, error) {
	random := util.GetRandom()
	now := util.GetCurrentTime()

	type Body struct {
		Fid  string `json:"fid"`
		Sig  string `json:"sig"`
		Time int64  `json:"time"`
	}

	body := new(Body)
	body.Fid = fid
	body.Sig = util.CalculateSignatureWithFid(s.AppKey, random, now, fid)
	body.Time = now

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Timeout: 60 * time.Second,
	}

	req, err := http.NewRequest("POST", fmt.Sprintf(statusVoiceFileUrl, s.AppID, random), bytes.NewBuffer(bodyJson))
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

	result := new(VoiceFileStatusResult)
	err = result.ParseFromHTTPResponseBody(b)
	if err != nil {
		return nil, err
	}

	return result, nil
}
