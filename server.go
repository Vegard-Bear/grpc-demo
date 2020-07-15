package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/grpc-demo/proto"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

func main() {

	//创建 gRPC Server 对象
	server := grpc.NewServer()

	//将 SearchService（其包含需要被调用的服务端接口）注册到 gRPC Server 的内部注册中心
	pb.RegisterSearchServiceServer(server, &SearchService{})

	//创建 Listen，监听 TCP 端口
	lis, err := net.Listen("tcp", ":"+"9001")
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	//gRPC Server 开始 lis.Accept，直到 Stop 或 GracefulStop
	server.Serve(lis)
}
