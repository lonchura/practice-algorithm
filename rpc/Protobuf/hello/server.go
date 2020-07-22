package main

import (
	"github.com/lonchura/hello/pb"
)

type HelloService struct{}

func (p *HelloService) Hello(request *pb.String, reply *pb.String) error {
	reply.Value = "hello:" + request.GetValue()
	return nil
}

func main()  {
	
}