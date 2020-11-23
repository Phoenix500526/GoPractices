package rpcutility

import (
	"net/rpc"
	"errors"
)

const HelloServiceName = "path/to/pkg.HelloService"
const DividedServiceName = "path/to/pkg.DividedService"

func init(){
	RegisterHelloServer(new(HelloServer))
	RegisterDividedServer(new(DividedServer))
}

type Number struct{
	A int32
	B int32
}

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

type DividedServiceInterface interface{
	Divide(num Number, result *Number) error
}

func RegisterHelloServer(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

func RegisterDividedServer(svc DividedServiceInterface) error {
	return rpc.RegisterName(DividedServiceName, svc)
}

type HelloServer struct{}

type DividedServer struct{}

func (h *HelloServer) Hello(request string, reply *string) error {
	*reply = "Hello : " + request
	return nil
}

func (d *DividedServer) Divide(num Number, result *Number) error {
	if num.B == 0 {
		return errors.New("Invalid Divisor")
	}
	result.A = num.A / num.B
	result.B = num.A % num.B
	return nil
}