package sms

import (
	"encoding/json"
	"fmt"
)

type VoiceFileStatusResult struct {
	Result int
	ErrMsg string
	Status int
}

func (v *VoiceFileStatusResult) String() string {
	return fmt.Sprintf("VoiceFileStatusResult: Result=%d, ErrMsg=%s, Status=%d", v.Result, v.ErrMsg, v.Status)
}

func (v *VoiceFileStatusResult) ParseFromHTTPResponseBody(body []byte) error {
	err := json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}
