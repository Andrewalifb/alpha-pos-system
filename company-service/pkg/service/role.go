package service

import (
	"context"

	pb "github.com/Andrewalifb/alpha-pos-system/company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/company-service/pkg/repository"
)

type PosRoleService interface {
	CreatePosRole(ctx context.Context, req *pb.CreatePosRoleRequest) (*pb.CreatePosRoleResponse, error)
	ReadPosRole(ctx context.Context, req *pb.ReadPosRoleRequest) (*pb.ReadPosRoleResponse, error)
	UpdatePosRole(ctx context.Context, req *pb.UpdatePosRoleRequest) (*pb.UpdatePosRoleResponse, error)
	DeletePosRole(ctx context.Context, req *pb.DeletePosRoleRequest) (*pb.DeletePosRoleResponse, error)
	ReadAllPosRoles(ctx context.Context, req *pb.ReadAllPosRolesRequest) (*pb.ReadAllPosRolesResponse, error)
}

type PosRoleServiceServer struct {
	pb.UnimplementedPosRoleServiceServer
	repo repository.PosRoleRepository
}

func NewPosRoleServiceServer(repo repository.PosRoleRepository) *PosRoleServiceServer {
	return &PosRoleServiceServer{
		repo: repo,
	}
}
