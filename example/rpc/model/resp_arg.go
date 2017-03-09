package model

import (
	"fmt"
)

// ResponseArg response arg example
type ResponseArg struct {
	RespArgOne string `json:"resp_arg_one"`
	CommonRPCResponse
}

// String response arg to string
func (resp *ResponseArg) String() string {
	return fmt.Sprintf("<Code: %d, Msg: %s, ResponseArg: %s>", resp.Code, resp.Msg, resp.RespArgOne)
}
