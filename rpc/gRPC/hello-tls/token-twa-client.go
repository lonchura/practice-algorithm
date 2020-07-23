package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/lonchura/hello-tls/pb"
	"github.com/lonchura/hello-tls/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

/*
# 证书签名请求文件(CN=client.io, Common Name)
$ openssl req -new \
    -subj "/C=GB/L=China/O=client/CN=client.io" \
    -key client.key \
    -out client.csr
# CA签名证书
$ openssl x509 -req -sha256 \
    -CA ca.crt -CAkey ca.key -CAcreateserial -days 3650 \
    -in client.csr \
    -out client-ca.crt
 */
func main()  {
	// auth token
	auth := token.Authentication{
		User: "lonchura",
		Password: "1234567",
	}

	// client crt
	certificate, err := tls.LoadX509KeyPair("tls/client-ca.crt", "tls/client.key")
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

	conn, err := grpc.Dial(
		"localhost:1234",
		//"192.168.1.9:1234",
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(&auth),
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