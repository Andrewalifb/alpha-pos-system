package repository

import (
	pb "github.com/Andrewalifb/alpha-pos-system/product-service/api/proto"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type PosSupplierRepository interface {
	CreatePosSupplier(posSupplier *pb.PosSupplier) error
	ReadPosSupplier(supplierID string) (*pb.PosSupplier, error)
	UpdatePosSupplier(posSupplier *pb.PosSupplier) error
	DeletePosSupplier(supplierID string) error
	ReadAllPosSuppliers() ([]*pb.PosSupplier, error)
}

type posSupplierRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosSupplierRepository(db *gorm.DB, redis *redis.Client) PosSupplierRepository {
	return &posSupplierRepository{
		db:    db,
		redis: redis,
	}
}
