package sms

import (
	"encoding/json"
	"fmt"
)

type Reply struct {
	NationCode string `json:"nationcode"`
	Mobile     string `json:"mobile"`
	Text       string `json:"text"`
	Sign       string `json:"sign"`
	Time       int64  `json:"time"`
	Extend     string `json:"extend"`
}

func (r *Reply) String() string {
	return fmt.Sprintf("Reply: NationCode=%s, Mobile=%s, Text=%s, Sign=%s, Time=%d, Extend=%s", r.NationCode, r.Mobile, r.Text, r.Sign, r.Time, r.Extend)
}

type StatusPullReplyResult struct {
	Result int
	ErrMsg string
	Count  int
	Reply  []*Reply `json:"data"`
}

func (s *StatusPullReplyResult) String() string {
	reply := ""
	if s.Reply != nil {
		for i, r := range s.Reply {
			reply += r.String()
			if i < len(s.Reply)-1 {
				reply += ", "
			}
		}
	}

	return fmt.Sprintf("StatusPullReplyResult: Result=%d, ErrMsg=%s, Count=%d, Reply=%s", s.Result, s.ErrMsg, s.Count, reply)
}

func (s *StatusPullReplyResult) ParseFromHTTPResponseBody(body []byte) error {
	err := json.Unmarshal(body, s)
	if err != nil {
		return err
	}

	return nil
}
