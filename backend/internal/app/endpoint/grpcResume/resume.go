package grpcResume

import (
	"context"
	"errors"
	pb "moscowhack/gen/go/resume"
	"moscowhack/internal/app/errorz"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Resume interface {
	ResumeParser(context.Context, []byte) (map[string]string, error)
}

type Endpoint struct {
	Resume Resume
	pb.UnimplementedResumeServiceServer
}

func New(resumeParser Resume) *Endpoint {
	return &Endpoint{
		Resume: resumeParser,
	}
}

func (e *Endpoint) ParseResume(ctx context.Context, req *pb.ResumeRequest) (*pb.ResumeResponse, error) {
	if req.ResumeDoc == nil {
		return nil, status.Error(codes.InvalidArgument, "username пустое значение")
	}

	resumeMap, err := e.Resume.ResumeParser(ctx, req.ResumeDoc)
	if err != nil {
		if errors.Is(err, errorz.ErrPanicHandle) {
			return nil, status.Error(codes.InvalidArgument, "ошибка ")
		}

		return nil, status.Error(codes.Internal, "ошибка аутентификации")
	}

	return &pb.ResumeResponse{ResumeMap: resumeMap}, nil
}
