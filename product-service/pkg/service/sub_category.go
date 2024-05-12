package service

import (
	"context"

	pb "github.com/Andrewalifb/alpha-pos-system/product-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/product-service/pkg/repository"
)

type PosProductSubCategoryService interface {
	CreatePosProductSubCategory(ctx context.Context, req *pb.CreatePosProductSubCategoryRequest) (*pb.CreatePosProductSubCategoryResponse, error)
	ReadPosProductSubCategory(ctx context.Context, req *pb.ReadPosProductSubCategoryRequest) (*pb.ReadPosProductSubCategoryResponse, error)
	UpdatePosProductSubCategory(ctx context.Context, req *pb.UpdatePosProductSubCategoryRequest) (*pb.UpdatePosProductSubCategoryResponse, error)
	DeletePosProductSubCategory(ctx context.Context, req *pb.DeletePosProductSubCategoryRequest) (*pb.DeletePosProductSubCategoryResponse, error)
	ReadAllPosProductSubCategories(ctx context.Context, req *pb.ReadAllPosProductSubCategoriesRequest) (*pb.ReadAllPosProductSubCategoriesResponse, error)
}

type posProductSubCategoryService struct {
	pb.UnimplementedPosProductSubCategoryServiceServer
	repo repository.PosProductSubCategoryRepository
}

func NewPosProductSubCategoryService(repo repository.PosProductSubCategoryRepository) *posProductSubCategoryService {
	return &posProductSubCategoryService{
		repo: repo,
	}
}
