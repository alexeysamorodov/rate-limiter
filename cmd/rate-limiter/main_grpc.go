package main

import (
	"log"
	"net"

	pb "github.com/alexeysamorodov/rate-limiter/gen"
	ratelimiter "github.com/alexeysamorodov/rate-limiter/internal/app/rate-limiter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	limiter := ratelimiter.NewRPSRateLimiter(1)
	defer limiter.Stop()

	// Создаем gRPC сервер с перехватчиком
	server := grpc.NewServer(
		grpc.UnaryInterceptor(ratelimiter.RateLimitInterceptor(limiter)),
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
