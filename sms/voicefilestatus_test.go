package sms

import "testing"

func TestStatusVoiceFile_Get(t *testing.T) {
	// 短信应用SDK AppID
	appID := 1400009099 // 1400开头

	// 短信应用SDK AppKey
	appKey := "9ff91d87c2cd7cd0ea762f141975d1df37481d48700d70ac37470aefc60f9bad"

	// Note: 这里fid来自`上传语音文件`接口返回的响应，要按语音
	// 文件fid发送语音通知，需要先上传语音文件获取fid
	fid := "2d86ecd5cee47fbe8fb06b358e334facef44bf77.mp3"

	file := NewStatusVoiceFile(appID, appKey)

	// result里会带有语音文件审核状态status, {0: 待审核, 1: 通过, 2: 拒绝, 3: 语音文件不存在}
	result, err := file.Get(fid)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(result)
	}
}
