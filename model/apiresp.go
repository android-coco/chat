package model

type ApiResp struct {
	ErrorNo  int64 `json:"errno"`
	ErrorMsg string `json:"errmsg"`
	Data     interface{}  `json:"data,omitempty"`
}
