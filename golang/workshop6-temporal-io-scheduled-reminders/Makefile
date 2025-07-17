.PHONY: proto build run worker clean test client

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		protos/remindme.proto

build: proto
	go build -o bin/server ./cmd/server
	go build -o bin/worker ./cmd/worker
	go build -o bin/client ./cmd/client

run: build
	./bin/server

worker: build
	./bin/worker

clean:
	rm -rf bin/
	rm -f protos/*.pb.go

deps:
	go mod tidy
	go mod download

up:
	docker compose up -d

down:
	docker compose down

logs:
	docker compose logs -f

test: build
	./bin/client

client: build
	./bin/client
