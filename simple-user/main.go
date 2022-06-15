package main

import (
	"context"
	"log"
	"net"

	"github.com/zinirun/learn-go-grpc/data"
	userpb "github.com/zinirun/learn-go-grpc/protos/v1/user"
	"google.golang.org/grpc"
)

const port = "9000"

type userServer struct {
	userpb.UserServer
}

func (s *userServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	userId := req.UserId

	var userMessage *userpb.UserMessage
	for _, u := range data.Users {
		if u.UserId == userId {
			userMessage = u
			break
		}
	}

	return &userpb.GetUserResponse{
		UserMessage: userMessage,
	}, nil
}

func (s *userServer) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	return &userpb.ListUsersResponse{
		UserMessages: data.Users,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServer(grpcServer, &userServer{})

	log.Printf("start gRPC server on %s port", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
