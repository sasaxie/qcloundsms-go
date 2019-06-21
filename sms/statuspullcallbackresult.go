package sms

import (
	"encoding/json"
	"fmt"
)

type Callback struct {
	UserReceiveTime string `json:"user_receive_time"`
	NationCode      string `json:"nationcode"`
	Mobile          string `json:"mobile"`
	ReportStatus    string `json:"report_status"`
	ErrMsg          string `json:"errmsg"`
	Description     string `json:"description"`
	Sid             string `json:"sid"`
}

func (c *Callback) String() string {
	return fmt.Sprintf("Callback: UserReceiveTime=%s, NationCode=%s, Mobile=%s, ReportStatus=%s, ErrMsg=%s, Description=%s, Sid=%s", c.UserReceiveTime, c.NationCode, c.Mobile, c.ReportStatus, c.ErrMsg, c.Description, c.Sid)
}

type StatusPullCallbackResult struct {
	Result    int
	ErrMsg    string
	Count     int
	Callbacks []*Callback `json:"data"`
}

func (s *StatusPullCallbackResult) String() string {
	callbacks := ""
	if s.Callbacks != nil {
		for i, callback := range s.Callbacks {
			callbacks += callback.String()
			if i < len(s.Callbacks)-1 {
				callbacks += ", "
			}
		}
	}

	return fmt.Sprintf("StatusPullCallbackResult: Result=%d, ErrMsg=%s, Count=%d, Callbacks=%s", s.Result, s.ErrMsg, s.Count, callbacks)
}

func (s *StatusPullCallbackResult) ParseFromHTTPResponseBody(body []byte) error {
	err := json.Unmarshal(body, s)
	if err != nil {
		return err
	}

	return nil
}
