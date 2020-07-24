package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/lonchura/hello-tls/pb"
	"github.com/grpc-ecosystem/go-grpc-middleware"
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

// 拦截器(tracing上报)
func tracingFilter(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	log.Println("Tracing filter")

	return handler(ctx, req)
}
// 拦截器(监控日志上报)
func prometheusFilter(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	log.Println("Prometheus filter")

	return handler(ctx, req)
}
// 拦截器(日志)
func logFilter(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	log.Println("Log filter")

	return handler(ctx, req)
}
// 拦截器(验签)
func signerFilter(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	log.Println("Signer filter")

	return handler(ctx, req)
}
// 拦截器(鉴权)
func authFilter(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	log.Println("Log filter")

	return handler(ctx, req)
}
// 拦截器(恢复)
func recoveryFilter(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	log.Println("Recovery filter")

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
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			tracingFilter, prometheusFilter, logFilter, signerFilter, authFilter, recoveryFilter,
		)),
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