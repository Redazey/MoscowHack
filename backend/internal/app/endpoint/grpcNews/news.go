package grpcNews

import (
	pb "moscowhack/gen/go/news"
	"strconv"
)

type Service interface {
	GetNewsFromDB(id int) (*pb.NewsItem, error)
}

type Endpoint struct {
	s Service
	pb.UnimplementedNewsServiceServer
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) NewsById(req *pb.NewsRequest) (*pb.NewsResponse, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return &pb.NewsResponse{}, err
	}

	newsData, err := e.s.GetNewsFromDB(id)
	if err != nil {
		return &pb.NewsResponse{}, err
	}

	newsItem := map[string]*pb.NewsItem{"NewsItem": newsData}

	return &pb.NewsResponse{News: newsItem}, nil
}
