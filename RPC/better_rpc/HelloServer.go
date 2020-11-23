package main

import (
	_"AdvancedGO/RPC/better_rpc/rpcutility"
	"net"
	"net/rpc"
	"log"
)

func main() {
	listener, err := net.Listen("tcp", ":2020")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}

}
