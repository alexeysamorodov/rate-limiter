package ratelimiter

import (
	"context"

	pb "github.com/alexeysamorodov/rate-limiter/gen"
)

func (i *Implementation) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{
		Message: "Hello, grpc! " + req.GetMessage(),
	}, nil
}
