package model

import "fmt"

type CommonRPCResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (this *CommonRPCResponse) String() string {
	return fmt.Sprintf("<Code: %d, Msg: %s>", this.Code, this.Msg)
}
