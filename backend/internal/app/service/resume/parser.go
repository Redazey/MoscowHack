package resume

import (
	"context"
	"fmt"
	"moscowhack/config"
)

type Service struct {
}

func New(cfg *config.Configuration) *Service {
	return &Service{}
}

func (s *Service) ResumeParser(ctx context.Context, resumeDoc []byte) (map[string]string, error) {
	resumeMap := make(map[string]string)

	fmt.Print(resumeDoc)

	return resumeMap, nil
}
