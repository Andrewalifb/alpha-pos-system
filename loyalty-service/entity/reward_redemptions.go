package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type PosRewardRedemption struct {
	RedemptionID   uuid.UUID `gorm:"type:uuid;primary_key" json:"redemption_id"`
	RewardID       uuid.UUID `gorm:"type:uuid;not null" json:"reward_id"`
	CustomerID     uuid.UUID `gorm:"type:uuid;not null" json:"customer_id"`
	RedemptionDate time.Time `gorm:"type:timestamp;not null" json:"redemption_date"`
	CompanyID      uuid.UUID `gorm:"type:uuid;not null" json:"company_id"`
	CreatedAt      time.Time `gorm:"type:timestamp" json:"created_at"`
	CreatedBy      uuid.UUID `gorm:"type:uuid" json:"created_by"`
	UpdatedAt      time.Time `gorm:"type:timestamp" json:"updated_at"`
	UpdatedBy      uuid.UUID `gorm:"type:uuid" json:"updated_by"`
}
