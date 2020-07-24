package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"github.com/lonchura/hello-tls/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *pb.String) (*pb.String, error) {
	reply := &pb.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

// 拦截器
func filter(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	// info
	log.Println("the full RPC method string:", info.FullMethod)
	log.Printf("info.Server type: %T\n", info.Server)
	if server, ok := info.Server.(*HelloServiceImpl); ok {
		log.Println(server)
	}

	// req
	log.Printf("req type: %T\n", req)
	if msg, ok := req.(*pb.String); ok {
		log.Println(msg.GetValue(), ok)
		if msg.GetValue() == "hello" {
			return nil, errors.New("hello failed")
		}
	}

	return handler(ctx, req)
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
	// server crt
	certificate, err := tls.LoadX509KeyPair("tls/server-ca.crt", "tls/server.key")
	if err != nil {
		log.Fatal(err)
	}

	// ca
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("tls/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append certs")
	}

	// creds
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert, // NOTE: this is optional!
		ClientCAs:    certPool,
	})

	// grpc server
	grpcServer := grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(filter),
	)
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