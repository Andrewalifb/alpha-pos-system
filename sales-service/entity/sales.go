package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type PosSale struct {
	SaleID          uuid.UUID `gorm:"type:uuid;primary_key" json:"sale_id"`
	ProductID       uuid.UUID `gorm:"type:uuid;not null" json:"product_id"`
	CustomerID      uuid.UUID `gorm:"type:uuid;not null" json:"customer_id"`
	Quantity        int       `gorm:"type:int;not null" json:"quantity"`
	SaleDate        time.Time `gorm:"type:timestamp;not null" json:"sale_date"`
	TotalPrice      float64   `gorm:"type:decimal(10,2);not null" json:"total_price"`
	StoreID         uuid.UUID `gorm:"type:uuid" json:"store_id"`
	CashierID       uuid.UUID `gorm:"type:uuid" json:"cashier_id"`
	PaymentMethodID uuid.UUID `gorm:"type:uuid;not null" json:"payment_method_id"`
	CompanyID       uuid.UUID `gorm:"type:uuid;not null" json:"company_id"`
	CreatedAt       time.Time `gorm:"type:timestamp" json:"created_at"`
	CreatedBy       uuid.UUID `gorm:"type:uuid" json:"created_by"`
	UpdatedAt       time.Time `gorm:"type:timestamp" json:"updated_at"`
	UpdatedBy       uuid.UUID `gorm:"type:uuid" json:"updated_by"`
}
