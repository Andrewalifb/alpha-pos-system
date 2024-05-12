package repository

import (
	pb "github.com/Andrewalifb/alpha-pos-system/company-service/api/proto"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type PosStoreRepository interface {
	CreatePosStore(posStore *pb.PosStore) error
	ReadPosStore(storeID string) (*pb.PosStore, error)
	UpdatePosStore(posStore *pb.PosStore) error
	DeletePosStore(storeID string) error
	ReadAllPosStores() ([]*pb.PosStore, error)
}

type posStoreRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosStoreRepository(db *gorm.DB, redis *redis.Client) PosStoreRepository {
	return &posStoreRepository{
		db:    db,
		redis: redis,
	}
}
