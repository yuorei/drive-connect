package application

import (
	"context"
	"drive-connect/db"

	"drive-connect/db/model"
	"drive-connect/lib/grpc_user"
	"time"

	"github.com/google/uuid"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewUserService(db *db.DB) *userService {
	return &userService{
		db: db,
	}
}

type userService struct {
	grpc_user.UnimplementedUserServiceServer
	db *db.DB
}

func (s *userService) GetUserById(ctx context.Context, request *grpc_user.ID) (*grpc_user.User, error) {
	user, err := s.db.GetUserById(request.Id)
	if err != nil {
		return nil, err
	}

	createAt := timestamppb.New(user.CreatedAt)
	updateAt := timestamppb.New(user.UpdatedAt)

	return &grpc_user.User{
		Id:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		Password:        user.Password,
		IsDriver:        user.Driver,
		ProfileImageURL: user.ProfileImageURL,
		CreatedAt:       createAt,
		UpdatedAt:       updateAt,
	}, nil
}

func (s *userService) CreateUser(ctx context.Context, request *grpc_user.User) (*grpc_user.User, error) {
	now := time.Now()
	user := model.User{
		ID:              uuid.New().String(),
		Name:            request.Name,
		Email:           request.Email,
		Password:        request.Password,
		Driver:          request.IsDriver,
		ProfileImageURL: request.ProfileImageURL,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	err := s.db.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	createAt := timestamppb.New(time.Now())
	updateAt := createAt

	return &grpc_user.User{
		Id:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		Password:        user.Password,
		IsDriver:        user.Driver,
		ProfileImageURL: user.ProfileImageURL,
		CreatedAt:       createAt,
		UpdatedAt:       updateAt,
	}, nil
}

func (s *userService) UpdateUser(ctx context.Context, request *grpc_user.User) (*grpc_user.User, error) {
	now := time.Now()
	user := &model.User{
		ID:              request.Id,
		Name:            request.Name,
		Email:           request.Email,
		Password:        request.Password,
		Driver:          request.IsDriver,
		ProfileImageURL: request.ProfileImageURL,
		CreatedAt:       request.CreatedAt.AsTime(),
		UpdatedAt:       now,
	}

	err := s.db.UpdateUser(user.ID, user)
	if err != nil {
		return nil, err
	}

	createAt := request.CreatedAt
	updateAt := timestamppb.New(now)

	return &grpc_user.User{
		Id:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		Password:        user.Password,
		IsDriver:        user.Driver,
		ProfileImageURL: user.ProfileImageURL,
		CreatedAt:       createAt,
		UpdatedAt:       updateAt,
	}, nil
}

func (s *userService) DeleteUser(ctx context.Context, request *grpc_user.ID) (*emptypb.Empty, error) {
	err := s.db.DeleteUser(request.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
