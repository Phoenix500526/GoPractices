package main

import (
	"custom2/rpcutility"
	"fmt"
	"log"
	"net/rpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(rpcutility.HelloServiceName+".Hello", request, reply)
}

func main() {
	client, err := DialHelloService("tcp", "localhost:2020")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello(" World", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
