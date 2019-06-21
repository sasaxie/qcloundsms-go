package sms

import "testing"

func TestStatusPuller_PullCallback(t *testing.T) {
	// 短信应用SDK AppID
	appID := 1400009099 // 1400开头

	// 短信应用SDK AppKey
	appKey := "9ff91d87c2cd7cd0ea762f141975d1df37481d48700d70ac37470aefc60f9bad"

	// Note: 短信拉取功能需要联系腾讯云短信技术支持(QQ:3012203387)开通权限
	maxNum := 10 // 单次拉取最大量
	puller := NewStatusPuller(appID, appKey)
	result, err := puller.PullCallback(maxNum)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(result)
	}
}

func TestStatusPuller_PullReply(t *testing.T) {
	// Note: 短信拉取功能需要联系腾讯云短信技术支持(QQ:3012203387)开通权限
	// 短信应用SDK AppID
	appID := 1400009099 // 1400开头

	// 短信应用SDK AppKey
	appKey := "9ff91d87c2cd7cd0ea762f141975d1df37481d48700d70ac37470aefc60f9bad"

	// Note: 短信拉取功能需要联系腾讯云短信技术支持(QQ:3012203387)开通权限
	maxNum := 10 // 单次拉取最大量
	puller := NewStatusPuller(appID, appKey)
	result, err := puller.PullReply(maxNum)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(result)
	}
}
