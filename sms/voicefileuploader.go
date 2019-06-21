package sms

import (
	"bytes"
	"fmt"
	_ "github.com/paulrosania/go-charset/data"
	"github.com/sasaxie/qcloundsms-go/util"
	"io/ioutil"
	"net/http"
	"time"
)

const uploadVoiceFileUrl = "https://cloud.tim.qq.com/v5/tlsvoicesvr/uploadvoicefile?sdkappid=%d&random=%d&time=%d"

type ContentType int

const (
	WAV ContentType = iota
	MP3
)

type VoiceFileUploader struct {
	Base
}

func NewVoiceFileUploader(appID int, appKey string) *VoiceFileUploader {
	return &VoiceFileUploader{
		Base{
			AppID:  appID,
			AppKey: appKey,
		},
	}
}

/*
 * 上传语音文件
 * fileContent 语音文件内容
 * contentType 语音文件类型
 */
func (v *VoiceFileUploader) Upload(fileContent []byte, contentType ContentType) (*VoiceFileUploaderResult, error) {
	random := util.GetRandom()
	now := util.GetCurrentTime()

	fileSha1Sum := util.Sha1Sum(fileContent)
	auth := util.CalculateAuth(v.AppKey, random, now, fileSha1Sum)

	fileType := "audio/mpeg"
	if contentType == WAV {
		fileType = "audio/wav"
	}

	httpClient := &http.Client{
		Timeout: 60 * time.Second,
	}

	req, err := http.NewRequest("POST", fmt.Sprintf(uploadVoiceFileUrl, v.AppID, random, now), bytes.NewBuffer(fileContent))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", fileType)
	req.Header.Add("x-content-sha1", fileSha1Sum)
	req.Header.Add("Authorization", auth)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := new(VoiceFileUploaderResult)
	err = result.ParseFromHTTPResponseBody(b)
	if err != nil {
		return nil, err
	}

	return result, nil
}
