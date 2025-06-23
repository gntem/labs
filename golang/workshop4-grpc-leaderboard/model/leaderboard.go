package model

type Leaderboard struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Score  int64  `json:"score"`
}
