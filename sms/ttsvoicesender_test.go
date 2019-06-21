package sms

import (
	"testing"
)

func TestTtsVoiceSender_Send(t *testing.T) {
	// 短信应用SDK AppID
	appID := 1400009099 // 1400开头

	// 短信应用SDK AppKey
	appKey := "9ff91d87c2cd7cd0ea762f141975d1df37481d48700d70ac37470aefc60f9bad"

	// 短信模版内容
	params := make([]string, 0)
	params = append(params, "5789")

	// 需要发送短信的手机号码
	phoneNumber := "12345678902"

	// 短信模版ID，需要在短信应用中申请
	// NOTE: 这里的模版`7839`只是一个示例，
	// 真实的模版ID需要在短信控制台中申请
	templateId := 7839

	// 播放次数
	// NOTE: 可选，最多3次，默认2次
	playTimes := 1

	ttsVoiceSender := NewTtsVoiceSender(appID, appKey)

	result, err := ttsVoiceSender.Send("86", phoneNumber, templateId, params, playTimes, "")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(result)
	}
}
