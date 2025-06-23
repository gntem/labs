.PHONY: proto build run clean test

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
GOGET=$(GOCMD) get

# Binary name
BINARY_NAME=leaderboard-server

# Build directory
BUILD_DIR=build

# Main package path
MAIN_PACKAGE=.

# Proto parameters
PROTO_DIR=./proto
PROTO_FILE=$(PROTO_DIR)/leaderboard.proto
PROTO_GEN_DIR=$(PROTO_DIR)

# Set the default target to build
all: proto build

# Install required tools
tools:
	$(GOGET) -u google.golang.org/protobuf/cmd/protoc-gen-go
	$(GOGET) -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

# Generate Go code from proto
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		$(PROTO_FILE)

# Install dependencies
deps:
	$(GOMOD) tidy

# Build the binary
build: deps proto
	mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)

# Run the server
run: build
	$(BUILD_DIR)/$(BINARY_NAME)

# Run the tests
test:
	$(GOTEST) -v ./...

# Clean up
clean:
	rm -rf $(BUILD_DIR)
	rm -f $(PROTO_GEN_DIR)/*.pb.go

# Default target
default: all
