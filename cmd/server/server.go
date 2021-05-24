package main

import (
	"github.com/acpereira/grpc-ads/pb"
	"github.com/acpereira/grpc-ads/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main()  {
	lis, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal("Could not connect %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterBookServiceServer(grpcServer, &services.BookService{})
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Could not serve %v", err)
	}
}