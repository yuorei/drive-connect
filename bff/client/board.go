package client

import (
	"context"
	"drive-connect-bff/graph/model"
	"drive-connect-bff/lib/grpc_back"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *Client) CreateBoard(input model.BoardInput, id string) (*grpc_back.Board, error) {
	t, err := time.Parse(time.RFC3339, input.StartTime)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return nil, err
	}
	ts := &timestamp.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}

	request := &grpc_back.Board{
		Id:                   id,
		Type:                 input.Type,
		Description:          input.Description,
		DepartureLatitude:    float32(input.DepartureLatitude),
		DepartureLongitude:   float32(input.DepartureLongitude),
		DestinationLatitude:  float32(input.DestinationLatitude),
		DestinationLongitude: float32(input.DestinationLongitude),
		Reward:               input.Reward,
		StartTime:            ts,
	}

	response, err := c.boardClient.CreateBoard(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetBoard(id string) (*grpc_back.Board, error) {
	request := &grpc_back.BoardID{
		Id: id,
	}

	response, err := c.boardClient.ReadBoard(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetBoards() ([]*grpc_back.Board, error) {
	response, err := c.boardClient.ReadAllBoard(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return response.Boards, nil
}
