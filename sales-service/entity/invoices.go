package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type PosInvoice struct {
	InvoiceID uuid.UUID `gorm:"type:uuid;primary_key" json:"invoice_id"`
	SaleID    uuid.UUID `gorm:"type:uuid;not null" json:"sale_id"`
	Date      time.Time `gorm:"type:timestamp;not null" json:"date"`
	Total     float64   `gorm:"type:decimal(10,2);not null" json:"total"`
	Discounts float64   `gorm:"type:decimal(10,2)" json:"discounts"`
	Taxes     float64   `gorm:"type:decimal(10,2)" json:"taxes"`
	CompanyID uuid.UUID `gorm:"type:uuid;not null" json:"company_id"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	CreatedBy uuid.UUID `gorm:"type:uuid" json:"created_by"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
	UpdatedBy uuid.UUID `gorm:"type:uuid" json:"updated_by"`
}
