package application

import (
	"context"
	"drive-connect/db"

	"drive-connect/db/model"
	"drive-connect/lib/grpc_board"
	"time"

	"github.com/google/uuid"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewBoardService(db *db.DB) *boardService {
	return &boardService{
		db: db,
	}
}

type boardService struct {
	grpc_board.UnimplementedBoardServiceServer
	db *db.DB
}

func (s *boardService) GetBoardById(ctx context.Context, request *grpc_board.ID) (*grpc_board.Board, error) {
	board, err := s.db.GetBoardById(request.Id)
	if err != nil {
		return nil, err
	}
	startTime := timestamppb.New(board.StartTime)
	createAt := timestamppb.New(board.CreatedAt)
	updateAt := timestamppb.New(board.UpdatedAt)

	return &grpc_board.Board{
		Id:                   board.ID,
		Type:                 board.Type,
		UserId:               board.UserID,
		Description:          board.Description,
		DepartureLatitude:    float32(board.DepartureLatitude),
		DepartureLongitude:   float32(board.DepartureLongitude),
		DestinationLatitude:  float32(board.DestinationLatitude),
		DestinationLongitude: float32(board.DestinationLongitude),
		Reward:               board.Reward,
		StartTime:            startTime,
		CreatedAt:            createAt,
		UpdatedAt:            updateAt,
	}, nil
}

func (s *boardService) CreateBoard(ctx context.Context, request *grpc_board.Board) (*grpc_board.Board, error) {
	now := time.Now()
	board := model.Board{
		ID:                   uuid.New().String(),
		Type:                 request.Type,
		UserID:               request.UserId,
		Description:          request.Description,
		DepartureLatitude:    float64(request.DepartureLatitude),
		DepartureLongitude:   float64(request.DepartureLongitude),
		DestinationLatitude:  float64(request.DestinationLatitude),
		DestinationLongitude: float64(request.DestinationLongitude),
		Reward:               request.Reward,
		StartTime:            request.StartTime.AsTime(),
		CreatedAt:            now,
		UpdatedAt:            now,
	}

	err := s.db.CreateBoard(&board)
	if err != nil {
		return nil, err
	}

	startTime := timestamppb.New(board.StartTime)
	createAt := timestamppb.New(time.Now())
	updateAt := createAt

	return &grpc_board.Board{
		Id:                   board.ID,
		Type:                 board.Type,
		UserId:               board.UserID,
		Description:          board.Description,
		DepartureLatitude:    float32(board.DepartureLatitude),
		DepartureLongitude:   float32(board.DepartureLongitude),
		DestinationLatitude:  float32(board.DestinationLatitude),
		DestinationLongitude: float32(board.DestinationLongitude),
		Reward:               board.Reward,
		StartTime:            startTime,
		CreatedAt:            createAt,
		UpdatedAt:            updateAt,
	}, nil
}

func (s *boardService) UpdateBoard(ctx context.Context, request *grpc_board.Board) (*grpc_board.Board, error) {
	now := time.Now()
	board := &model.Board{
		ID:                   request.Id,
		Type:                 request.Type,
		UserID:               request.UserId,
		Description:          request.Description,
		DepartureLatitude:    float64(request.DepartureLatitude),
		DepartureLongitude:   float64(request.DepartureLongitude),
		DestinationLatitude:  float64(request.DestinationLatitude),
		DestinationLongitude: float64(request.DestinationLongitude),
		Reward:               request.Reward,
		StartTime:            request.StartTime.AsTime(),
		CreatedAt:            request.CreatedAt.AsTime(),
		UpdatedAt:            now,
	}

	err := s.db.UpdateBoard(board.ID, board)
	if err != nil {
		return nil, err
	}

	startTime := timestamppb.New(board.StartTime)
	createAt := request.CreatedAt
	updateAt := timestamppb.New(now)

	return &grpc_board.Board{
		Id:                   board.ID,
		Type:                 board.Type,
		UserId:               board.UserID,
		Description:          board.Description,
		DepartureLatitude:    float32(board.DepartureLatitude),
		DepartureLongitude:   float32(board.DepartureLongitude),
		DestinationLatitude:  float32(board.DestinationLatitude),
		DestinationLongitude: float32(board.DestinationLongitude),
		Reward:               board.Reward,
		StartTime:            startTime,
		CreatedAt:            createAt,
		UpdatedAt:            updateAt,
	}, nil
}

func (s *boardService) DeleteBoard(ctx context.Context, request *grpc_board.ID) (*emptypb.Empty, error) {
	err := s.db.DeleteBoard(request.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
