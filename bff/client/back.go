package client

import (
	"drive-connect-bff/lib/grpc_back"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (c *Client) NewClient() {
	// TCPサーバーのアドレスを指定
	userAddress := "localhost:50051"
	// サーバーに接続する
	conn, err := grpc.Dial(
		userAddress,
		// コネクションでSSL/TLSを使用しない
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// コネクションが確立されるまで待機する(同期処理をする)
		grpc.WithBlock(),
	)

	if err != nil {
		log.Fatal("Connection failed. err: ", err)
		return
	}

	// gRPCクライアントを生成
	c.userClient = grpc_back.NewUserServiceClient(conn)
	c.boardClient = grpc_back.NewBoardServiceClient(conn)
	c.authClient = grpc_back.NewAuthServiceClient(conn)
}
