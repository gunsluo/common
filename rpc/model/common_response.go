package model

import "fmt"

// CommonRPCResponse define common RPC Response
type CommonRPCResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// String common RPC Response to string
func (resp *CommonRPCResponse) String() string {
	return fmt.Sprintf("<Code: %d, Msg: %s>", resp.Code, resp.Msg)
}
