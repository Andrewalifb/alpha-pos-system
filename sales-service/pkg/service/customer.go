package service

import (
	"context"

	pb "github.com/Andrewalifb/alpha-pos-system/sales-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/sales-service/pkg/repository"
)

type PosCustomerService interface {
	CreatePosCustomer(ctx context.Context, req *pb.CreatePosCustomerRequest) (*pb.CreatePosCustomerResponse, error)
	ReadPosCustomer(ctx context.Context, req *pb.ReadPosCustomerRequest) (*pb.ReadPosCustomerResponse, error)
	UpdatePosCustomer(ctx context.Context, req *pb.UpdatePosCustomerRequest) (*pb.UpdatePosCustomerResponse, error)
	DeletePosCustomer(ctx context.Context, req *pb.DeletePosCustomerRequest) (*pb.DeletePosCustomerResponse, error)
	ReadAllPosCustomers(ctx context.Context, req *pb.ReadAllPosCustomersRequest) (*pb.ReadAllPosCustomersResponse, error)
}

type posCustomerService struct {
	pb.UnimplementedPosCustomerServiceServer
	repo repository.PosCustomerRepository
}

func NewPosCustomerService(repo repository.PosCustomerRepository) *posCustomerService {
	return &posCustomerService{
		repo: repo,
	}
}
