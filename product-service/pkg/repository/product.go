package repository

import (
	pb "github.com/Andrewalifb/alpha-pos-system/product-service/api/proto"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type PosProductRepository interface {
	CreatePosProduct(posProduct *pb.PosProduct) error
	ReadPosProduct(productID string) (*pb.PosProduct, error)
	UpdatePosProduct(posProduct *pb.PosProduct) error
	DeletePosProduct(productID string) error
	ReadAllPosProducts() ([]*pb.PosProduct, error)
}

type posProductRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosProductRepository(db *gorm.DB, redis *redis.Client) PosProductRepository {
	return &posProductRepository{
		db:    db,
		redis: redis,
	}
}
