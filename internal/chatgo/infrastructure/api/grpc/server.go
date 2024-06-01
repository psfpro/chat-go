package grpc

import (
	"chatgo/internal/chatgo/application"
	"chatgo/proto"
	"context"
	"github.com/gofrs/uuid"
)

type ChatGoServer struct {
	proto.UnimplementedChatGoServer
	addTaskHandler *application.AddTask
}

func NewChatGoServer(addTaskHandler *application.AddTask) *ChatGoServer {
	return &ChatGoServer{addTaskHandler: addTaskHandler}
}

func (s *ChatGoServer) AddTask(_ context.Context, in *proto.AddTaskRequest) (*proto.AddTaskResponse, error) {
	id, err := s.addTaskHandler.Handle(in.Task.Title, in.Task.Description)
	if err != nil {
		return &proto.AddTaskResponse{
			Id:    uuid.Nil.String(),
			Error: err.Error(),
		}, nil
	}

	return &proto.AddTaskResponse{
		Id:    id.String(),
		Error: "",
	}, nil
}
