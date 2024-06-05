.PHONY: run-http
run-http:
	go run cmd/rate-limiter/main_http.go

.PHONY: run-grpc
run-grpc:
	go run cmd/rate-limiter/main_grpc.go

.PHONY: gen-go
gen-go:
	protoc -I=api --go_out=internal/pb --go-grpc_out=internal/pb api/service.proto

.PHONY: gen
gen:
	buf generate