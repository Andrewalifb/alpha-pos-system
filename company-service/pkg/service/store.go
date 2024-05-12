package service

import (
	"context"

	pb "github.com/Andrewalifb/alpha-pos-system/company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/company-service/pkg/repository"
)

type PosStoreService interface {
	CreatePosStore(ctx context.Context, req *pb.CreatePosStoreRequest) (*pb.CreatePosStoreResponse, error)
	ReadPosStore(ctx context.Context, req *pb.ReadPosStoreRequest) (*pb.ReadPosStoreResponse, error)
	UpdatePosStore(ctx context.Context, req *pb.UpdatePosStoreRequest) (*pb.UpdatePosStoreResponse, error)
	DeletePosStore(ctx context.Context, req *pb.DeletePosStoreRequest) (*pb.DeletePosStoreResponse, error)
	ReadAllPosStores(ctx context.Context, req *pb.ReadAllPosStoresRequest) (*pb.ReadAllPosStoresResponse, error)
}

type PosStoreServiceServer struct {
	pb.UnimplementedPosStoreServiceServer
	repo repository.PosStoreRepository
}

func NewPosStoreServiceServer(repo repository.PosStoreRepository) *PosStoreServiceServer {
	return &PosStoreServiceServer{
		repo: repo,
	}
}
