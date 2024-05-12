package repository

import (
	pb "github.com/Andrewalifb/alpha-pos-system/product-service/api/proto"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type PosProductCategoryRepository interface {
	CreatePosProductCategory(posProductCategory *pb.PosProductCategory) error
	ReadPosProductCategory(categoryID string) (*pb.PosProductCategory, error)
	UpdatePosProductCategory(posProductCategory *pb.PosProductCategory) error
	DeletePosProductCategory(categoryID string) error
	ReadAllPosProductCategories() ([]*pb.PosProductCategory, error)
}

type posProductCategoryRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosProductCategoryRepository(db *gorm.DB, redis *redis.Client) PosProductCategoryRepository {
	return &posProductCategoryRepository{
		db:    db,
		redis: redis,
	}
}
