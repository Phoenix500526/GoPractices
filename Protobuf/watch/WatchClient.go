package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
	"watch/rpcutility"
)

type WatchServiceClient struct {
	*rpc.Client
}

func DialWatchService(network, address string) (*WatchServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &WatchServiceClient{Client: c}, nil
}

func (p *WatchServiceClient) Watch(timeout int, keyChanged *string) error {
	return p.Client.Call(rpcutility.WatchServiceName+".Watch", timeout, keyChanged)
}

func (p *WatchServiceClient) Set(key string, value string) error {
	return p.Client.Call(rpcutility.WatchServiceName+".Set", [2]string{key, value}, new(struct{}))
}

func main() {
	client, err := DialWatchService("tcp", "localhost:2020")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	go func() {
		var keyChanged string
		err := client.Watch(30, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("watch: ", keyChanged)
	}()

	err = client.Set("abc", "abc-value")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 3)
}
