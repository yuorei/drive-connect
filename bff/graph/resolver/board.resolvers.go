package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"drive-connect-bff/graph/model"
	"drive-connect-bff/middleware"
	"fmt"
)

// CreateBoard is the resolver for the createBoard field.
func (r *mutationResolver) CreateBoard(ctx context.Context, input model.BoardInput) (*model.Board, error) {
	board, err := r.client.CreateBoard(input, middleware.CtxValue(ctx).UserID)
	if err != nil {
		return nil, err
	}

	return &model.Board{
		ID: board.Id,
		User: &model.User{
			ID: middleware.CtxValue(ctx).UserID,
		},
		Description:          board.Description,
		DepartureLatitude:    float64(board.DepartureLatitude),
		DepartureLongitude:   float64(board.DepartureLongitude),
		DestinationLatitude:  float64(board.DestinationLatitude),
		DestinationLongitude: float64(board.DestinationLongitude),
		Reward:               board.Reward,
		StartTime:            board.StartTime.String(),
	}, nil
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input model.CommentInput) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented: CreateComment - createComment"))
}

// UpdateBoard is the resolver for the updateBoard field.
func (r *mutationResolver) UpdateBoard(ctx context.Context, id string, input model.BoardInput) (*model.Board, error) {
	panic(fmt.Errorf("not implemented: UpdateBoard - updateBoard"))
}

// UpdateComment is the resolver for the updateComment field.
func (r *mutationResolver) UpdateComment(ctx context.Context, id string, input model.CommentInput) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented: UpdateComment - updateComment"))
}

// DeleteBoard is the resolver for the deleteBoard field.
func (r *mutationResolver) DeleteBoard(ctx context.Context, id string) (*model.Board, error) {
	panic(fmt.Errorf("not implemented: DeleteBoard - deleteBoard"))
}

// DeleteComment is the resolver for the deleteComment field.
func (r *mutationResolver) DeleteComment(ctx context.Context, id string) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented: DeleteComment - deleteComment"))
}

// Boards is the resolver for the boards field.
func (r *queryResolver) Boards(ctx context.Context) ([]*model.Board, error) {
	boards, err := r.client.GetBoards()
	if err != nil {
		return nil, err
	}
	var result []*model.Board
	timeformat := "2006-01-02T15:04:05Z"
	for _, board := range boards {
		startTime := board.StartTime.AsTime().Format(timeformat)
		createdAt := board.CreatedAt.AsTime().Format(timeformat)
		updatedAt := board.StartTime.AsTime().Format(timeformat)
		result = append(result, &model.Board{
			ID: board.Id,
			User: &model.User{
				ID: board.UserId,
			},
			Description:          board.Description,
			DepartureLatitude:    float64(board.DepartureLatitude),
			DepartureLongitude:   float64(board.DepartureLongitude),
			DestinationLatitude:  float64(board.DestinationLatitude),
			DestinationLongitude: float64(board.DestinationLongitude),
			Reward:               board.Reward,
			StartTime:            startTime,
			CreatedAt:            &createdAt,
			UpdatedAt:            &updatedAt,
		})
	}

	return result, nil
}

// Board is the resolver for the board field.
func (r *queryResolver) Board(ctx context.Context, id string) (*model.Board, error) {
	board, err := r.client.GetBoard(id)
	if err != nil {
		return nil, err
	}
	timeformat := "2006-01-02T15:04:05Z"
	startTime := board.StartTime.AsTime().Format(timeformat)
	createdAt := board.CreatedAt.AsTime().Format(timeformat)
	updatedAt := board.StartTime.AsTime().Format(timeformat)
	return &model.Board{
		ID: board.Id,
		User: &model.User{
			ID: board.UserId,
		},
		Description:          board.Description,
		DepartureLatitude:    float64(board.DepartureLatitude),
		DepartureLongitude:   float64(board.DepartureLongitude),
		DestinationLatitude:  float64(board.DestinationLatitude),
		DestinationLongitude: float64(board.DestinationLongitude),
		Reward:               board.Reward,
		StartTime:            startTime,
		CreatedAt:            &createdAt,
		UpdatedAt:            &updatedAt,
	}, nil
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented: Comments - comments"))
}

// Comment is the resolver for the comment field.
func (r *queryResolver) Comment(ctx context.Context, id string) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented: Comment - comment"))
}
