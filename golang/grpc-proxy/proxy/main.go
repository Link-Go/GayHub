package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/mwitkow/grpc-proxy/proxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

var (
	director proxy.StreamDirector
)

func main() {
	director = func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		fmt.Println(fullMethodName) // /helloworld.Greeter/SayHello
		// 取出 ctx 中的信息，可以进行鉴权、传递信息等操作
		md, ok := metadata.FromIncomingContext(ctx)
		fmt.Println(md, ok) // map[:authority:[localhost:50051] content-type:[application/grpc] key1:[value1] key2:[value2] user-agent:[grpc-go/1.54.0]] true

		// 可以建立一张 fullMethodName 与 target 的配置映射表
		// 创建不同的 clientConn
		target := "localhost:50052"
		outCtx, _ := context.WithCancel(ctx)
		outCtx = metadata.NewOutgoingContext(outCtx, md.Copy())
		conn, err := grpc.DialContext(ctx, target, grpc.WithTransportCredentials(insecure.NewCredentials()))
		return outCtx, conn, err
	}

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	unknown := grpc.UnknownServiceHandler(proxy.TransparentHandler(director))
	s := grpc.NewServer(unknown)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
