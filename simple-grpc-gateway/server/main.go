package main

import (
	"context"
	"log"
	"net"

	"github.com/zinirun/learn-go-grpc/data"
	userpb "github.com/zinirun/learn-go-grpc/protos/v2/user"
	"google.golang.org/grpc"
)

const port = "9001"

type userServer struct {
	userpb.UserServer
}

func (s *userServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	userId := req.UserId

	var userMessage *userpb.UserMessage
	for _, u := range data.UsersV2 {
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
		UserMessages: data.UsersV2,
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
