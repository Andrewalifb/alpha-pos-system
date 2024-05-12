package repository

import (
	pb "github.com/Andrewalifb/alpha-pos-system/sales-service/api/proto"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type PosInvoiceRepository interface {
	CreatePosInvoice(posInvoice *pb.PosInvoice) error
	ReadPosInvoice(invoiceID string) (*pb.PosInvoice, error)
	UpdatePosInvoice(posInvoice *pb.PosInvoice) error
	DeletePosInvoice(invoiceID string) error
	ReadAllPosInvoices() ([]*pb.PosInvoice, error)
}

type posInvoiceRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosInvoiceRepository(db *gorm.DB, redis *redis.Client) PosInvoiceRepository {
	return &posInvoiceRepository{
		db:    db,
		redis: redis,
	}
}
