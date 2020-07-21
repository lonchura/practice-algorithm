package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main()  {
	client, err := rpc.Dial("tcp", "192.168.1.9:1234")
	if err != nil {
		log.Fatal("Dial failed:", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal("Call failed:", err)
	}

	fmt.Println(reply)
}