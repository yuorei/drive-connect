package client

import "drive-connect-bff/lib/grpc_back"

type Client struct {
	userClient  grpc_back.UserServiceClient
	boardClient grpc_back.BoardServiceClient
}

func NewClient() *Client {
	client := &Client{}
	client.NewClient()
	return client
}
