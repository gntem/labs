package repository

import "workshop4-grpc-leaderboard/model"

type LeaderboardRepository interface {
	ListLeaderboards(pageSize, pageNumber int32) ([]*model.Leaderboard, int32, error)
}
