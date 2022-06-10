package model

// 统一响应

type CustomResp struct {
	Code int `json:"code"`

	Msg string `json:"msg"`

	Data interface{} `json:"data"`
}
