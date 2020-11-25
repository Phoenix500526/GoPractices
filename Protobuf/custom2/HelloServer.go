package main

import (
	"custom2/rpcutility"
	"log"
	"net"
	"net/rpc"
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
		go func() {
			defer conn.Close()
			p := rpc.NewServer()
			p.RegisterName(rpcutility.HelloServiceName, rpcutility.NewHelloService(&conn))
			p.ServeConn(conn)
		}()
	}

}
