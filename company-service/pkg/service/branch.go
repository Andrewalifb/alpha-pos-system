package service

import (
	"context"

	pb "github.com/Andrewalifb/alpha-pos-system/company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/company-service/pkg/repository"
)

type PosStoreBranchService interface {
	CreatePosStoreBranch(ctx context.Context, req *pb.CreatePosStoreBranchRequest) (*pb.CreatePosStoreBranchResponse, error)
	ReadPosStoreBranch(ctx context.Context, req *pb.ReadPosStoreBranchRequest) (*pb.ReadPosStoreBranchResponse, error)
	UpdatePosStoreBranch(ctx context.Context, req *pb.UpdatePosStoreBranchRequest) (*pb.UpdatePosStoreBranchResponse, error)
	DeletePosStoreBranch(ctx context.Context, req *pb.DeletePosStoreBranchRequest) (*pb.DeletePosStoreBranchResponse, error)
	ReadAllPosStoreBranches(ctx context.Context, req *pb.ReadAllPosStoreBranchesRequest) (*pb.ReadAllPosStoreBranchesResponse, error)
}

type PosStoreBranchServiceServer struct {
	pb.UnimplementedPosStoreBranchServiceServer
	repo repository.PosStoreBranchRepository
}

func NewPosStoreBranchServiceServer(repo repository.PosStoreBranchRepository) *PosStoreBranchServiceServer {
	return &PosStoreBranchServiceServer{
		repo: repo,
	}
}
