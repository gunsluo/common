#公共代码 
[![Build Status](https://travis-ci.org/gunsluo/common.svg?branch=master)](https://travis-ci.org/gunsluo/common) [![Coverage Status](https://coveralls.io/repos/github/gunsluo/common/badge.svg?branch=master)](https://coveralls.io/github/gunsluo/common?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/gunsluo/common)](https://goreportcard.com/report/github.com/gunsluo/common)

##RPC

###客户端

```Go
package main

import (
	"fmt"

	"github.com/gunsluo/common/example/rpc/model"
	"github.com/gunsluo/common/rpc"
)

func main() {

	client := rpc.NewClient("127.0.0.1:9999", 10)

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
```

###服务端

```Go
package main

import (
	"fmt"

	"github.com/gunsluo/common/example/rpc/model"
	"github.com/gunsluo/common/rpc"
)

// Handle define handle
type Handle int

// Test  test is example
func (t *Handle) Test(args *model.RequestArg, reply *model.ResponseArg) error {

	fmt.Println("receive:", args)

	reply.RespArgOne = "response"
	reply.Code = 200
	reply.Msg = "Success"

	fmt.Println("response:", reply)

	return nil
}

func main() {

	server := rpc.NewServer("0.0.0.0:9999")
	server.Register(new(Handle))

	server.Run()
}
```
