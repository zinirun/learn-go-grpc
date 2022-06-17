package data

import (
	postpb "github.com/zinirun/learn-go-grpc/protos/v1/post"
)

type PostData struct {
	UserId string
	Posts  []*postpb.PostMessage
}

var UserPosts = []*PostData{
	{
		UserId: "1",
		Posts: []*postpb.PostMessage{
			{
				PostId: "1",
				Author: "", // filled from user server
				Title:  "Title 1",
				Body:   "Body 1",
				Tags:   []string{"gRPC", "Golang", "server", "protobuf"},
			},
			{
				PostId: "2",
				Author: "", // filled from user server
				Title:  "Title 2",
				Body:   "Body 2",
				Tags:   []string{"gRPC", "Golang", "server", "protobuf"},
			},
			{
				PostId: "3",
				Author: "", // filled from user server
				Title:  "Title 3",
				Body:   "Body 3",
				Tags:   []string{"gRPC", "Golang", "server", "protobuf"},
			},
		},
	},
}
