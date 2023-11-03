package main

import (
	"drive-connect/application"
	"drive-connect/db"
	"drive-connect/lib/grpc_board"
	"drive-connect/lib/grpc_user"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := 50051
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	dbConn := db.ConnectionDB()
	dbConn.CreateTable()
	grpc_user.RegisterUserServiceServer(s, application.NewUserService(dbConn))
	grpc_board.RegisterBoardServiceServer(s, application.NewBoardService(dbConn))

	reflection.Register(s)

	log.Printf("start gRPC server port: %v", port)
	s.Serve(listener)

}
