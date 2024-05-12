package repository

import (
	pb "github.com/Andrewalifb/alpha-pos-system/company-service/api/proto"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type PosCompanyRepository interface {
	CreatePosCompany(posCompany *pb.PosCompany) error
	ReadPosCompany(companyID string) (*pb.PosCompany, error)
	UpdatePosCompany(posCompany *pb.PosCompany) error
	DeletePosCompany(companyID string) error
	ReadAllPosCompanies() ([]*pb.PosCompany, error)
}

type posCompanyRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosCompanyRepository(db *gorm.DB, redis *redis.Client) PosCompanyRepository {
	return &posCompanyRepository{
		db:    db,
		redis: redis,
	}
}
