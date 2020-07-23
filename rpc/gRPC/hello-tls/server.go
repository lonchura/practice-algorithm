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
# 生成私钥
$ openssl genrsa -out server.key 2048
# 通过私钥在本地申请公钥证书
$ openssl req -new -x509 -days 3650 \
    -subj "/C=GB/L=China/O=grpc-server/CN=server.grpc.io" \
    -key server.key -out server.crt
 */
func main()  {
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