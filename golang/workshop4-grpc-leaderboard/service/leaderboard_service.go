package service

import (
	"context"

	"workshop4-grpc-leaderboard/proto"
	"workshop4-grpc-leaderboard/repository"
)

type LeaderboardService struct {
	proto.UnimplementedLeaderboardServiceServer
	repo repository.LeaderboardRepository
}

func NewLeaderboardService(repo repository.LeaderboardRepository) *LeaderboardService {
	return &LeaderboardService{
		repo: repo,
	}
}

func (s *LeaderboardService) ListLeaderboards(ctx context.Context, req *proto.ListLeaderboardsRequest) (*proto.ListLeaderboardsResponse, error) {
	leaderboards, totalCount, err := s.repo.ListLeaderboards(req.PageSize, req.PageNumber)
	if err != nil {
		return nil, err
	}

	protoLeaderboards := make([]*proto.Leaderboard, len(leaderboards))
	for i, lb := range leaderboards {
		protoLeaderboards[i] = &proto.Leaderboard{
			Id:     lb.ID,
			UserId: lb.UserID,
			Name:   lb.Name,
			Score:  lb.Score,
		}
	}

	return &proto.ListLeaderboardsResponse{
		Leaderboards: protoLeaderboards,
		TotalCount:   totalCount,
		PageNumber:   req.PageNumber,
	}, nil
}
