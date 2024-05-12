package repository

import (
	pb "github.com/Andrewalifb/alpha-pos-system/loyalty-service/api/proto"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type PosPointsTransactionRepository interface {
	CreatePosPointsTransaction(posPointsTransaction *pb.PosPointsTransaction) error
	ReadPosPointsTransaction(transactionID string) (*pb.PosPointsTransaction, error)
	UpdatePosPointsTransaction(posPointsTransaction *pb.PosPointsTransaction) error
	DeletePosPointsTransaction(transactionID string) error
	ReadAllPosPointsTransactions() ([]*pb.PosPointsTransaction, error)
}

type posPointsTransactionRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosPointsTransactionRepository(db *gorm.DB, redis *redis.Client) PosPointsTransactionRepository {
	return &posPointsTransactionRepository{
		db:    db,
		redis: redis,
	}
}
