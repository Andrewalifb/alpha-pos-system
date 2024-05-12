package repository

import (
	pb "github.com/Andrewalifb/alpha-pos-system/company-service/api/proto"
	"github.com/Andrewalifb/alpha-pos-system/company-service/entity"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PosUserRepository interface {
	CreatePosUser(posUser *entity.PosUser) error
	ReadPosUser(userID string) (*pb.PosUser, error)
	ReadPosUserByUsername(username string) (*pb.PosUser, error) // Add this line
	UpdatePosUser(posUser *pb.PosUser) error
	DeletePosUser(userID string) error
	ReadAllPosUsers() ([]*pb.PosUser, error)
}

type posUserRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewPosUserRepository(db *gorm.DB, redis *redis.Client) PosUserRepository {
	return &posUserRepository{
		db:    db,
		redis: redis,
	}
}

func (r *posUserRepository) CreatePosUser(posUser *entity.PosUser) error {
	result := r.db.Create(posUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *posUserRepository) ReadPosUserByUsername(username string) (*pb.PosUser, error) {
	var user pb.PosUser
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	t := user.CreatedAt.AsTime()
	user.CreatedAt = timestamppb.New(t)
	return &user, nil
}
