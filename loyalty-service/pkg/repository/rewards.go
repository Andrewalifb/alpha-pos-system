package repository

import (
	pb "github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/proto"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type PosRewardRepository interface {
	CreatePosReward(posReward *pb.PosReward) error
	ReadPosReward(rewardID string) (*pb.PosReward, error)
	UpdatePosReward(posReward *pb.PosReward) error
	DeletePosReward(rewardID string) error
	ReadAllPosRewards() ([]*pb.PosReward, error)
}

type posRewardRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosRewardRepository(db *gorm.DB, redis *redis.Client) PosRewardRepository {
	return &posRewardRepository{
		db:    db,
		redis: redis,
	}
}
