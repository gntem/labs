package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "remindme/protos"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "9000"
	}

	conn, err := grpc.Dial("localhost:"+grpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewRemindmeServiceClient(conn)
	ctx := context.Background()

	reminderTime := time.Now().Add(10 * time.Second).Unix()

	resp, err := client.SetReminder(ctx, &pb.SetReminderRequest{
		UserId:       "user123",
		ReminderText: "Test reminder - Hello from the past!",
		ReminderTime: reminderTime,
	})
	if err != nil {
		log.Fatalf("failed to set reminder: %v", err)
	}
	log.Printf("set reminder response: %s", resp.Message)

	reminders, err := client.GetReminders(ctx, &pb.GetRemindersRequest{
		UserId: "user123",
	})
	if err != nil {
		log.Fatalf("failed to get reminders: %v", err)
	}
	log.Printf("current reminders count: %d", len(reminders.Reminders))
	for _, reminder := range reminders.Reminders {
		log.Printf("- %s (time: %d)", reminder.ReminderText, reminder.ReminderTime)
	}

	log.Println("Waiting 15 seconds for reminder to trigger...")
	time.Sleep(15 * time.Second)

	reminders, err = client.GetReminders(ctx, &pb.GetRemindersRequest{
		UserId: "user123",
	})
	if err != nil {
		log.Fatalf("failed to get reminders: %v", err)
	}
	log.Printf("reminders after trigger count: %d", len(reminders.Reminders))
	for _, reminder := range reminders.Reminders {
		log.Printf("- %s (time: %d)", reminder.ReminderText, reminder.ReminderTime)
	}
}
