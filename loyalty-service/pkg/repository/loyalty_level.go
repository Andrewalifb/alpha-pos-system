package repository

import (
	pb "github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/proto"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type PosLoyaltyLevelRepository interface {
	CreatePosLoyaltyLevel(posLoyaltyLevel *pb.PosLoyaltyLevel) error
	ReadPosLoyaltyLevel(levelID string) (*pb.PosLoyaltyLevel, error)
	UpdatePosLoyaltyLevel(posLoyaltyLevel *pb.PosLoyaltyLevel) error
	DeletePosLoyaltyLevel(levelID string) error
	ReadAllPosLoyaltyLevels() ([]*pb.PosLoyaltyLevel, error)
}

type posLoyaltyLevelRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosLoyaltyLevelRepository(db *gorm.DB, redis *redis.Client) PosLoyaltyLevelRepository {
	return &posLoyaltyLevelRepository{
		db:    db,
		redis: redis,
	}
}
