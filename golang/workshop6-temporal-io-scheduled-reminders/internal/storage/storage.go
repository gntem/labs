package storage

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

type Reminder struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	ReminderText string    `json:"reminder_text"`
	ReminderTime int64     `json:"reminder_time"`
	CreatedAt    time.Time `json:"created_at"`
}

type Storage struct {
	filename string
	mu       sync.RWMutex
}

func NewStorage(filename string) *Storage {
	return &Storage{
		filename: filename,
	}
}

func (s *Storage) SaveReminder(reminder *Reminder) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	reminders, err := s.loadReminders()
	if err != nil {
		return err
	}

	reminders = append(reminders, *reminder)
	return s.saveReminders(reminders)
}

func (s *Storage) GetReminders(userID string) ([]Reminder, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	reminders, err := s.loadReminders()
	if err != nil {
		return nil, err
	}

	var userReminders []Reminder
	for _, reminder := range reminders {
		if reminder.UserID == userID {
			userReminders = append(userReminders, reminder)
		}
	}

	return userReminders, nil
}

func (s *Storage) loadReminders() ([]Reminder, error) {
	if _, err := os.Stat(s.filename); os.IsNotExist(err) {
		return []Reminder{}, nil
	}

	data, err := os.ReadFile(s.filename)
	if err != nil {
		return nil, err
	}

	var reminders []Reminder
	if len(data) > 0 {
		err = json.Unmarshal(data, &reminders)
		if err != nil {
			return nil, err
		}
	}

	return reminders, nil
}

func (s *Storage) saveReminders(reminders []Reminder) error {
	data, err := json.MarshalIndent(reminders, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filename, data, 0644)
}
