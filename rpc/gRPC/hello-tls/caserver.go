package main

import (
	"context"
	"github.com/lonchura/hello-tls/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *pb.String) (*pb.String, error) {
	reply := &pb.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}
/*
# CA根证书
$ openssl genrsa -out ca.key 2048
$ openssl req -new -x509 -days 3650 \
    -subj "/C=GB/L=China/O=gobook/CN=github.com" \
    -key ca.key -out ca.crt

# 证书签名请求文件(CN=server.io, Common Name)
$ openssl req -new \
    -subj "/C=GB/L=China/O=server/CN=server.io" \
    -key server.key \
    -out server.csr
# CA签名证书
$ openssl x509 -req -sha256 \
    -CA ca.crt -CAkey ca.key -CAcreateserial -days 3650 \
    -in server.csr \
    -out server.crt
 */
func main()  {
	// creds
	creds, err := credentials.NewServerTLSFromFile("tls/server-ca.crt", "tls/server.key")
	if err != nil {
		log.Fatal(err)
	}

	// grpc server
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	// register hello service
	pb.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	// listen
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	// serve
	grpcServer.Serve(lis)
}