package repository

import (
	pb "github.com/Andrewalifb/alpha-pos-system/sales-service/api/proto"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type PosReturnRepository interface {
	CreatePosReturn(posReturn *pb.PosReturn) error
	ReadPosReturn(returnID string) (*pb.PosReturn, error)
	UpdatePosReturn(posReturn *pb.PosReturn) error
	DeletePosReturn(returnID string) error
	ReadAllPosReturns() ([]*pb.PosReturn, error)
}

type posReturnRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosReturnRepository(db *gorm.DB, redis *redis.Client) PosReturnRepository {
	return &posReturnRepository{
		db:    db,
		redis: redis,
	}
}
