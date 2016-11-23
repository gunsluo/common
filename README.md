#公共代码

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
```

###服务端

```Go
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
```
