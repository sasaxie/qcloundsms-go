package sms

import (
	"testing"
	"time"
)

func TestMobileStatusPuller_PullCallback(t *testing.T) {
	// 短信应用SDK AppID
	appID := 1400009099 // 1400开头

	// 短信应用SDK AppKey
	appKey := "9ff91d87c2cd7cd0ea762f141975d1df37481d48700d70ac37470aefc60f9bad"

	beginTime := time.Now().Unix() - 3600
	endTime := time.Now().Unix()
	maxNum := 10 // 单次拉取最大量
	puller := NewMobileStatusPuller(appID, appKey)
	result, err := puller.PullCallback("86", "12345678902", beginTime, endTime, maxNum)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(result)
	}
}

func TestMobileStatusPuller_PullReply(t *testing.T) {
	// 短信应用SDK AppID
	appID := 1400009099 // 1400开头

	// 短信应用SDK AppKey
	appKey := "9ff91d87c2cd7cd0ea762f141975d1df37481d48700d70ac37470aefc60f9bad"

	beginTime := time.Now().Unix() - 3600
	endTime := time.Now().Unix()
	maxNum := 10 // 单次拉取最大量
	puller := NewMobileStatusPuller(appID, appKey)
	result, err := puller.PullReply("86", "12345678902", beginTime, endTime, maxNum)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(result)
	}
}
