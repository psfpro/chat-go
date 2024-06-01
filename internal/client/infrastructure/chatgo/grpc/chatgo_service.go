package grpc

import (
	"chatgo/proto"
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type ChatGoService struct {
	client proto.ChatGoClient
}

func NewChatGoService() *ChatGoService {
	conn, err := grpc.NewClient(":3200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := proto.NewChatGoClient(conn)
	return &ChatGoService{client: client}
}

func (s *ChatGoService) AddTask(title string, description string) (uuid.UUID, error) {
	res, err := s.client.AddTask(context.Background(), &proto.AddTaskRequest{Task: &proto.Task{
		Title:       title,
		Description: description,
	}})
	if err != nil {
		return uuid.Nil, err
	}
	if res.Error != "" {
		return uuid.Nil, fmt.Errorf("add task err: %v", res.Error)
	}
	id, err := uuid.FromString(res.Id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid id: %v", err)
	}

	return id, nil
}
