package workflows

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

type ReminderRequest struct {
	ID           string
	UserID       string
	ReminderText string
	ReminderTime int64
}

func RemindmeWorkflow(ctx workflow.Context, req ReminderRequest) error {
	reminderTime := time.Unix(req.ReminderTime, 0)
	currentTime := workflow.Now(ctx)

	if reminderTime.After(currentTime) {
		timer := workflow.NewTimer(ctx, reminderTime.Sub(currentTime))
		if err := timer.Get(ctx, nil); err != nil {
			return err
		}
	}

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	return workflow.ExecuteActivity(ctx, "RemindActivity", req).Get(ctx, nil)
}
