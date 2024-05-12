package service

import (
	"context"

	pb "github.com/Andrewalifb/alpha-pos-system/product-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/product-service/pkg/repository"
)

type PosProductCategoryService interface {
	CreatePosProductCategory(ctx context.Context, req *pb.CreatePosProductCategoryRequest) (*pb.CreatePosProductCategoryResponse, error)
	ReadPosProductCategory(ctx context.Context, req *pb.ReadPosProductCategoryRequest) (*pb.ReadPosProductCategoryResponse, error)
	UpdatePosProductCategory(ctx context.Context, req *pb.UpdatePosProductCategoryRequest) (*pb.UpdatePosProductCategoryResponse, error)
	DeletePosProductCategory(ctx context.Context, req *pb.DeletePosProductCategoryRequest) (*pb.DeletePosProductCategoryResponse, error)
	ReadAllPosProductCategories(ctx context.Context, req *pb.ReadAllPosProductCategoriesRequest) (*pb.ReadAllPosProductCategoriesResponse, error)
}

type posProductCategoryService struct {
	pb.UnimplementedPosProductCategoryServiceServer
	repo repository.PosProductCategoryRepository
}

func NewPosProductCategoryService(repo repository.PosProductCategoryRepository) *posProductCategoryService {
	return &posProductCategoryService{
		repo: repo,
	}
}
