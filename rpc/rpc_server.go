package rpc

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type RPCServer struct {
	server *rpc.Server
	Listen string
}

func NewRPCServer(listen string) *RPCServer {
	rpcServer := new(RPCServer)
	rpcServer.Listen = listen
	rpcServer.server = rpc.NewServer()

	return rpcServer
}

func (this *RPCServer) Run() {

	tcpAddr, err := net.ResolveTCPAddr("tcp", this.Listen)
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}
	log.Println("rpc listening", this.Listen)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("listener.Accept occur error:", err)
			continue
		}
		// go rpc.ServeConn(conn)
		go this.server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

func (this *RPCServer) Register(rcvr interface{}) error {
	return this.server.Register(rcvr)
}
