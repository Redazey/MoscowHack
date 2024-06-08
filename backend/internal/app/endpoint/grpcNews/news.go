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
	GetNewsService(ctx context.Context) (*pb.NewsItem, error)
	GetNewsByIdService(ctx context.Context, id int) (*pb.NewsItem, error)
	GetNewsByCategoryService(ctx context.Context, categoryId string) (*pb.NewsItem, error)
	AddNewsService(ctx context.Context, title string, text string, datetime string, categories string) (int, error)
	DelNewsService(ctx context.Context, newsID int) error
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
	if req.Categories == "" {
		return nil, errors.New("id категории не указан")
	}

	newsData, err := e.News.GetNewsByCategoryService(ctx, req.Categories)
	if err != nil {
		return &pb.NewsResponse{}, err
	}

	newsItem := map[string]*pb.NewsItem{"NewsItem": newsData}

	return &pb.NewsResponse{News: newsItem}, nil
}

func (e *Endpoint) AddNews(ctx context.Context, req *pb.NewsRequest) (*pb.NewsResponse, error) {
	if req.Title == "" {
		return nil, errors.New("заголовок новости не указан")
	}
	if req.Text == "" {
		return nil, errors.New("текст новости не указан")
	}
	if req.Datetime == "" {
		return nil, errors.New("дата публикации новости не указан")
	}
	if req.Categories == "" {
		return nil, errors.New("id категорий новости не указан")
	}

	id, err := e.News.AddNewsService(ctx, req.Title, req.Text, req.Datetime, req.Categories)
	if err != nil {
		return &pb.NewsResponse{Err: error.Error(err)}, err
	}

	return &pb.NewsResponse{Id: fmt.Sprint(id)}, nil
}

func (e *Endpoint) DelNews(ctx context.Context, req *pb.NewsRequest) (*pb.NewsResponse, error) {
	if req.Id == "" {
		return nil, errors.New("id новости не указан")
	}

	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return &pb.NewsResponse{Err: error.Error(err)}, err
	}

	err = e.News.DelNewsService(ctx, id)
	if err != nil {
		return &pb.NewsResponse{Err: error.Error(err)}, err
	}

	return &pb.NewsResponse{Err: ""}, nil
}
