package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {}

// Hello方法必须满足Go语言的RPC规则：
// 	- 方法只能有两个可序列化的参数
//	- 其中第二个参数是指针类型
//	- 并且返回一个error类型
//	- 同时必须是公开的方法
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main()  {
	err := rpc.RegisterName("HelloService", new(HelloService))
	if err != nil {
		log.Fatal("rpc Register failed error:", err)
	}

	// test: nc -l 1234
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}