package grpcNews

import (
	"context"
	"errors"
	pb "moscowhack/gen/go/news"
)

type News interface {
	GetNewsService() (map[string]*pb.GetNewsItem, error)
	GetNewsByIdService(int32) (*pb.GetNewsByIdResponse, error)
	GetNewsByCategoryService(string) (map[string]*pb.GetNewsItem, error)
	AddNewsService(string, string, string, string) (int32, error)
	DelNewsService(int32) error
}

type Endpoint struct {
	News News
	pb.UnimplementedNewsServiceServer
}

func (e *Endpoint) GetNews(_ context.Context, _ *pb.GetNewsRequest) (*pb.GetNewsResponse, error) {
	newsData, err := e.News.GetNewsService()
	if err != nil {
		return &pb.GetNewsResponse{}, err
	}

	return &pb.GetNewsResponse{News: newsData}, nil
}

func (e *Endpoint) GetNewsById(_ context.Context, req *pb.GetNewsByIdRequest) (*pb.GetNewsByIdResponse, error) {
	if req.Id == 0 {
		return nil, errors.New("id новости не указан")
	}

	newsData, err := e.News.GetNewsByIdService(req.Id)
	if err != nil {
		return &pb.GetNewsByIdResponse{}, err
	}

	return newsData, nil
}

func (e *Endpoint) GetNewsByCategory(_ context.Context, req *pb.GetNewsByCategoryRequest) (*pb.GetNewsResponse, error) {
	if req.Categories == "" {
		return nil, errors.New("id категории не указан")
	}

	newsData, err := e.News.GetNewsByCategoryService(req.Categories)
	if err != nil {
		return &pb.GetNewsResponse{}, err
	}

	return &pb.GetNewsResponse{News: newsData}, nil
}

func (e *Endpoint) AddNews(_ context.Context, req *pb.AddNewsRequest) (*pb.AddNewsResponse, error) {
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

	id, err := e.News.AddNewsService(req.Title, req.Text, req.Datetime, req.Categories)
	if err != nil {
		return &pb.AddNewsResponse{Err: error.Error(err)}, err
	}

	return &pb.AddNewsResponse{Id: id}, nil
}

func (e *Endpoint) DelNews(_ context.Context, req *pb.DelNewsRequest) (*pb.DelNewsResponse, error) {
	if req.Id == 0 {
		return nil, errors.New("id новости не указан")
	}

	err := e.News.DelNewsService(req.Id)
	if err != nil {
		return &pb.DelNewsResponse{Err: error.Error(err)}, err
	}

	return &pb.DelNewsResponse{Err: ""}, nil
}
