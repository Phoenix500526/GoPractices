package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main(){
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil{
		log.Fatal("dialing:", err)
	}
	var reply String
	request := String{Value:"RPC World"}
	err = client.Call("HelloService.Hello", &request, &reply)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("%v\n",reply.GetValue())
}