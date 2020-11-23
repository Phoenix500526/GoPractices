package main

import (
	"AdvancedGO/RPC/better_rpc/rpcutility"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type DividedServiceClient struct {
	*rpc.Client
}

func DialDividedService(network, address string) (*DividedServiceClient, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &DividedServiceClient{Client: rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))}, nil
}

func (p *DividedServiceClient) Divide(num rpcutility.Number, result *rpcutility.Number) error {
	return p.Client.Call(rpcutility.DividedServiceName+".Divide", num, result)
}

func main() {
	client, err := DialDividedService("tcp", "localhost:2020")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	num := rpcutility.Number{10, 0}
	var result rpcutility.Number
	err = client.Divide(num, &result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("quo = %d, rem = %d\n", result.A, result.B)
}
