package service

import (
	"context"

	pb "github.com/Andrewalifb/alpha-pos-system/sales-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/sales-service/pkg/repository"
)

type PosCashDrawerService interface {
	CreatePosCashDrawer(ctx context.Context, req *pb.CreatePosCashDrawerRequest) (*pb.CreatePosCashDrawerResponse, error)
	ReadPosCashDrawer(ctx context.Context, req *pb.ReadPosCashDrawerRequest) (*pb.ReadPosCashDrawerResponse, error)
	UpdatePosCashDrawer(ctx context.Context, req *pb.UpdatePosCashDrawerRequest) (*pb.UpdatePosCashDrawerResponse, error)
	DeletePosCashDrawer(ctx context.Context, req *pb.DeletePosCashDrawerRequest) (*pb.DeletePosCashDrawerResponse, error)
	ReadAllPosCashDrawers(ctx context.Context, req *pb.ReadAllPosCashDrawersRequest) (*pb.ReadAllPosCashDrawersResponse, error)
}

type posCashDrawerService struct {
	pb.UnimplementedPosCashDrawerServiceServer
	repo repository.PosCashDrawerRepository
}

func NewPosCashDrawerService(repo repository.PosCashDrawerRepository) *posCashDrawerService {
	return &posCashDrawerService{
		repo: repo,
	}
}
