package rpc

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Server rpc server
type Server struct {
	server *rpc.Server
	Listen string
}

// NewServer new rpc server
func NewServer(listen string) *Server {
	rpcServer := new(Server)
	rpcServer.Listen = listen
	rpcServer.server = rpc.NewServer()

	return rpcServer
}

// Run listen rpc server
func (server *Server) Run() {

	tcpAddr, err := net.ResolveTCPAddr("tcp", server.Listen)
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}
	log.Println("rpc listening", server.Listen)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("listener.Accept occur error:", err)
			continue
		}
		// go rpc.ServeConn(conn)
		go server.server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

// Register rpc server register recipient
func (server *Server) Register(rcvr interface{}) error {
	return server.server.Register(rcvr)
}
