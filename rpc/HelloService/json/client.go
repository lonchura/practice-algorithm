package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main()  {
	conn, err := net.Dial("tcp", "192.168.1.9:1234")
	if err != nil {
		log.Fatal("Dial failed:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)

	if err != nil {
		log.Fatal("Call failed:", err)
	}

	fmt.Println(reply)
}