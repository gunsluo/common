package rpc

import (
	"log"
	"math"
	"net/rpc"
	"sync"
	"time"
)

// Client define rpc client
type Client struct {
	sync.Mutex
	rpcClient *rpc.Client
	rpcServer string
	Timeout   time.Duration
}

// NewClient new rpc client
func NewClient(rpcServer string, timeout int) *Client {
	client := new(Client)
	client.rpcServer = rpcServer
	client.Timeout = time.Duration(timeout) * time.Second

	return client
}

// Close rpc client close
func (client *Client) Close() {
	if client.rpcClient != nil {
		client.rpcClient.Close()
		client.rpcClient = nil
	}
}

func (client *Client) insureConn() {
	if client.rpcClient != nil {
		return
	}

	var err error
	retry := 1

	for {
		if client.rpcClient != nil {
			return
		}

		client.rpcClient, err = NewJSONClient("tcp", client.rpcServer, client.Timeout)
		if err == nil {
			return
		}

		log.Printf("dial %s fail: %v", client.rpcServer, err)

		if retry > 6 {
			retry = 1
		}

		time.Sleep(time.Duration(math.Pow(2.0, float64(retry))) * time.Second)

		retry++
	}
}

// Call rpc client send msg
func (client *Client) Call(method string, args interface{}, reply interface{}) error {

	client.Lock()
	defer client.Unlock()

	client.insureConn()

	timeout := time.Duration(50 * time.Second)
	done := make(chan error)

	go func() {
		err := client.rpcClient.Call(method, args, reply)
		done <- err
	}()

	select {
	case <-time.After(timeout):
		log.Printf("[WARN] rpc call timeout %v => %v", client.rpcClient, client.rpcServer)
		client.Close()
	case err := <-done:
		if err != nil {
			client.Close()
			return err
		}
	}

	return nil
}
