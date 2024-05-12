package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type PosReward struct {
	RewardID       uuid.UUID `gorm:"type:uuid;primary_key" json:"reward_id"`
	RewardName     string    `gorm:"type:varchar(255);not null" json:"reward_name"`
	PointsRequired int       `gorm:"type:int;not null" json:"points_required"`
	CompanyID      uuid.UUID `gorm:"type:uuid;not null" json:"company_id"`
	CreatedAt      time.Time `gorm:"type:timestamp" json:"created_at"`
	CreatedBy      uuid.UUID `gorm:"type:uuid" json:"created_by"`
	UpdatedAt      time.Time `gorm:"type:timestamp" json:"updated_at"`
	UpdatedBy      uuid.UUID `gorm:"type:uuid" json:"updated_by"`
}
