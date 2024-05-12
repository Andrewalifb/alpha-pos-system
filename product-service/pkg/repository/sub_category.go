package repository

import (
	pb "github.com/Andrewalifb/alpha-pos-system/product-service/api/proto"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type PosProductSubCategoryRepository interface {
	CreatePosProductSubCategory(posProductSubCategory *pb.PosProductSubCategory) error
	ReadPosProductSubCategory(subCategoryID string) (*pb.PosProductSubCategory, error)
	UpdatePosProductSubCategory(posProductSubCategory *pb.PosProductSubCategory) error
	DeletePosProductSubCategory(subCategoryID string) error
	ReadAllPosProductSubCategories() ([]*pb.PosProductSubCategory, error)
}

type posProductSubCategoryRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosProductSubCategoryRepository(db *gorm.DB, redis *redis.Client) PosProductSubCategoryRepository {
	return &posProductSubCategoryRepository{
		db:    db,
		redis: redis,
	}
}
