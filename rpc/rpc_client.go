package rpc

import (
	"log"
	"math"
	"net/rpc"
	"sync"
	"time"
)

type RPCClient struct {
	sync.Mutex
	rpcClient *rpc.Client
	RpcServer string
	Timeout   time.Duration
}

func NewRPCClient(rpcServer string, timeout int) *RPCClient {
	client := new(RPCClient)
	client.RpcServer = rpcServer
	client.Timeout = time.Duration(timeout) * time.Second

	return client
}

func (this *RPCClient) Close() {
	if this.rpcClient != nil {
		this.rpcClient.Close()
		this.rpcClient = nil
	}
}

func (this *RPCClient) insureConn() {
	if this.rpcClient != nil {
		return
	}

	var err error
	var retry int = 1

	for {
		if this.rpcClient != nil {
			return
		}

		this.rpcClient, err = NewJsonClient("tcp", this.RpcServer, this.Timeout)
		if err == nil {
			return
		}

		log.Printf("dial %s fail: %v", this.RpcServer, err)

		if retry > 6 {
			retry = 1
		}

		time.Sleep(time.Duration(math.Pow(2.0, float64(retry))) * time.Second)

		retry++
	}
}

func (this *RPCClient) Call(method string, args interface{}, reply interface{}) error {

	this.Lock()
	defer this.Unlock()

	this.insureConn()

	timeout := time.Duration(50 * time.Second)
	done := make(chan error)

	go func() {
		err := this.rpcClient.Call(method, args, reply)
		done <- err
	}()

	select {
	case <-time.After(timeout):
		log.Printf("[WARN] rpc call timeout %v => %v", this.rpcClient, this.RpcServer)
		this.Close()
	case err := <-done:
		if err != nil {
			this.Close()
			return err
		}
	}

	return nil
}
