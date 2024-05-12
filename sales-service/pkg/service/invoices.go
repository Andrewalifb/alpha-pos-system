package service

import (
	"context"

	pb "github.com/Andrewalifb/alpha-pos-system/sales-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/sales-service/pkg/repository"
)

type PosInvoiceService interface {
	CreatePosInvoice(ctx context.Context, req *pb.CreatePosInvoiceRequest) (*pb.CreatePosInvoiceResponse, error)
	ReadPosInvoice(ctx context.Context, req *pb.ReadPosInvoiceRequest) (*pb.ReadPosInvoiceResponse, error)
	UpdatePosInvoice(ctx context.Context, req *pb.UpdatePosInvoiceRequest) (*pb.UpdatePosInvoiceResponse, error)
	DeletePosInvoice(ctx context.Context, req *pb.DeletePosInvoiceRequest) (*pb.DeletePosInvoiceResponse, error)
	ReadAllPosInvoices(ctx context.Context, req *pb.ReadAllPosInvoicesRequest) (*pb.ReadAllPosInvoicesResponse, error)
}

type posInvoiceService struct {
	pb.UnimplementedPosInvoiceServiceServer
	repo repository.PosInvoiceRepository
}

func NewPosInvoiceService(repo repository.PosInvoiceRepository) *posInvoiceService {
	return &posInvoiceService{
		repo: repo,
	}
}
