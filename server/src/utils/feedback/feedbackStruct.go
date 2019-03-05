package feedback

import "net/http"

type FeedBack struct {
	DistWriter http.ResponseWriter `json:"-"`
	Code int `json:"code"`
	Msg string `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Total int `json:"total,omitempty"`
}