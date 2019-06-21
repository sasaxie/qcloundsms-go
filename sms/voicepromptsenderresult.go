package sms

import (
	"encoding/json"
	"fmt"
)

type VoicePromptSenderResult struct {
	Result int
	ErrMsg string
	Ext    string
	CallId string
}

func (v *VoicePromptSenderResult) String() string {
	return fmt.Sprintf("VoicePromptSenderResult: Result=%d, ErrMsg=%s, Ext=%s, CallId=%s", v.Result, v.ErrMsg, v.Ext, v.CallId)
}

func (v *VoicePromptSenderResult) ParseFromHTTPResponseBody(body []byte) error {
	err := json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}
