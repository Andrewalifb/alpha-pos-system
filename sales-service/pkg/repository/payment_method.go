package repository

import (
	pb "github.com/Andrewalifb/alpha-pos-system/sales-service/api/proto"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type PosPaymentMethodRepository interface {
	CreatePosPaymentMethod(posPaymentMethod *pb.PosPaymentMethod) error
	ReadPosPaymentMethod(paymentMethodID string) (*pb.PosPaymentMethod, error)
	UpdatePosPaymentMethod(posPaymentMethod *pb.PosPaymentMethod) error
	DeletePosPaymentMethod(paymentMethodID string) error
	ReadAllPosPaymentMethods() ([]*pb.PosPaymentMethod, error)
}

type posPaymentMethodRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosPaymentMethodRepository(db *gorm.DB, redis *redis.Client) PosPaymentMethodRepository {
	return &posPaymentMethodRepository{
		db:    db,
		redis: redis,
	}
}
