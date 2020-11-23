package main

import (
	"log"
	"net/rpc"
	"net"
)

const HelloServiceName = "path/to/pkg.HelloService"


type HelloServiceInterface interface{
	Hello(request string, reply *string) error
}

func RegisterHelloServer(svc HelloServiceInterface) error{
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloServer struct{}

func (h *HelloServer) Hello(request string, reply *string) error{
	*reply = "Hello : " + request
	return nil
}

func main(){
	RegisterHelloServer(new(HelloServer))
	listener, err := net.Listen("tcp", ":2020")
	if err != nil{
		log.Fatal("ListenTCP error:", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}
	rpc.ServeConn(conn)
}