package client

import (
	"context"
	"drive-connect-bff/graph/model"
	"drive-connect-bff/lib/grpc_back"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *Client) CreateUser(input model.UserInput) (*grpc_back.User, error) {
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

func (c *Client) GetUserList() ([]*grpc_back.User, error) {
	response, err := c.userClient.GetUserList(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return response.Users, nil
}

func (c *Client) GetUserById(id string) (*grpc_back.User, error) {
	request := &grpc_back.UserID{
		Id: id,
	}

	response, err := c.userClient.GetUserById(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
