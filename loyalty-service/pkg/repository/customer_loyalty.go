package repository

import (
	pb "github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/proto"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type PosCustomerLoyaltyRepository interface {
	CreatePosCustomerLoyalty(posCustomerLoyalty *pb.PosCustomerLoyalty) error
	ReadPosCustomerLoyalty(customerID string) (*pb.PosCustomerLoyalty, error)
	UpdatePosCustomerLoyalty(posCustomerLoyalty *pb.PosCustomerLoyalty) error
	DeletePosCustomerLoyalty(customerID string) error
	ReadAllPosCustomerLoyalties() ([]*pb.PosCustomerLoyalty, error)
}

type posCustomerLoyaltyRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosCustomerLoyaltyRepository(db *gorm.DB, redis *redis.Client) PosCustomerLoyaltyRepository {
	return &posCustomerLoyaltyRepository{
		db:    db,
		redis: redis,
	}
}
