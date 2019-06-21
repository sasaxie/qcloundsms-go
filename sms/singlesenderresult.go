package sms

import (
	"encoding/json"
	"fmt"
)

type SingleSenderResult struct {
	Result int
	ErrMsg string
	Ext    string
	Sid    string
	Fee    int
}

func (s *SingleSenderResult) String() string {
	return fmt.Sprintf("SingleSenderResult: Result=%d, ErrMsg=%s, Ext=%s, Sid=%s, Fee=%d", s.Result, s.ErrMsg, s.Ext, s.Sid, s.Fee)
}

func (s *SingleSenderResult) ParseFromHTTPResponseBody(body []byte) error {
	err := json.Unmarshal(body, s)
	if err != nil {
		return err
	}

	return nil
}
