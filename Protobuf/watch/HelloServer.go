package main

import (
	"log"
	"net"
	"net/rpc"
	_ "watch/rpcutility"
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
