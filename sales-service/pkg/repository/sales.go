package repository

import (
	pb "github.com/Andrewalifb/alpha-pos-system/sales-service/api/proto"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type PosSaleRepository interface {
	CreatePosSale(posSale *pb.PosSale) error
	ReadPosSale(saleID string) (*pb.PosSale, error)
	UpdatePosSale(posSale *pb.PosSale) error
	DeletePosSale(saleID string) error
	ReadAllPosSales() ([]*pb.PosSale, error)
}

type posSaleRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosSaleRepository(db *gorm.DB, redis *redis.Client) PosSaleRepository {
	return &posSaleRepository{
		db:    db,
		redis: redis,
	}
}
