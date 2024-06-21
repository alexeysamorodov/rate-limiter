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

.PHONY: deps-go
deps-go:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest