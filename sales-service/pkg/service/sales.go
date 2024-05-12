package service

import (
	"context"

	pb "github.com/Andrewalifb/alpha-pos-system/sales-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/sales-service/pkg/repository"
)

type PosSaleService interface {
	CreatePosSale(ctx context.Context, req *pb.CreatePosSaleRequest) (*pb.CreatePosSaleResponse, error)
	ReadPosSale(ctx context.Context, req *pb.ReadPosSaleRequest) (*pb.ReadPosSaleResponse, error)
	UpdatePosSale(ctx context.Context, req *pb.UpdatePosSaleRequest) (*pb.UpdatePosSaleResponse, error)
	DeletePosSale(ctx context.Context, req *pb.DeletePosSaleRequest) (*pb.DeletePosSaleResponse, error)
	ReadAllPosSales(ctx context.Context, req *pb.ReadAllPosSalesRequest) (*pb.ReadAllPosSalesResponse, error)
}

type posSaleService struct {
	pb.UnimplementedPosSaleServiceServer
	repo repository.PosSaleRepository
}

func NewPosSaleService(repo repository.PosSaleRepository) *posSaleService {
	return &posSaleService{
		repo: repo,
	}
}
