package activities

import (
	"context"
	"fmt"
	"time"

	"remindme/internal/storage"
)

type RemindActivity struct {
	storage *storage.Storage
}

func NewRemindActivity(storage *storage.Storage) *RemindActivity {
	return &RemindActivity{storage: storage}
}

type ReminderRequest struct {
	ID           string
	UserID       string
	ReminderText string
	ReminderTime int64
}

func (a *RemindActivity) Execute(ctx context.Context, req ReminderRequest) error {
	reminder := &storage.Reminder{
		ID:           req.ID,
		UserID:       req.UserID,
		ReminderText: req.ReminderText,
		ReminderTime: req.ReminderTime,
		CreatedAt:    time.Now(),
	}

	if err := a.storage.SaveReminder(reminder); err != nil {
		return fmt.Errorf("failed to save reminder: %w", err)
	}

	fmt.Printf("reminder triggered for user %s: %s\n", req.UserID, req.ReminderText)
	return nil
}
