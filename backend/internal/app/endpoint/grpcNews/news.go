package grpcNews

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	pb "moscowhack/gen/go/news"
	"strconv"
)

type News interface {
	GetNews() (*pb.NewsItem, error)
	GetNewsById(id int) (*pb.NewsItem, error)
}

type Endpoint struct {
<<<<<<< HEAD:backend/internal/app/endpoint/grpcNews/news.go
	s Service
=======
	getNews News
>>>>>>> 60922d964ce532a3c8a9fed716754e95e23ddd83:backend/internal/app/endpoint/news/news.go
	pb.UnimplementedNewsServiceServer
}

func Register(gRPCServer *grpc.Server, news News) {
	pb.RegisterNewsServiceServer(gRPCServer, &Endpoint{getNews: news})
}

func New(news News) *Endpoint {
	return &Endpoint{
		getNews: news,
	}
}

func (e *Endpoint) GetNews(ctx context.Context, req *pb.NewsRequest) (*pb.NewsResponse, error) {
	fmt.Println(e, " ААА ", e.getNews)
	newsData, err := e.getNews.GetNews()
	if err != nil {
		return &pb.NewsResponse{}, err
	}

	newsItem := map[string]*pb.NewsItem{"NewsItem": newsData}

	return &pb.NewsResponse{News: newsItem}, nil
}

func (e *Endpoint) GetNewsById(ctx context.Context, req *pb.NewsRequest) (*pb.NewsResponse, error) {
	if req.Id == "" {
		return nil, errors.New("id новости не указан")
	}

	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return &pb.NewsResponse{}, err
	}

	newsData, err := e.getNews.GetNewsById(id)
	if err != nil {
		return &pb.NewsResponse{}, err
	}

	newsItem := map[string]*pb.NewsItem{"NewsItem": newsData}

	return &pb.NewsResponse{News: newsItem}, nil
}
