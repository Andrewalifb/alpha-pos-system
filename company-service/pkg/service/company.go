// Service
package service

import (
	"context"

	pb "github.com/Andrewalifb/alpha-pos-system/company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/company-service/pkg/repository"
)

type PosCompanyService interface {
	CreatePosCompany(ctx context.Context, req *pb.CreatePosCompanyRequest) (*pb.CreatePosCompanyResponse, error)
	ReadPosCompany(ctx context.Context, req *pb.ReadPosCompanyRequest) (*pb.ReadPosCompanyResponse, error)
	UpdatePosCompany(ctx context.Context, req *pb.UpdatePosCompanyRequest) (*pb.UpdatePosCompanyResponse, error)
	DeletePosCompany(ctx context.Context, req *pb.DeletePosCompanyRequest) (*pb.DeletePosCompanyResponse, error)
	ReadAllPosCompanies(ctx context.Context, req *pb.ReadAllPosCompaniesRequest) (*pb.ReadAllPosCompaniesResponse, error)
}

type PosCompanyServiceServer struct {
	pb.UnimplementedPosCompanyServiceServer
	repo repository.PosCompanyRepository
}

func NewPosCompanyServiceServer(repo repository.PosCompanyRepository) *PosCompanyServiceServer {
	return &PosCompanyServiceServer{
		repo: repo,
	}
}
