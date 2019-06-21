package sms

import (
	"encoding/json"
	"fmt"
)

type TtsVoiceSenderResult struct {
	Result int
	ErrMsg string
	Ext    string
	CallId string
}

func (t *TtsVoiceSenderResult) String() string {
	return fmt.Sprintf("TtsVoiceSenderResult: Result=%d, ErrMsg=%s, Ext=%s, CallId=%s", t.Result, t.ErrMsg, t.Ext, t.CallId)
}

func (t *TtsVoiceSenderResult) ParseFromHTTPResponseBody(body []byte) error {
	err := json.Unmarshal(body, t)
	if err != nil {
		return err
	}

	return nil
}
