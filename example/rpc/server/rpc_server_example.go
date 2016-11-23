package main

import (
	"fmt"

	"github.com/gunsluo/common/example/rpc/model"
	"github.com/gunsluo/common/rpc"
)

type Handle int

func (t *Handle) Test(args *model.RequestArg, reply *model.ResponseArg) error {

	fmt.Println("receive:", args)

	reply.RespArgOne = "response"
	reply.Code = 200
	reply.Msg = "Success"

	fmt.Println("respone:", reply)

	return nil
}

func main() {

	server := rpc.NewRPCServer("0.0.0.0:9999")
	server.Register(new(Handle))

	server.Run()
}
