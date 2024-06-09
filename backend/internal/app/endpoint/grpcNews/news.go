package grpcNews

import (
	"context"
	"errors"
	pb "moscowhack/gen/go/news"
)

type News interface {
	GetNewsService(ctx context.Context) (map[string]*pb.NewsItem, error)
	GetNewsByIdService(ctx context.Context, id int32) (map[string]*pb.NewsItem, error)
	GetNewsByCategoryService(ctx context.Context, categoryId string) (map[string]*pb.NewsItem, error)
	AddNewsService(ctx context.Context, title string, text string, datetime string, categories string) (int32, error)
	DelNewsService(ctx context.Context, newsID int32) error
}

type Endpoint struct {
	News News
	pb.UnimplementedNewsServiceServer
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

	return &pb.NewsResponse{News: newsData}, nil
}

func (e *Endpoint) GetNewsById(ctx context.Context, req *pb.NewsRequest) (*pb.NewsResponse, error) {
	if req.Id == 0 {
		return nil, errors.New("id новости не указан")
	}

	newsData, err := e.News.GetNewsByIdService(ctx, req.Id)
	if err != nil {
		return &pb.NewsResponse{}, err
	}

	return &pb.NewsResponse{News: newsData}, nil
}

func (e *Endpoint) GetNewsByCategory(ctx context.Context, req *pb.NewsRequest) (*pb.NewsResponse, error) {
	if req.Categories == "" {
		return nil, errors.New("id категории не указан")
	}

	newsData, err := e.News.GetNewsByCategoryService(ctx, req.Categories)
	if err != nil {
		return &pb.NewsResponse{}, err
	}

	return &pb.NewsResponse{News: newsData}, nil
}

func (e *Endpoint) AddNews(ctx context.Context, req *pb.NewsRequest) (*pb.ChangeNewsResponse, error) {
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
		return &pb.ChangeNewsResponse{Err: error.Error(err)}, err
	}

	return &pb.ChangeNewsResponse{Id: id}, nil
}

func (e *Endpoint) DelNews(ctx context.Context, req *pb.NewsRequest) (*pb.ChangeNewsResponse, error) {
	if req.Id == 0 {
		return nil, errors.New("id новости не указан")
	}

	err := e.News.DelNewsService(ctx, req.Id)
	if err != nil {
		return &pb.ChangeNewsResponse{Err: error.Error(err)}, err
	}

	return &pb.ChangeNewsResponse{Err: ""}, nil
}
