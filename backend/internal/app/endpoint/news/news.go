package news

import (
	"moscowhack/internal/app/service/news"
	pb "moscowhack/protos/news"
	"strconv"
)

type Message struct {
	news map[string]map[string]string
}

type Service interface {
	GetNewsFromDB(id int) (map[string]*news.NewsItem, error)
}

type Endpoint struct {
	s      Service
	server NewsServiceServer
}

type NewsServiceServer struct {
	pb.UnimplementedAuthServiceServer
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) News(req *pb.NewsRequest) (*pb.NewsResponse, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return &pb.NewsResponse{}, err
	}

	newsData, err := news.GetNewsFromDB(id)
	if err != nil {
		return &pb.NewsResponse{}, err
	}

	return &pb.NewsResponse{News: newsData}, nil
}
