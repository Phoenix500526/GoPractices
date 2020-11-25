package rpcutility

import (
	"net"
)

const HelloServiceName = "path/to/pkg.HelloService"

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

type HelloService struct{
	conn net.Conn
}

func NewHelloService(conn *net.Conn) *HelloService{
	return &HelloService{conn:*conn}
}


func (h *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello : " + request + ", from " + h.conn.RemoteAddr().String()
	return nil
}