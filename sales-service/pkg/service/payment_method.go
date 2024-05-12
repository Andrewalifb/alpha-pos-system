package service

import (
	"context"

	pb "github.com/Andrewalifb/alpha-pos-system/sales-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/sales-service/pkg/repository"
)

type PosPaymentMethodService interface {
	CreatePosPaymentMethod(ctx context.Context, req *pb.CreatePosPaymentMethodRequest) (*pb.CreatePosPaymentMethodResponse, error)
	ReadPosPaymentMethod(ctx context.Context, req *pb.ReadPosPaymentMethodRequest) (*pb.ReadPosPaymentMethodResponse, error)
	UpdatePosPaymentMethod(ctx context.Context, req *pb.UpdatePosPaymentMethodRequest) (*pb.UpdatePosPaymentMethodResponse, error)
	DeletePosPaymentMethod(ctx context.Context, req *pb.DeletePosPaymentMethodRequest) (*pb.DeletePosPaymentMethodResponse, error)
	ReadAllPosPaymentMethods(ctx context.Context, req *pb.ReadAllPosPaymentMethodsRequest) (*pb.ReadAllPosPaymentMethodsResponse, error)
}

type posPaymentMethodService struct {
	pb.UnimplementedPosPaymentMethodServiceServer
	repo repository.PosPaymentMethodRepository
}

func NewPosPaymentMethodService(repo repository.PosPaymentMethodRepository) *posPaymentMethodService {
	return &posPaymentMethodService{
		repo: repo,
	}
}
