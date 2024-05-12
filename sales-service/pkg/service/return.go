package service

import (
	"context"

	pb "github.com/Andrewalifb/alpha-pos-system/sales-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/sales-service/pkg/repository"
)

type PosReturnService interface {
	CreatePosReturn(ctx context.Context, req *pb.CreatePosReturnRequest) (*pb.CreatePosReturnResponse, error)
	ReadPosReturn(ctx context.Context, req *pb.ReadPosReturnRequest) (*pb.ReadPosReturnResponse, error)
	UpdatePosReturn(ctx context.Context, req *pb.UpdatePosReturnRequest) (*pb.UpdatePosReturnResponse, error)
	DeletePosReturn(ctx context.Context, req *pb.DeletePosReturnRequest) (*pb.DeletePosReturnResponse, error)
	ReadAllPosReturns(ctx context.Context, req *pb.ReadAllPosReturnsRequest) (*pb.ReadAllPosReturnsResponse, error)
}

type posReturnService struct {
	pb.UnimplementedPosReturnServiceServer
	repo repository.PosReturnRepository
}

func NewPosReturnService(repo repository.PosReturnRepository) *posReturnService {
	return &posReturnService{
		repo: repo,
	}
}
