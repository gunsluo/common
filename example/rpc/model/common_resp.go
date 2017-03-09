package model

import "fmt"

// CommonRPCResponse RPC Response
type CommonRPCResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// String RPC Response to string
func (resp *CommonRPCResponse) String() string {
	return fmt.Sprintf("<Code: %d, Msg: %s>", resp.Code, resp.Msg)
}
