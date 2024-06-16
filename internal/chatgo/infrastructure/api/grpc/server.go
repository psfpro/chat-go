package grpc

import (
	"chatgo/internal/chatgo/application"
	"chatgo/internal/chatgo/domain"
	"chatgo/proto"
	"context"
	"github.com/gofrs/uuid"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ChatGoServer struct {
	proto.UnimplementedChatGoServer
	authenticationService   application.AuthenticationService
	userLoginHandler        *application.UserLoginHandler
	userRegistrationHandler *application.UserRegistrationHandler
	addTaskHandler          *application.AddTask
	taskRepository          domain.TaskRepository
}

func NewChatGoServer(
	authenticationService application.AuthenticationService,
	userLoginHandler *application.UserLoginHandler,
	userRegistrationHandler *application.UserRegistrationHandler,
	addTaskHandler *application.AddTask,
	taskRepository domain.TaskRepository,
) *ChatGoServer {
	return &ChatGoServer{
		authenticationService:   authenticationService,
		userLoginHandler:        userLoginHandler,
		userRegistrationHandler: userRegistrationHandler,
		addTaskHandler:          addTaskHandler,
		taskRepository:          taskRepository,
	}
}

func (s *ChatGoServer) UserLogin(ctx context.Context, in *proto.UserLoginRequest) (*proto.UserAuthenticationResponse, error) {
	accessToken, refreshToken, err := s.userLoginHandler.Handle(ctx, in.Login, in.Password)
	if err != nil {
		return &proto.UserAuthenticationResponse{
			AccessToken:  "",
			RefreshToken: "",
			Error:        err.Error(),
		}, nil
	}

	return &proto.UserAuthenticationResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Error:        "",
	}, nil
}

func (s *ChatGoServer) UserRegistration(ctx context.Context, in *proto.UserRegistrationRequest) (*proto.UserAuthenticationResponse, error) {
	accessToken, refreshToken, err := s.userRegistrationHandler.Handle(ctx, in.Login, in.Password)
	if err != nil {
		return &proto.UserAuthenticationResponse{
			AccessToken:  "",
			RefreshToken: "",
			Error:        err.Error(),
		}, nil
	}

	return &proto.UserAuthenticationResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Error:        "",
	}, nil
}

func (s *ChatGoServer) AddTask(ctx context.Context, in *proto.AddTaskRequest) (*proto.AddTaskResponse, error) {
	userID, err := UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	id, err := s.addTaskHandler.Handle(ctx, userID, in.Title, in.Description)
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

func (s *ChatGoServer) GetTasks(ctx context.Context, _ *proto.GetTasksRequest) (*proto.GetTasksResponse, error) {
	userID, err := UserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	tasks, err := s.taskRepository.GetAllByUserId(ctx, userID)
	if err != nil {
		return nil, err
	}
	var res []*proto.Task
	for _, task := range tasks {
		res = append(res, &proto.Task{
			Id:          task.ID.String(),
			Title:       task.Title,
			Description: task.Description,
			State:       string(task.State),
		})
	}

	return &proto.GetTasksResponse{
		Task: res,
	}, nil
}

func (s *ChatGoServer) AuthFunc(ctx context.Context, method string) (context.Context, error) {
	if method == "/chatgo.ChatGo/UserLogin" || method == "/chatgo.ChatGo/UserRegistration" {
		return ctx, nil
	}
	token, err := grpcauth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	userId, err := s.authenticationService.GetUserID(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	newCtx := context.WithValue(ctx, &key{}, userId)
	return newCtx, nil
}

func UserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	userID, ok := ctx.Value(&key{}).(uuid.UUID)
	if !ok {
		return uuid.Nil, status.Errorf(codes.Unauthenticated, "user unauthenticated")
	}

	return userID, nil
}

type key struct{}

// UnaryServerInterceptor returns a new unary server interceptors that performs per-request auth.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var newCtx context.Context
		var err error
		if overrideSrv, ok := info.Server.(*ChatGoServer); ok {
			newCtx, err = overrideSrv.AuthFunc(ctx, info.FullMethod)
		}
		if err != nil {
			return nil, err
		}
		return handler(newCtx, req)
	}
}

// StreamServerInterceptor returns a new unary server interceptors that performs per-request auth.
func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		var newCtx context.Context
		var err error
		if overrideSrv, ok := srv.(ChatGoServer); ok {
			newCtx, err = overrideSrv.AuthFunc(stream.Context(), info.FullMethod)
		}
		if err != nil {
			return err
		}
		wrapped := grpcmiddleware.WrapServerStream(stream)
		wrapped.WrappedContext = newCtx
		return handler(srv, wrapped)
	}
}
