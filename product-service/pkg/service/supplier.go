package service

import (
	"context"

	pb "github.com/Andrewalifb/alpha-pos-system/product-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/product-service/pkg/repository"
)

type PosSupplierService interface {
	CreatePosSupplier(ctx context.Context, req *pb.CreatePosSupplierRequest) (*pb.CreatePosSupplierResponse, error)
	ReadPosSupplier(ctx context.Context, req *pb.ReadPosSupplierRequest) (*pb.ReadPosSupplierResponse, error)
	UpdatePosSupplier(ctx context.Context, req *pb.UpdatePosSupplierRequest) (*pb.UpdatePosSupplierResponse, error)
	DeletePosSupplier(ctx context.Context, req *pb.DeletePosSupplierRequest) (*pb.DeletePosSupplierResponse, error)
	ReadAllPosSuppliers(ctx context.Context, req *pb.ReadAllPosSuppliersRequest) (*pb.ReadAllPosSuppliersResponse, error)
}

type posSupplierService struct {
	pb.UnimplementedPosSupplierServiceServer
	repo repository.PosSupplierRepository
}

func NewPosSupplierService(repo repository.PosSupplierRepository) *posSupplierService {
	return &posSupplierService{
		repo: repo,
	}
}
