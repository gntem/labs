package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"workshop4-grpc-leaderboard/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	serverAddr := flag.String("addr", "localhost:50051", "The server address in the format host:port")
	flag.Parse()

	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewLeaderboardServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pageSize := int32(2)
	pageNumber := int32(1)

	fmt.Printf("Fetching leaderboards (page %d, size %d)...\n", pageNumber, pageSize)

	resp, err := client.ListLeaderboards(ctx, &proto.ListLeaderboardsRequest{
		PageSize:   pageSize,
		PageNumber: pageNumber,
	})
	if err != nil {
		log.Fatalf("ListLeaderboards failed: %v", err)
	}

	fmt.Printf("Total leaderboards: %d\n", resp.TotalCount)
	fmt.Printf("Current page: %d\n", resp.PageNumber)
	fmt.Println("Leaderboards:")
	for i, lb := range resp.Leaderboards {
		fmt.Printf("%d. ID: %s, User ID: %s, Name: %s, Score: %d\n", i+1, lb.Id, lb.UserId, lb.Name, lb.Score)
	}

	if resp.TotalCount > pageSize*pageNumber {
		pageNumber++
		fmt.Printf("\nFetching next page (page %d, size %d)...\n", pageNumber, pageSize)

		resp, err = client.ListLeaderboards(ctx, &proto.ListLeaderboardsRequest{
			PageSize:   pageSize,
			PageNumber: pageNumber,
		})
		if err != nil {
			log.Fatalf("ListLeaderboards failed: %v", err)
		}

		fmt.Printf("Total leaderboards: %d\n", resp.TotalCount)
		fmt.Printf("Current page: %d\n", resp.PageNumber)
		fmt.Println("Leaderboards:")
		for i, lb := range resp.Leaderboards {
			fmt.Printf("%d. ID: %s, User ID: %s, Name: %s, Score: %d\n", i+1, lb.Id, lb.UserId, lb.Name, lb.Score)
		}
	}
}
