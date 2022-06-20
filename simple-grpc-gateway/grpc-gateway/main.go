package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	userpb "github.com/zinirun/learn-go-grpc/protos/v2/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port           = "9000"
	gRPCServerPort = "9001"
)

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	options := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := userpb.RegisterUserHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:"+gRPCServerPort,
		options,
	); err != nil {
		log.Fatalf("failed to register gRPC gateway: %v", err)
	}

	log.Printf("start HTTP server on %s port", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
