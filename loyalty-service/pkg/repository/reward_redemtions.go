package repository

import (
	pb "github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/proto"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type PosRewardRedemptionRepository interface {
	CreatePosRewardRedemption(posRewardRedemption *pb.PosRewardRedemption) error
	ReadPosRewardRedemption(redemptionID string) (*pb.PosRewardRedemption, error)
	UpdatePosRewardRedemption(posRewardRedemption *pb.PosRewardRedemption) error
	DeletePosRewardRedemption(redemptionID string) error
	ReadAllPosRewardRedemptions() ([]*pb.PosRewardRedemption, error)
}

type posRewardRedemptionRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosRewardRedemptionRepository(db *gorm.DB, redis *redis.Client) PosRewardRedemptionRepository {
	return &posRewardRedemptionRepository{
		db:    db,
		redis: redis,
	}
}
