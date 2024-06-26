package grpcAuth

import (
	"context"
	"errors"
	pb "moscowhack/gen/go/auth"
	"moscowhack/internal/app/errorz"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	UserLogin(string, string) (string, error)
	NewUserRegistration(string, string) (string, error)
	IsAdmin(string) (bool, error)
}

type Endpoint struct {
	Auth Auth
	pb.UnimplementedAuthServiceServer
}

func (e *Endpoint) Login(_ context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "username пустое значение")
	}

	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password пустое значение")
	}

	token, err := e.Auth.UserLogin(req.Email, req.Password)
	if err != nil {
		if errors.Is(err, errorz.ErrUserNotFound) {
			return nil, status.Error(codes.InvalidArgument, "неверный username или password")
		}

		return nil, status.Error(codes.Internal, "ошибка аутентификации")
	}

	return &pb.AuthResponse{Key: token}, nil
}

func (e *Endpoint) Registration(_ context.Context, req *pb.RegistrationRequest) (*pb.AuthResponse, error) {
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "username пустое значение")
	}

	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password пустое значение")
	}

	token, err := e.Auth.NewUserRegistration(ctx, req)
	if err != nil {
		if errors.Is(err, errorz.ErrUserExists) {
			return nil, status.Error(codes.InvalidArgument, "пользователь с таким именем уже существует")
		}

		return nil, status.Error(codes.Internal, "ошибка регистрации")
	}

	return &pb.AuthResponse{Key: token}, nil
}

func (e *Endpoint) IsAdmin(_ context.Context, req *pb.IsAdminRequest) (*pb.IsAdminResponse, error) {
	if req.JwtToken == "" {
		return nil, status.Error(codes.InvalidArgument, "jwtToken пустое значение")
	}

	isAdmin, err := e.Auth.IsAdmin(req.JwtToken)
	if err != nil {
		if errors.Is(err, errorz.ErrUserNotFound) {
			return nil, status.Error(codes.InvalidArgument, "пользователь с таким именем не существует")
		}

		return nil, status.Error(codes.Internal, "ошибка проверки прав пользователя")
	}

	return &pb.IsAdminResponse{IsAdmin: isAdmin}, nil
}
