package endpoint

type Service interface {
}

type Endpoint struct {
	s Service
}

type NewsRequest struct{}

type NewsResponse struct{}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}
