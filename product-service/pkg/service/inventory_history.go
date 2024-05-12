package service

import (
	"context"

	pb "github.com/Andrewalifb/alpha-pos-system/product-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/product-service/pkg/repository"
)

type PosInventoryHistoryService interface {
	CreatePosInventoryHistory(ctx context.Context, req *pb.CreatePosInventoryHistoryRequest) (*pb.CreatePosInventoryHistoryResponse, error)
	ReadPosInventoryHistory(ctx context.Context, req *pb.ReadPosInventoryHistoryRequest) (*pb.ReadPosInventoryHistoryResponse, error)
	UpdatePosInventoryHistory(ctx context.Context, req *pb.UpdatePosInventoryHistoryRequest) (*pb.UpdatePosInventoryHistoryResponse, error)
	DeletePosInventoryHistory(ctx context.Context, req *pb.DeletePosInventoryHistoryRequest) (*pb.DeletePosInventoryHistoryResponse, error)
	ReadAllPosInventoryHistories(ctx context.Context, req *pb.ReadAllPosInventoryHistoriesRequest) (*pb.ReadAllPosInventoryHistoriesResponse, error)
}

type posInventoryHistoryService struct {
	pb.UnimplementedPosInventoryHistoryServiceServer
	repo repository.PosInventoryHistoryRepository
}

func NewPosInventoryHistoryService(repo repository.PosInventoryHistoryRepository) *posInventoryHistoryService {
	return &posInventoryHistoryService{
		repo: repo,
	}
}
