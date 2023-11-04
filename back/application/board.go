package application

import (
	"context"
	"drive-connect/db"
	"fmt"

	"drive-connect/db/model"
	"drive-connect/lib/grpc_back"
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
	grpc_back.UnimplementedBoardServiceServer
	db *db.DB
}

func (s *boardService) ReadBoard(ctx context.Context, request *grpc_back.BoardID) (*grpc_back.Board, error) {
	board, err := s.db.GetBoardById(request.Id)
	if err != nil {
		return nil, err
	}
	startTime := timestamppb.New(board.StartTime)
	createAt := timestamppb.New(board.CreatedAt)
	updateAt := timestamppb.New(board.UpdatedAt)

	return &grpc_back.Board{
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

func (s *boardService) CreateBoard(ctx context.Context, request *grpc_back.Board) (*grpc_back.Board, error) {
	now := time.Now()
	fmt.Println(request.StartTime.AsTime(), "!!!!!!!!!!!!!!")
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

	return &grpc_back.Board{
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

func (s *boardService) UpdateBoard(ctx context.Context, request *grpc_back.Board) (*grpc_back.Board, error) {
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

	return &grpc_back.Board{
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

func (s *boardService) DeleteBoard(ctx context.Context, request *grpc_back.BoardID) (*emptypb.Empty, error) {
	err := s.db.DeleteBoard(request.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *boardService) ReadAllBoard(ctx context.Context, _ *emptypb.Empty) (*grpc_back.BoardList, error) {
	boards, err := s.db.ReadAllBoard()
	if err != nil {
		return nil, err
	}

	var boardList []*grpc_back.Board

	for _, board := range boards {
		startTime := timestamppb.New(board.StartTime)
		createAt := timestamppb.New(board.CreatedAt)
		updateAt := timestamppb.New(board.UpdatedAt)

		boardList = append(boardList, &grpc_back.Board{
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
		})
	}

	return &grpc_back.BoardList{
		Boards: boardList,
	}, nil
}

func (s *boardService) GetCommentById(ctx context.Context, request *grpc_back.BoardID) (*grpc_back.Comment, error) {
	board, err := s.db.GetCommentById(request.Id)
	if err != nil {
		return nil, err
	}

	createAt := timestamppb.New(board.CreatedAt)
	updateAt := timestamppb.New(board.UpdatedAt)

	return &grpc_back.Comment{
		Id:          board.ID,
		PostId:      board.PostID,
		CommenterId: board.CommenterID,
		Content:     board.Content,
		CreatedAt:   createAt,
		UpdatedAt:   updateAt,
	}, nil
}

func (s *boardService) CreateComment(ctx context.Context, request *grpc_back.Comment) (*grpc_back.Comment, error) {
	now := time.Now()
	board := model.Comment{
		ID:          uuid.New().String(),
		PostID:      request.PostId,
		CommenterID: request.CommenterId,
		Content:     request.Content,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	err := s.db.CreateComment(&board)
	if err != nil {
		return nil, err
	}

	createAt := timestamppb.New(time.Now())
	updateAt := createAt

	return &grpc_back.Comment{
		Id:          board.ID,
		PostId:      board.PostID,
		CommenterId: board.CommenterID,
		Content:     board.Content,
		CreatedAt:   createAt,
		UpdatedAt:   updateAt,
	}, nil
}

func (s *boardService) UpdateComment(ctx context.Context, request *grpc_back.Comment) (*grpc_back.Comment, error) {
	now := time.Now()
	board := &model.Comment{
		ID:          uuid.New().String(),
		PostID:      request.PostId,
		CommenterID: request.CommenterId,
		Content:     request.Content,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	err := s.db.UpdateComment(board.ID, board)
	if err != nil {
		return nil, err
	}

	createAt := request.CreatedAt
	updateAt := timestamppb.New(now)

	return &grpc_back.Comment{
		Id:          board.ID,
		PostId:      board.PostID,
		CommenterId: board.CommenterID,
		Content:     board.Content,
		CreatedAt:   createAt,
		UpdatedAt:   updateAt,
	}, nil
}

func (s *boardService) DeleteComment(ctx context.Context, request *grpc_back.BoardID) (*emptypb.Empty, error) {
	err := s.db.DeleteComment(request.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *boardService) GetCommentList(ctx context.Context, _ *emptypb.Empty) (*grpc_back.CommentList, error) {
	boards, err := s.db.GetCommentList()
	if err != nil {
		return nil, err
	}

	var boardList []*grpc_back.Comment

	for _, board := range boards {
		createAt := timestamppb.New(board.CreatedAt)
		updateAt := timestamppb.New(board.UpdatedAt)

		boardList = append(boardList, &grpc_back.Comment{
			Id:          board.ID,
			PostId:      board.PostID,
			CommenterId: board.CommenterID,
			Content:     board.Content,
			CreatedAt:   createAt,
			UpdatedAt:   updateAt,
		})
	}

	return &grpc_back.CommentList{
		Comments: boardList,
	}, nil
}
