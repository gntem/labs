package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"workshop4-grpc-leaderboard/proto"
	"workshop4-grpc-leaderboard/repository"
	"workshop4-grpc-leaderboard/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	repo := repository.NewMockLeaderboardRepository()
	leaderboardService := service.NewLeaderboardService(repo)
	proto.RegisterLeaderboardServiceServer(grpcServer, leaderboardService)

	reflection.Register(grpcServer)

	go func() {
		log.Printf("Starting gRPC server on port %d", port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down gRPC server...")
	grpcServer.GracefulStop()
	log.Println("Server stopped")
}
