package model

import (
	"fmt"

	"git.oschina.net/wormholelogic/lib/rpc/model"
)

type ResponseArg struct {
	RespArgOne string `json:"resp_arg_one"`
	model.CommonRPCResponse
}

func (this *ResponseArg) String() string {
	return fmt.Sprintf("<Code: %d, Msg: %s, ResponseArg: %s>", this.Code, this.Msg, this.RespArgOne)
}
