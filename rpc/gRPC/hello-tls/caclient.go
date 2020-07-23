package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/lonchura/hello-tls/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

/*
$ openssl genrsa -out client.key 2048
$ openssl req -new -x509 -days 3650 \
    -subj "/C=GB/L=China/O=grpc-client/CN=client.grpc.io" \
    -key client.key -out client.crt
 */
func main()  {
	// client crt
	certificate, err := tls.LoadX509KeyPair("tls/client.crt", "tls/client.key")
	if err != nil {
		log.Fatal(err)
	}

	// root ca
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("tls/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs")
	}

	// creds
	creds := credentials.NewTLS(&tls.Config{
		Certificates:       []tls.Certificate{certificate},
		ServerName:         "server.io", // NOTE: this is required!
		RootCAs:            certPool,
	})

	//conn, err := grpc.Dial("192.168.1.9:1234")
	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &pb.String{Value: "hello"})
	if err != nil {
		fmt.Println(123)
		log.Fatal(err)
	}

	fmt.Println(reply.GetValue())
}