package main

import (
	"fmt"
	"log"
	"net"
	pb "go/gokit/protocol"

	grpc_transport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":5001"
)

// 使用go-kit包装grpc服务
type server struct{
	sayHelloHandle grpc_transport.Handler
	buyProHandle grpc_transport.Handler
}



func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	_, rsp, err := s.sayHelloHandle.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}

	// 通过go-kit包装
	return rsp.(*pb.HelloReply), err
	//return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) BuyPro(ctx context.Context, in *pb.BuyRequest) (*pb.BuyReply, error) {
	_, rsp, err := s.buyProHandle.ServeGRPC(ctx, in)
	if err != nil {
		return nil, err
	}

	return rsp.(*pb.BuyReply), err
	//return &pb.BuyReply{Message: in.Name, Price: 10}, nil
}

// 创建 sayHello 的 EndPoint
func makeSayHelloEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return &pb.HelloReply{Message: "Hello " + pb.HelloRequest{}.Name}, nil
	}
}
// 创建 buyPro 的 EndPoint
func makeBuyProEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return &pb.BuyReply{Message: pb.BuyRequest{}.Name, Price: 10}, nil
	}
}


func decodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeResponse(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}



func main() {
	newServer := new(server)
	sayHelloHandle := grpc_transport.NewServer(
		makeSayHelloEndpoint(),
		decodeRequest,
		encodeResponse,
		)
	newServer.sayHelloHandle = sayHelloHandle

	buyProHandle := grpc_transport.NewServer(
		makeBuyProEndpoint(),
		decodeRequest,
		encodeResponse,
		)
	newServer.buyProHandle = buyProHandle


	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, newServer)
	pb.RegisterBuyServerServer(s, newServer)
	s.Serve(lis)
}