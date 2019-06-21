package sms

import (
	"encoding/json"
	"fmt"
)

type VoiceFileUploaderResult struct {
	Result int
	ErrMsg string
	Fid    string
}

func (v *VoiceFileUploaderResult) String() string {
	return fmt.Sprintf("VoiceFileUploaderResult: Result=%d, ErrMsg=%s, Fid=%s", v.Result, v.ErrMsg, v.Fid)
}

func (v *VoiceFileUploaderResult) ParseFromHTTPResponseBody(body []byte) error {
	err := json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}
