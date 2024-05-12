package service

import (
	"context"

	pb "github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/loyalty-service/pkg/repository"
)

type PosPointsTransactionService interface {
	CreatePosPointsTransaction(ctx context.Context, req *pb.CreatePosPointsTransactionRequest) (*pb.CreatePosPointsTransactionResponse, error)
	ReadPosPointsTransaction(ctx context.Context, req *pb.ReadPosPointsTransactionRequest) (*pb.ReadPosPointsTransactionResponse, error)
	UpdatePosPointsTransaction(ctx context.Context, req *pb.UpdatePosPointsTransactionRequest) (*pb.UpdatePosPointsTransactionResponse, error)
	DeletePosPointsTransaction(ctx context.Context, req *pb.DeletePosPointsTransactionRequest) (*pb.DeletePosPointsTransactionResponse, error)
	ReadAllPosPointsTransactions(ctx context.Context, req *pb.ReadAllPosPointsTransactionsRequest) (*pb.ReadAllPosPointsTransactionsResponse, error)
}

type posPointsTransactionService struct {
	pb.UnimplementedPosPointsTransactionServiceServer
	repo repository.PosPointsTransactionRepository
}

func NewPosPointsTransactionService(repo repository.PosPointsTransactionRepository) *posPointsTransactionService {
	return &posPointsTransactionService{
		repo: repo,
	}
}
