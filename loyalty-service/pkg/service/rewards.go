package service

import (
	"context"

	pb "github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/loyalty-service/pkg/repository"
)

type PosRewardService interface {
	CreatePosReward(ctx context.Context, req *pb.CreatePosRewardRequest) (*pb.CreatePosRewardResponse, error)
	ReadPosReward(ctx context.Context, req *pb.ReadPosRewardRequest) (*pb.ReadPosRewardResponse, error)
	UpdatePosReward(ctx context.Context, req *pb.UpdatePosRewardRequest) (*pb.UpdatePosRewardResponse, error)
	DeletePosReward(ctx context.Context, req *pb.DeletePosRewardRequest) (*pb.DeletePosRewardResponse, error)
	ReadAllPosRewards(ctx context.Context, req *pb.ReadAllPosRewardsRequest) (*pb.ReadAllPosRewardsResponse, error)
}

type posRewardService struct {
	pb.UnimplementedPosRewardServiceServer
	repo repository.PosRewardRepository
}

func NewPosRewardService(repo repository.PosRewardRepository) *posRewardService {
	return &posRewardService{
		repo: repo,
	}
}
