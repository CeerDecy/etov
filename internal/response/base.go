package response

import "net/http"

type BaseResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func NewBaseResp(code int, msg string, data any) *BaseResp {
	return &BaseResp{Code: code, Msg: msg, Data: data}
}

func SuccessResp(data any) *BaseResp {
	return NewBaseResp(http.StatusOK, "success", data)
}

func ErrorResp(err error) *BaseResp {
	return NewBaseResp(http.StatusInternalServerError, err.Error(), nil)
}

func ErrorMsgResp(msg string) *BaseResp {
	return NewBaseResp(http.StatusInternalServerError, msg, nil)
}
