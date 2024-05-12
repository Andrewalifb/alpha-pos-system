package service

import (
	"context"

	pb "github.com/Andrewalifb/alpha-pos-system/sales-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/sales-service/pkg/repository"
)

type PosOnlinePaymentService interface {
	CreatePosOnlinePayment(ctx context.Context, req *pb.CreatePosOnlinePaymentRequest) (*pb.CreatePosOnlinePaymentResponse, error)
	ReadPosOnlinePayment(ctx context.Context, req *pb.ReadPosOnlinePaymentRequest) (*pb.ReadPosOnlinePaymentResponse, error)
	UpdatePosOnlinePayment(ctx context.Context, req *pb.UpdatePosOnlinePaymentRequest) (*pb.UpdatePosOnlinePaymentResponse, error)
	DeletePosOnlinePayment(ctx context.Context, req *pb.DeletePosOnlinePaymentRequest) (*pb.DeletePosOnlinePaymentResponse, error)
	ReadAllPosOnlinePayments(ctx context.Context, req *pb.ReadAllPosOnlinePaymentsRequest) (*pb.ReadAllPosOnlinePaymentsResponse, error)
}

type posOnlinePaymentService struct {
	pb.UnimplementedPosOnlinePaymentServiceServer
	repo repository.PosOnlinePaymentRepository
}

func NewPosOnlinePaymentService(repo repository.PosOnlinePaymentRepository) *posOnlinePaymentService {
	return &posOnlinePaymentService{
		repo: repo,
	}
}
