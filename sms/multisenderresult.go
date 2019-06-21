package sms

import (
	"encoding/json"
	"fmt"
)

type Detail struct {
	Result     int
	ErrMsg     string
	Mobile     string
	NationCode string
	Sid        string
	Fee        int
}

func (d *Detail) String() string {
	return fmt.Sprintf("Detail: Result=%d, ErrMsg=%s, Mobile=%s, NationCode=%s, Sid=%s, Fee=%d", d.Result, d.ErrMsg, d.Mobile, d.NationCode, d.Sid, d.Fee)
}

type MultiSenderResult struct {
	Result int
	ErrMsg string
	Ext    string
	Detail []*Detail
}

func (m *MultiSenderResult) String() string {
	details := ""
	if m.Detail != nil {
		for i, detail := range m.Detail {
			details += detail.String()
			if i < len(m.Detail)-1 {
				details += ", "
			}
		}
	}

	return fmt.Sprintf("MultiSenderResult: Result=%d, ErrMsg=%s, Ext=%s, Detail=%s", m.Result, m.ErrMsg, m.Ext, details)
}

func (m *MultiSenderResult) ParseFromHTTPResponseBody(body []byte) error {
	err := json.Unmarshal(body, m)
	if err != nil {
		return err
	}

	return nil
}
