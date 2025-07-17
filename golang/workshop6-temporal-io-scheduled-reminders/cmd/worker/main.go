package main

import (
	"log"
	"os"

	"remindme/internal/activities"
	"remindme/internal/storage"
	"remindme/internal/workflows"

	"github.com/joho/godotenv"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	temporalHost := os.Getenv("TEMPORAL_HOST")
	if temporalHost == "" {
		temporalHost = "localhost:7233"
	}

	storageFile := os.Getenv("REMINDER_DATA_FILE")
	if storageFile == "" {
		storageFile = "reminders.json"
	}

	c, err := client.Dial(client.Options{
		HostPort: temporalHost,
	})
	if err != nil {
		log.Fatalf("Unable to create Temporal client: %v", err)
	}
	defer c.Close()

	w := worker.New(c, "remindme-task-queue", worker.Options{})

	storage := storage.NewStorage(storageFile)
	remindActivity := activities.NewRemindActivity(storage)

	w.RegisterWorkflow(workflows.RemindmeWorkflow)
	w.RegisterActivityWithOptions(remindActivity.Execute, activity.RegisterOptions{Name: "RemindActivity"})

	log.Println("Starting Temporal worker...")
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalf("Unable to start worker: %v", err)
	}
}
