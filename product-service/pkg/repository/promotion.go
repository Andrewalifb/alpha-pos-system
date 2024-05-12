package repository

import (
	pb "github.com/Andrewalifb/alpha-pos-system/product-service/api/proto"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type PosPromotionRepository interface {
	CreatePosPromotion(posPromotion *pb.PosPromotion) error
	ReadPosPromotion(promotionID string) (*pb.PosPromotion, error)
	UpdatePosPromotion(posPromotion *pb.PosPromotion) error
	DeletePosPromotion(promotionID string) error
	ReadAllPosPromotions() ([]*pb.PosPromotion, error)
}

type posPromotionRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosPromotionRepository(db *gorm.DB, redis *redis.Client) PosPromotionRepository {
	return &posPromotionRepository{
		db:    db,
		redis: redis,
	}
}
