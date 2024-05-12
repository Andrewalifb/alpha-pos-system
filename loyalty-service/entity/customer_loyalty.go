package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type PosCustomerLoyalty struct {
	CustomerID   uuid.UUID `gorm:"type:uuid;not null" json:"customer_id"`
	LevelID      uuid.UUID `gorm:"type:uuid;not null" json:"level_id"`
	RewardPoints int       `gorm:"type:int;default:0" json:"reward_points"`
	CompanyID    uuid.UUID `gorm:"type:uuid;not null" json:"company_id"`
	CreatedAt    time.Time `gorm:"type:timestamp" json:"created_at"`
	CreatedBy    uuid.UUID `gorm:"type:uuid" json:"created_by"`
	UpdatedAt    time.Time `gorm:"type:timestamp" json:"updated_at"`
	UpdatedBy    uuid.UUID `gorm:"type:uuid" json:"updated_by"`
}
