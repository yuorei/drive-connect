package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"drive-connect-bff/graph/generated"
	"drive-connect-bff/graph/model"
	"drive-connect-bff/middleware"
	"fmt"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	responseUser, err := r.client.CreateUser(input)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:        responseUser.Id,
		Name:      responseUser.Name,
		Email:     responseUser.Email,
		IsDriver:  responseUser.IsDriver,
		CreatedAt: responseUser.CreatedAt.String(),
		UpdatedAt: responseUser.UpdatedAt.String(),
	}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	responseUser, err := r.client.GetUserList()
	if err != nil {
		return nil, err
	}

	var users []*model.User
	timeformat := "2006-01-02T15:04:05Z"
	for _, user := range responseUser {
		createdAt := user.CreatedAt.AsTime().Format(timeformat)
		updatedAt := user.UpdatedAt.AsTime().Format(timeformat)

		users = append(users, &model.User{
			ID:        user.Id,
			Name:      user.Name,
			Email:     user.Email,
			IsDriver:  user.IsDriver,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}

	return users, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	responseUser, err := r.client.GetUserById(middleware.CtxValue(ctx).UserID)
	if err != nil {
		return nil, err
	}

	timeformat := "2006-01-02T15:04:05Z"

	createdAt := responseUser.CreatedAt.AsTime().Format(timeformat)
	updatedAt := responseUser.UpdatedAt.AsTime().Format(timeformat)

	return &model.User{
		ID:        responseUser.Id,
		Name:      responseUser.Name,
		Email:     responseUser.Email,
		IsDriver:  responseUser.IsDriver,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
