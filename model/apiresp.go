package model

type ApiResp struct {
	ErrorNo  int64       `json:"code"`
	ErrorMsg string      `json:"msg"`
	Data     interface{} `json:"data,omitempty"`
	Rows     interface{} `json:"rows,omitempty"`
	Total    interface{} `json:"total,omitempty"`
}
