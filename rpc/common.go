package rpc

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

// NewRPCClient new client
func NewRPCClient(network, address string, timeout time.Duration) (*rpc.Client, error) {
	conn, err := net.DialTimeout(network, address, timeout)
	if err != nil {
		return nil, err
	}
	return rpc.NewClient(conn), nil
}

// NewJSONClient new jsonclient
func NewJSONClient(network, address string, timeout time.Duration) (*rpc.Client, error) {
	conn, err := net.DialTimeout(network, address, timeout)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), err
}
