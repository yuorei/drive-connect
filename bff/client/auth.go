package client

import (
	"context"
	"drive-connect-bff/graph/model"
	"drive-connect-bff/lib/grpc_back"
)

func (c *Client) Login(input model.LoginInput) (*grpc_back.LoginResponse, error) {

	// リクエストの生成
	request := &grpc_back.LoginRequest{
		Email:    input.Email,
		Password: input.Password,
	}

	response, err := c.authClient.Login(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
