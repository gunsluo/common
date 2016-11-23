package main

import (
	"fmt"

	"github.com/gunsluo/common/example/rpc/model"
	"github.com/gunsluo/common/rpc"
)

func main() {

	client := rpc.NewRPCClient("127.0.0.1:9999", 10)

	req := new(model.RequestArg)
	req.ArgOne = "test"
	var resp model.ResponseArg
	fmt.Printf("send: %+v\n", req)
	err := client.Call("Handle.Test", req, &resp)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Printf("resp: %+v\n", resp)
	}
}
