package repository

import (
	pb "github.com/Andrewalifb/alpha-pos-system/sales-service/api/proto"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type PosOnlinePaymentRepository interface {
	CreatePosOnlinePayment(posOnlinePayment *pb.PosOnlinePayment) error
	ReadPosOnlinePayment(paymentID string) (*pb.PosOnlinePayment, error)
	UpdatePosOnlinePayment(posOnlinePayment *pb.PosOnlinePayment) error
	DeletePosOnlinePayment(paymentID string) error
	ReadAllPosOnlinePayments() ([]*pb.PosOnlinePayment, error)
}

type posOnlinePaymentRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosOnlinePaymentRepository(db *gorm.DB, redis *redis.Client) PosOnlinePaymentRepository {
	return &posOnlinePaymentRepository{
		db:    db,
		redis: redis,
	}
}
