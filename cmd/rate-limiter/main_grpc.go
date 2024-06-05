package main

import (
	"context"
	"log"
	"net"

	ratelimiter "github.com/alexeysamorodov/rate-limiter/internal/app/rate-limiter"
	pb "github.com/alexeysamorodov/rate-limiter/internal/pb/github.com/alexeysamorodov/rate-limiter/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func main() {
	limiter := ratelimiter.NewRPSRateLimiter(1) // 2 rps

	// Создаем gRPC сервер с перехватчиком
	server := grpc.NewServer(
		grpc.UnaryInterceptor(RateLimitInterceptor(limiter)),
	)

	// Регистрируем ваш сервис
	pb.RegisterExampleServiceServer(server, &ratelimiter.Implementation{})

	// Включаем рефлексию (для gRPC CLI и других инструментов)
	reflection.Register(server)

	// Настраиваем прослушивание порта
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Запускаем сервер
	log.Printf("Server is listening on port %s", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func RateLimitInterceptor(rl *ratelimiter.RPSRateLimiter) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if !rl.Allow() {
			return nil, status.Errorf(codes.ResourceExhausted, "rate limit exceeded")
		}
		return handler(ctx, req)
	}
}
