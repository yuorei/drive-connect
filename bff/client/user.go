package client

import (
	"context"
	"drive-connect-bff/graph/model"
	"drive-connect-bff/lib/grpc_back"
)

func (c *Client) CreateUser(input model.UserInput) (*grpc_back.User, error) {

	// リクエストの生成
	request := &grpc_back.User{
		Name:  input.Name,
		Email: input.Email,
		// TODO: パスワードのハッシュ化
		Password: input.Password,
		IsDriver: input.IsDriver,
	}

	response, err := c.userClient.CreateUser(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
