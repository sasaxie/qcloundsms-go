package sms

import (
	"encoding/json"
	"fmt"
)

type FileVoiceSenderResult struct {
	Result int
	ErrMsg string
	Ext    string
	CallId string
}

func (f *FileVoiceSenderResult) String() string {
	return fmt.Sprintf("FileVoiceSenderResult: Result=%d, ErrMsg=%s, Ext=%s, CallId=%s", f.Result, f.ErrMsg, f.Ext, f.CallId)
}

func (f *FileVoiceSenderResult) ParseFromHTTPResponseBody(body []byte) error {
	err := json.Unmarshal(body, f)
	if err != nil {
		return err
	}

	return nil
}
