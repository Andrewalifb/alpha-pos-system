package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type PosCashDrawer struct {
	DrawerID        uuid.UUID `gorm:"type:uuid;primary_key" json:"drawer_id"`
	StoreID         uuid.UUID `gorm:"type:uuid" json:"store_id"`
	EmployeeID      uuid.UUID `gorm:"type:uuid" json:"employee_id"`
	CashIn          float64   `gorm:"type:decimal(10,2);not null" json:"cash_in"`
	CashOut         float64   `gorm:"type:decimal(10,2);not null" json:"cash_out"`
	TransactionTime time.Time `gorm:"type:timestamp;not null" json:"transaction_time"`
	RoleID          uuid.UUID `gorm:"type:uuid;not null" json:"role_id"`
	CompanyID       uuid.UUID `gorm:"type:uuid;not null" json:"company_id"`
	CreatedAt       time.Time `gorm:"type:timestamp" json:"created_at"`
	CreatedBy       uuid.UUID `gorm:"type:uuid" json:"created_by"`
	UpdatedAt       time.Time `gorm:"type:timestamp" json:"updated_at"`
	UpdatedBy       uuid.UUID `gorm:"type:uuid" json:"updated_by"`
}
