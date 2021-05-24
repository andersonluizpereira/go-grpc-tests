package main

import (
	"context"
	"fmt"
	"github.com/acpereira/grpc-ads/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	connection, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer connection.Close()

	client := pb.NewBookServiceClient(connection)
	AddBook(client)
}

func AddBook(client pb.BookServiceClient) {
	req := &pb.Book{
		Id:    "1",
		Description:  "Xuxucao DVD para baixainhos",
		Title: "Xuxucao Dvd",
	}
	res, err := client.AddBook(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}
	fmt.Println(res)
}
