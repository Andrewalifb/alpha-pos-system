package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type PosReturn struct {
	ReturnID   uuid.UUID `gorm:"type:uuid;primary_key" json:"return_id"`
	SaleID     uuid.UUID `gorm:"type:uuid;not null" json:"sale_id"`
	ProductID  uuid.UUID `gorm:"type:uuid;not null" json:"product_id"`
	Quantity   int       `gorm:"type:int;not null" json:"quantity"`
	ReturnDate time.Time `gorm:"type:timestamp;not null" json:"return_date"`
	Reason     string    `gorm:"type:text" json:"reason"`
	CompanyID  uuid.UUID `gorm:"type:uuid;not null" json:"company_id"`
	CreatedAt  time.Time `gorm:"type:timestamp" json:"created_at"`
	CreatedBy  uuid.UUID `gorm:"type:uuid" json:"created_by"`
	UpdatedAt  time.Time `gorm:"type:timestamp" json:"updated_at"`
	UpdatedBy  uuid.UUID `gorm:"type:uuid" json:"updated_by"`
}
