package sms

import (
	"io/ioutil"
	"testing"
)

func TestVoiceFileUploader_Upload(t *testing.T) {
	// 短信应用SDK AppID
	appID := 1400009099 // 1400开头

	// 短信应用SDK AppKey
	appKey := "9ff91d87c2cd7cd0ea762f141975d1df37481d48700d70ac37470aefc60f9bad"

	// Note: 语音文件大小上传限制400K字节
	filename := "../resource/example.mp3"
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error(err)
		return
	}

	uploader := NewVoiceFileUploader(appID, appKey)
	result, err := uploader.Upload(fileContent, MP3)
	if err != nil {
		t.Error(err)
		return
	} else {
		t.Log(result)
	}
}
