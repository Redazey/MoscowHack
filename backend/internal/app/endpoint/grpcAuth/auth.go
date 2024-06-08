package grpcAuth

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "moscowhack/gen/go/auth"
	"moscowhack/internal/app/errorz"
)

type Auth interface {
	UserLogin(context.Context, string, string) (string, error)
	NewUserRegistration(context.Context, string, string) (string, error)
}

type Endpoint struct {
	Auth Auth
	pb.UnimplementedAuthServiceServer
}

func New(auth Auth) *Endpoint {
	return &Endpoint{
		Auth: auth,
	}
}

func (e *Endpoint) Login(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	if req.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "username пустое значение")
	}

	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password пустое значение")
	}

	token, err := e.Auth.UserLogin(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		if errors.Is(err, errorz.ErrUserNotFound) {
			return nil, status.Error(codes.InvalidArgument, "неверный username или password")
		}

		return nil, status.Error(codes.Internal, "ошибка аутентификации")
	}

	return &pb.AuthResponse{Key: token}, nil
}

// передаем в эту функцию username и password
func (e *Endpoint) Registration(ctx context.Context, req *pb.AuthRequest) (*pb.AuthResponse, error) {
	if req.Username == "" {
		return nil, status.Error(codes.InvalidArgument, "username пустое значение")
	}

	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password пустое значение")
	}

	token, err := e.Auth.NewUserRegistration(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		if errors.Is(err, errorz.ErrUserExists) {
			return nil, status.Error(codes.InvalidArgument, "пользователь с таким именем уже существует")
		}

		return nil, status.Error(codes.Internal, "ошибка регистрации")
	}

	return &pb.AuthResponse{Key: token}, nil
}
