package main

import (
	"context"
	"fmt"
	"github.com/lonchura/hello-tls/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

/*
$ openssl genrsa -out client.key 2048
$ openssl req -new -x509 -days 3650 \
    -subj "/C=GB/L=China/O=grpc-client/CN=client.grpc.io" \
    -key client.key -out client.crt
*/
func main()  {
	// creds with server crt
	creds, err := credentials.NewClientTLSFromFile(
		"tls/server.crt", "server.grpc.io",
	)
	if err != nil {
		log.Fatal(err)
	}

	//conn, err := grpc.Dial("192.168.1.9:1234")
	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(creds))
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