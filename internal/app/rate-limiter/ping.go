package ratelimiter

import (
	"context"

	pb "github.com/alexeysamorodov/rate-limiter/internal/pb/github.com/alexeysamorodov/rate-limiter/api"
)

func (i *Implementation) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{
		Message: "Hello, grpc!",
	}, nil
}
