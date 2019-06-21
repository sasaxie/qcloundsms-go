package sms

import (
	"encoding/json"
	"fmt"
)

type VoiceVerifyCodeSenderResult struct {
	Result int
	ErrMsg string
	Ext    string
	CallId string
}

func (v *VoiceVerifyCodeSenderResult) String() string {
	return fmt.Sprintf("VoiceVerifyCodeSenderResult: Result=%d, ErrMsg=%s, Ext=%s, CallId=%s", v.Result, v.ErrMsg, v.Ext, v.CallId)
}

func (v *VoiceVerifyCodeSenderResult) ParseFromHTTPResponseBody(body []byte) error {
	err := json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}
