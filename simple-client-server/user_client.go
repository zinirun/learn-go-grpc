package simple_client_server

import (
	"log"
	"sync"

	userpb "github.com/zinirun/learn-go-grpc/protos/v1/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	once sync.Once
	cli  userpb.UserClient
)

func GetUserClient(serviceHost string) userpb.UserClient {
	// singleton - init client once and use continuously
	once.Do(func() {
		// connection with target(grpc service)
		// - without secure options
		// - block before got connection (default: connection in background)
		conn, err := grpc.Dial(serviceHost, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
		if err != nil {
			log.Fatalf("failed to get user grpc-client: %v", err)
		}
		cli = userpb.NewUserClient((conn))
	})
	return cli
}
