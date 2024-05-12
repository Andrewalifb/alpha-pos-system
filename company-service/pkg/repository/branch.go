package repository

import (
	pb "github.com/Andrewalifb/alpha-pos-system/company-service/api/proto"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

type PosStoreBranchRepository interface {
	CreatePosStoreBranch(posStoreBranch *pb.PosStoreBranch) error
	ReadPosStoreBranch(branchID string) (*pb.PosStoreBranch, error)
	UpdatePosStoreBranch(posStoreBranch *pb.PosStoreBranch) error
	DeletePosStoreBranch(branchID string) error
	ReadAllPosStoreBranches() ([]*pb.PosStoreBranch, error)
}

type posStoreBranchRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosStoreBranchRepository(db *gorm.DB, redis *redis.Client) PosStoreBranchRepository {
	return &posStoreBranchRepository{
		db:    db,
		redis: redis,
	}
}
