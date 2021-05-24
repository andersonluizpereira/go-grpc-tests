package services

import (
	"context"
	"fmt"
	"github.com/acpereira/grpc-ads/pb"
)

type BookService struct {
	pb.UnimplementedBookServiceServer
}

func NewBookService() *BookService{
	return &BookService{}
}

func (*BookService) AddBook(ctx context.Context, req *pb.Book) (*pb.Book, error) {
	fmt.Println(req.Description)
	return &pb.Book{
		Id:    "123",
		Description:  req.GetDescription(),
		Title: req.GetTitle(),
	}, nil
}
