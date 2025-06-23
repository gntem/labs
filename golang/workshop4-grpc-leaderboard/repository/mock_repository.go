package repository

import (
	"workshop4-grpc-leaderboard/model"
)

type MockLeaderboardRepository struct {
	leaderboards []*model.Leaderboard
}

func NewMockLeaderboardRepository() *MockLeaderboardRepository {
	return &MockLeaderboardRepository{
		leaderboards: []*model.Leaderboard{
			{ID: "1", UserID: "user1", Name: "Global Rankings", Score: 1000},
			{ID: "2", UserID: "user2", Name: "Weekly Challenge", Score: 850},
			{ID: "3", UserID: "user3", Name: "Monthly Tournament", Score: 1200},
			{ID: "4", UserID: "user4", Name: "Boss Battle", Score: 750},
			{ID: "5", UserID: "user5", Name: "Speedrun", Score: 950},
		},
	}
}

func (r *MockLeaderboardRepository) ListLeaderboards(pageSize, pageNumber int32) ([]*model.Leaderboard, int32, error) {
	if pageSize <= 0 {
		pageSize = 10
	}

	if pageNumber <= 0 {
		pageNumber = 1
	}

	totalCount := int32(len(r.leaderboards))
	start := (pageNumber - 1) * pageSize
	end := start + pageSize

	if start >= totalCount {
		return []*model.Leaderboard{}, totalCount, nil
	}

	if end > totalCount {
		end = totalCount
	}

	return r.leaderboards[start:end], totalCount, nil
}
