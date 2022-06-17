package main

import (
	"context"
	"log"
	"net"

	"github.com/zinirun/learn-go-grpc/data"
	postpb "github.com/zinirun/learn-go-grpc/protos/v1/post"
	userpb "github.com/zinirun/learn-go-grpc/protos/v1/user"
	client "github.com/zinirun/learn-go-grpc/simple-client-server"
	"google.golang.org/grpc"
)

const port = "9001"

type postServer struct {
	postpb.PostServer
	userCli userpb.UserClient
}

func (s *postServer) ListPostsByUserId(ctx context.Context, req *postpb.ListPostsByUserIdRequest) (*postpb.ListPostsByUserIdResponse, error) {
	userId := req.UserId

	// server:server grpc connection
	resp, err := s.userCli.GetUser(ctx, &userpb.GetUserRequest{UserId: userId})
	if err != nil {
		return nil, err
	}

	var postMessages []*postpb.PostMessage
	for _, userPost := range data.UserPosts {
		// find and fill "Author"
		if userPost.UserId == userId {
			for _, post := range userPost.Posts {
				post.Author = resp.UserMessage.Name
			}
			postMessages = userPost.Posts
		}
	}

	return &postpb.ListPostsByUserIdResponse{
		PostMessages: postMessages,
	}, nil
}

func (s *postServer) ListPosts(ctx context.Context, req *postpb.ListPostsRequest) (*postpb.ListPostsResponse, error) {
	var postMessages []*postpb.PostMessage

	for _, userPost := range data.UserPosts {
		resp, err := s.userCli.GetUser(ctx, &userpb.GetUserRequest{UserId: userPost.UserId})
		if err != nil {
			return nil, err
		}
		for _, post := range userPost.Posts {
			post.Author = resp.UserMessage.Name
		}
		postMessages = append(postMessages, userPost.Posts...)
	}

	return &postpb.ListPostsResponse{
		PostMessages: postMessages,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	userCli := client.GetUserClient("localhost:9000")
	grpcServer := grpc.NewServer()
	postpb.RegisterPostServer(grpcServer, &postServer{
		userCli: userCli,
	})

	log.Printf("start gRPC server on %s port", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
