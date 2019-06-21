package sms

import "testing"

func TestVoiceVerifySender_Send(t *testing.T) {
	// 短信应用SDK AppID
	appID := 1400009099 // 1400开头

	// 短信应用SDK AppKey
	appKey := "9ff91d87c2cd7cd0ea762f141975d1df37481d48700d70ac37470aefc60f9bad"

	sender := NewVoiceVerifyCodeSender(appID, appKey)
	result, err := sender.Send("86", "12345678902", "5678", 2, "")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(result)
	}
}
