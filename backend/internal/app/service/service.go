package service

import "moscowhack/internal/app/service/news"

type Service struct {
	n news.Service
}

func New() *Service {
	return &Service{}
}
