package main

import (
	"fmt"
	"log"
	"net/rpc"
	"AdvancedGO/RPC/better_rpc/rpcutility"
)

type DividedServiceClient struct {
	*rpc.Client
}

func DialDividedService(network, address string) (*DividedServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &DividedServiceClient{Client:c}, nil
}

func (p *DividedServiceClient) Divide(num rpcutility.Number, result *rpcutility.Number) error {
	return p.Client.Call(rpcutility.DividedServiceName+".Divide", num, result)
}

func main() {
	client, err := DialDividedService("tcp", "localhost:2020")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	num := rpcutility.Number{10, 7}
	var result rpcutility.Number
	err = client.Divide(num, &result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("quo = %d, rem = %d\n", result.A, result.B)
}
