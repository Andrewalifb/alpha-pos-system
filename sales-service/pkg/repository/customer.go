package repository

import (
	pb "github.com/Andrewalifb/alpha-pos-system/sales-service/api/proto"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type PosCustomerRepository interface {
	CreatePosCustomer(posCustomer *pb.PosCustomer) error
	ReadPosCustomer(customerID string) (*pb.PosCustomer, error)
	UpdatePosCustomer(posCustomer *pb.PosCustomer) error
	DeletePosCustomer(customerID string) error
	ReadAllPosCustomers() ([]*pb.PosCustomer, error)
}

type posCustomerRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosCustomerRepository(db *gorm.DB, redis *redis.Client) PosCustomerRepository {
	return &posCustomerRepository{
		db:    db,
		redis: redis,
	}
}
