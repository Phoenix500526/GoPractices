package main

import (
	"net/rpc"
	"log"
	"fmt"
)

const HelloServiceName = "path/to/pkg.HelloService"

func main(){
	client, err := rpc.Dial("tcp", "localhost:2020")
	if err != nil{
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Call(HelloServiceName + ".Hello", "RPC World", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}