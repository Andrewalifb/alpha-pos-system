package service

import (
	"context"

	pb "github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/loyalty-service/pkg/repository"
)

type PosCustomerLoyaltyService interface {
	CreatePosCustomerLoyalty(ctx context.Context, req *pb.CreatePosCustomerLoyaltyRequest) (*pb.CreatePosCustomerLoyaltyResponse, error)
	ReadPosCustomerLoyalty(ctx context.Context, req *pb.ReadPosCustomerLoyaltyRequest) (*pb.ReadPosCustomerLoyaltyResponse, error)
	UpdatePosCustomerLoyalty(ctx context.Context, req *pb.UpdatePosCustomerLoyaltyRequest) (*pb.UpdatePosCustomerLoyaltyResponse, error)
	DeletePosCustomerLoyalty(ctx context.Context, req *pb.DeletePosCustomerLoyaltyRequest) (*pb.DeletePosCustomerLoyaltyResponse, error)
	ReadAllPosCustomerLoyalties(ctx context.Context, req *pb.ReadAllPosCustomerLoyaltiesRequest) (*pb.ReadAllPosCustomerLoyaltiesResponse, error)
}

type posCustomerLoyaltyService struct {
	pb.UnimplementedPosCustomerLoyaltyServiceServer
	repo repository.PosCustomerLoyaltyRepository
}

func NewPosCustomerLoyaltyService(repo repository.PosCustomerLoyaltyRepository) *posCustomerLoyaltyService {
	return &posCustomerLoyaltyService{
		repo: repo,
	}
}
