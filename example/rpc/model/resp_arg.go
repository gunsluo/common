package model

import (
	"fmt"
)

type ResponseArg struct {
	RespArgOne string `json:"resp_arg_one"`
	CommonRPCResponse
}

func (this *ResponseArg) String() string {
	return fmt.Sprintf("<Code: %d, Msg: %s, ResponseArg: %s>", this.Code, this.Msg, this.RespArgOne)
}
