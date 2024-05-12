package service

import (
	"context"

	pb "github.com/Andrewalifb/alpha-pos-system/product-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/product-service/pkg/repository"
)

type PosProductService interface {
	CreatePosProduct(ctx context.Context, req *pb.CreatePosProductRequest) (*pb.CreatePosProductResponse, error)
	ReadPosProduct(ctx context.Context, req *pb.ReadPosProductRequest) (*pb.ReadPosProductResponse, error)
	UpdatePosProduct(ctx context.Context, req *pb.UpdatePosProductRequest) (*pb.UpdatePosProductResponse, error)
	DeletePosProduct(ctx context.Context, req *pb.DeletePosProductRequest) (*pb.DeletePosProductResponse, error)
	ReadAllPosProducts(ctx context.Context, req *pb.ReadAllPosProductsRequest) (*pb.ReadAllPosProductsResponse, error)
}

type posProductService struct {
	pb.UnimplementedPosProductServiceServer
	repo repository.PosProductRepository
}

func NewPosProductService(repo repository.PosProductRepository) *posProductService {
	return &posProductService{
		repo: repo,
	}
}
