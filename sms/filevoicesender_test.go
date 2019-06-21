package sms

import "testing"

func TestFileVoiceSender_Send(t *testing.T) {
	// 短信应用SDK AppID
	appID := 1400009099 // 1400开头

	// 短信应用SDK AppKey
	appKey := "9ff91d87c2cd7cd0ea762f141975d1df37481d48700d70ac37470aefc60f9bad"

	// Note: 这里fid来自`上传语音文件`接口返回的响应，要按语音
	// 文件fid发送语音通知，需要先上传语音文件获取fid
	fid := "2d86ecd5cee47fbe8fb06b358e334facef44bf77.mp3"

	sender := NewFileVoiceSender(appID, appKey)

	result, err := sender.Send("86", "12345678902", fid, 2, "")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(result)
	}
}
