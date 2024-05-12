package repository

import (
	pb "github.com/Andrewalifb/alpha-pos-system/product-service/api/proto"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type PosInventoryHistoryRepository interface {
	CreatePosInventoryHistory(posInventoryHistory *pb.PosInventoryHistory) error
	ReadPosInventoryHistory(inventoryID string) (*pb.PosInventoryHistory, error)
	UpdatePosInventoryHistory(posInventoryHistory *pb.PosInventoryHistory) error
	DeletePosInventoryHistory(inventoryID string) error
	ReadAllPosInventoryHistories() ([]*pb.PosInventoryHistory, error)
}

type posInventoryHistoryRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosInventoryHistoryRepository(db *gorm.DB, redis *redis.Client) PosInventoryHistoryRepository {
	return &posInventoryHistoryRepository{
		db:    db,
		redis: redis,
	}
}
