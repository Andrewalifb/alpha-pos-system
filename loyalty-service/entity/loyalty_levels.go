package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type PosLoyaltyLevel struct {
	LevelID        uuid.UUID `gorm:"type:uuid;primary_key" json:"level_id"`
	LevelName      string    `gorm:"type:varchar(50);not null" json:"level_name"`
	PointsRequired int       `gorm:"type:int;not null" json:"points_required"`
	CompanyID      uuid.UUID `gorm:"type:uuid;not null" json:"company_id"`
	CreatedAt      time.Time `gorm:"type:timestamp" json:"created_at"`
	CreatedBy      uuid.UUID `gorm:"type:uuid" json:"created_by"`
	UpdatedAt      time.Time `gorm:"type:timestamp" json:"updated_at"`
	UpdatedBy      uuid.UUID `gorm:"type:uuid" json:"updated_by"`
}
