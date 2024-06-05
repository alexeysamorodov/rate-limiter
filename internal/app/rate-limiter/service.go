package ratelimiter

import pb "github.com/alexeysamorodov/rate-limiter/internal/pb/github.com/alexeysamorodov/rate-limiter/api"

// Server - структура сервера
type Implementation struct {
	pb.ExampleServiceServer
}
