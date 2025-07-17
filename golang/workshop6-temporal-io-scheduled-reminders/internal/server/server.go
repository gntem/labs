package server

import (
	"context"
	"fmt"
	"os"
	"time"

	"remindme/internal/storage"
	"remindme/internal/workflows"
	pb "remindme/protos"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RemindmeServer struct {
	pb.UnimplementedRemindmeServiceServer
	temporalClient client.Client
	storage        *storage.Storage
}

func NewRemindmeServer(temporalClient client.Client) *RemindmeServer {
	storageFile := os.Getenv("REMINDER_DATA_FILE")
	if storageFile == "" {
		storageFile = "reminders.json"
	}

	return &RemindmeServer{
		temporalClient: temporalClient,
		storage:        storage.NewStorage(storageFile),
	}
}

func (s *RemindmeServer) SetReminder(ctx context.Context, req *pb.SetReminderRequest) (*pb.SetReminderResponse, error) {
	if req.UserId == "" || req.ReminderText == "" {
		return &pb.SetReminderResponse{
			Success: false,
			Message: "user_id and reminder_text are required",
		}, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.ReminderTime <= time.Now().Unix() {
		return &pb.SetReminderResponse{
			Success: false,
			Message: "reminder_time must be in the future",
		}, status.Error(codes.InvalidArgument, "invalid reminder time")
	}

	reminderID := uuid.New().String()
	workflowID := fmt.Sprintf("reminder-%s", reminderID)

	workflowReq := workflows.ReminderRequest{
		ID:           reminderID,
		UserID:       req.UserId,
		ReminderText: req.ReminderText,
		ReminderTime: req.ReminderTime,
	}

	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "remindme-task-queue",
	}

	_, err := s.temporalClient.ExecuteWorkflow(ctx, options, workflows.RemindmeWorkflow, workflowReq)
	if err != nil {
		return &pb.SetReminderResponse{
			Success: false,
			Message: fmt.Sprintf("failed to start workflow: %v", err),
		}, status.Error(codes.Internal, "failed to start workflow")
	}

	return &pb.SetReminderResponse{
		Success: true,
		Message: fmt.Sprintf("reminder set with ID: %s", reminderID),
	}, nil
}

func (s *RemindmeServer) GetReminders(ctx context.Context, req *pb.GetRemindersRequest) (*pb.GetRemindersResponse, error) {
	if req.UserId == "" {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	reminders, err := s.storage.GetReminders(req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get reminders")
	}

	var pbReminders []*pb.Reminder
	for _, reminder := range reminders {
		pbReminders = append(pbReminders, &pb.Reminder{
			Id:           reminder.ID,
			UserId:       reminder.UserID,
			ReminderText: reminder.ReminderText,
			ReminderTime: reminder.ReminderTime,
		})
	}

	return &pb.GetRemindersResponse{
		Reminders: pbReminders,
	}, nil
}
