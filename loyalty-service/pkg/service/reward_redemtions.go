package service

import (
	"context"

	pb "github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/loyalty-service/pkg/repository"
)

type PosRewardRedemptionService interface {
	CreatePosRewardRedemption(ctx context.Context, req *pb.CreatePosRewardRedemptionRequest) (*pb.CreatePosRewardRedemptionResponse, error)
	ReadPosRewardRedemption(ctx context.Context, req *pb.ReadPosRewardRedemptionRequest) (*pb.ReadPosRewardRedemptionResponse, error)
	UpdatePosRewardRedemption(ctx context.Context, req *pb.UpdatePosRewardRedemptionRequest) (*pb.UpdatePosRewardRedemptionResponse, error)
	DeletePosRewardRedemption(ctx context.Context, req *pb.DeletePosRewardRedemptionRequest) (*pb.DeletePosRewardRedemptionResponse, error)
	ReadAllPosRewardRedemptions(ctx context.Context, req *pb.ReadAllPosRewardRedemptionsRequest) (*pb.ReadAllPosRewardRedemptionsResponse, error)
}

type posRewardRedemptionService struct {
	pb.UnimplementedPosRewardRedemptionServiceServer
	repo repository.PosRewardRedemptionRepository
}

func NewPosRewardRedemptionService(repo repository.PosRewardRedemptionRepository) *posRewardRedemptionService {
	return &posRewardRedemptionService{
		repo: repo,
	}
}
