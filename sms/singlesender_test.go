package sms

import (
	"testing"
)

func TestSingleSender_Send(t *testing.T) {
	// 短信应用SDK AppID
	appID := 1400009099 // 1400开头

	// 短信应用SDK AppKey
	appKey := "9ff91d87c2cd7cd0ea762f141975d1df37481d48700d70ac37470aefc60f9bad"

	// 需要发送短信的手机号码
	phoneNumber := "12345678902"

	singleSender := NewSingleSender(appID, appKey)
	result, err := singleSender.Send(0, "86", phoneNumber, "【腾讯云】您的验证码是: 5678", "", "")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(result)
	}
}

func TestSingleSender_SendWithParam(t *testing.T) {
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

	// 签名
	// NOTE: 这里的签名"腾讯云"只是一个示例，
	// 真实的签名需要在短信控制台中申请，另外
	// 签名参数使用的示`签名内容`，而不是`签名ID`
	smsSign := "腾讯云"

	singleSender := NewSingleSender(appID, appKey)

	result, err := singleSender.SendWithParam("86", phoneNumber, templateId, params, smsSign, "", "")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(result)
	}
}
