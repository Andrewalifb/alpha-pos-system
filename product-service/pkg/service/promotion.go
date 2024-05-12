package service

import (
	"context"

	pb "github.com/Andrewalifb/alpha-pos-system/product-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/product-service/pkg/repository"
)

type PosPromotionService interface {
	CreatePosPromotion(ctx context.Context, req *pb.CreatePosPromotionRequest) (*pb.CreatePosPromotionResponse, error)
	ReadPosPromotion(ctx context.Context, req *pb.ReadPosPromotionRequest) (*pb.ReadPosPromotionResponse, error)
	UpdatePosPromotion(ctx context.Context, req *pb.UpdatePosPromotionRequest) (*pb.UpdatePosPromotionResponse, error)
	DeletePosPromotion(ctx context.Context, req *pb.DeletePosPromotionRequest) (*pb.DeletePosPromotionResponse, error)
	ReadAllPosPromotions(ctx context.Context, req *pb.ReadAllPosPromotionsRequest) (*pb.ReadAllPosPromotionsResponse, error)
}

type posPromotionService struct {
	pb.UnimplementedPosPromotionServiceServer
	repo repository.PosPromotionRepository
}

func NewPosPromotionService(repo repository.PosPromotionRepository) *posPromotionService {
	return &posPromotionService{
		repo: repo,
	}
}
