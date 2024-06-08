package grpcNews

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	pb "moscowhack/gen/go/news"
	"strconv"
	"strings"
)

type News interface {
	GetNewsService(ctx context.Context) (*pb.NewsItem, error)
	GetNewsByIdService(ctx context.Context, id int) (*pb.NewsItem, error)
	GetNewsByCategoryService(ctx context.Context, categoryId []string) (*pb.NewsItem, error)
}

type Endpoint struct {
	News News
	pb.UnimplementedNewsServiceServer
}

func Register(gRPCServer *grpc.Server, news News) {
	pb.RegisterNewsServiceServer(gRPCServer, &Endpoint{News: news})
}

func New(news News) *Endpoint {
	return &Endpoint{
		News: news,
	}
}

func (e *Endpoint) GetNews(ctx context.Context, req *pb.NewsRequest) (*pb.NewsResponse, error) {
	newsData, err := e.News.GetNewsService(ctx)
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

	newsData, err := e.News.GetNewsByIdService(ctx, id)
	if err != nil {
		return &pb.NewsResponse{}, err
	}

	newsItem := map[string]*pb.NewsItem{"NewsItem": newsData}

	return &pb.NewsResponse{News: newsItem}, nil
}

func (e *Endpoint) GetNewsByCategory(ctx context.Context, req *pb.NewsRequest) (*pb.NewsResponse, error) {
	if req.Category == "" {
		return nil, errors.New("id категории не указан")
	}

	categorySlice := strings.Split(req.Category, ",")

	newsData, err := e.News.GetNewsByCategoryService(ctx, categorySlice)
	if err != nil {
		return &pb.NewsResponse{}, err
	}

	newsItem := map[string]*pb.NewsItem{"NewsItem": newsData}

	return &pb.NewsResponse{News: newsItem}, nil
}
