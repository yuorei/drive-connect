package application

import (
	"context"
	"drive-connect/db"
	"drive-connect/lib/grpc_back"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func NewAuthService(db *db.DB) *authService {
	return &authService{
		db: db,
	}
}

type authService struct {
	grpc_back.UnimplementedAuthServiceServer
	db *db.DB
}

func (s *authService) Login(ctx context.Context, request *grpc_back.LoginRequest) (*grpc_back.LoginResponse, error) {
	user, err := s.db.GetUserByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
	}).SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &grpc_back.LoginResponse{
		Token: token,
	}, nil
}
