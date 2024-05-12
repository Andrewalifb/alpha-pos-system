package service

import (
	"context"

	pb "github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/loyalty-service/pkg/repository"
)

type PosLoyaltyLevelService interface {
	CreatePosLoyaltyLevel(ctx context.Context, req *pb.CreatePosLoyaltyLevelRequest) (*pb.CreatePosLoyaltyLevelResponse, error)
	ReadPosLoyaltyLevel(ctx context.Context, req *pb.ReadPosLoyaltyLevelRequest) (*pb.ReadPosLoyaltyLevelResponse, error)
	UpdatePosLoyaltyLevel(ctx context.Context, req *pb.UpdatePosLoyaltyLevelRequest) (*pb.UpdatePosLoyaltyLevelResponse, error)
	DeletePosLoyaltyLevel(ctx context.Context, req *pb.DeletePosLoyaltyLevelRequest) (*pb.DeletePosLoyaltyLevelResponse, error)
	ReadAllPosLoyaltyLevels(ctx context.Context, req *pb.ReadAllPosLoyaltyLevelsRequest) (*pb.ReadAllPosLoyaltyLevelsResponse, error)
}

type posLoyaltyLevelService struct {
	pb.UnimplementedPosLoyaltyLevelServiceServer
	repo repository.PosLoyaltyLevelRepository
}

func NewPosLoyaltyLevelService(repo repository.PosLoyaltyLevelRepository) *posLoyaltyLevelService {
	return &posLoyaltyLevelService{
		repo: repo,
	}
}
