package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type PosPointsTransaction struct {
	TransactionID   uuid.UUID `gorm:"type:uuid;primary_key" json:"transaction_id"`
	CustomerID      uuid.UUID `gorm:"type:uuid;not null" json:"customer_id"`
	Points          int       `gorm:"type:int;not null" json:"points"`
	TransactionDate time.Time `gorm:"type:timestamp;not null" json:"transaction_date"`
	TransactionType string    `gorm:"type:varchar(50);not null" json:"transaction_type"` // e.g., 'Earned', 'Redeemed'
	CompanyID       uuid.UUID `gorm:"type:uuid;not null" json:"company_id"`
	CreatedAt       time.Time `gorm:"type:timestamp" json:"created_at"`
	CreatedBy       uuid.UUID `gorm:"type:uuid" json:"created_by"`
	UpdatedAt       time.Time `gorm:"type:timestamp" json:"updated_at"`
	UpdatedBy       uuid.UUID `gorm:"type:uuid" json:"updated_by"`
}
