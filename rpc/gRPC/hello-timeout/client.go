package main

import (
	"context"
	"fmt"
	"github.com/lonchura/hello-timeout/pb"
	"google.golang.org/grpc"
	"log"
)

/*
 * protoc --go_out=plugins=grpc:. pb/hello.proto
 */
func main()  {
	conn, err := grpc.Dial(
		//"192.168.1.9:1234",
		"localhost:1234",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &pb.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply.GetValue())
}