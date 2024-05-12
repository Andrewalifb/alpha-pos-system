package repository

import (
	pb "github.com/Andrewalifb/alpha-pos-system/company-service/api/proto"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type PosRoleRepository interface {
	CreatePosRole(posRole *pb.PosRole) error
	ReadPosRole(roleID string) (*pb.PosRole, error)
	UpdatePosRole(posRole *pb.PosRole) error
	DeletePosRole(roleID string) error
	ReadAllPosRoles() ([]*pb.PosRole, error)
}

type posRoleRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosRoleRepository(db *gorm.DB, redis *redis.Client) PosRoleRepository {
	return &posRoleRepository{
		db:    db,
		redis: redis,
	}
}
