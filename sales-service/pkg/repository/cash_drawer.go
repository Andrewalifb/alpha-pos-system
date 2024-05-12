package repository

import (
	pb "github.com/Andrewalifb/alpha-pos-system/sales-service/api/proto"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type PosCashDrawerRepository interface {
	CreatePosCashDrawer(posCashDrawer *pb.PosCashDrawer) error
	ReadPosCashDrawer(drawerID string) (*pb.PosCashDrawer, error)
	UpdatePosCashDrawer(posCashDrawer *pb.PosCashDrawer) error
	DeletePosCashDrawer(drawerID string) error
	ReadAllPosCashDrawers() ([]*pb.PosCashDrawer, error)
}

type posCashDrawerRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosCashDrawerRepository(db *gorm.DB, redis *redis.Client) PosCashDrawerRepository {
	return &posCashDrawerRepository{
		db:    db,
		redis: redis,
	}
}
