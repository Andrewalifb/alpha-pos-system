package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type PosOnlinePayment struct {
	PaymentID     uuid.UUID `gorm:"type:uuid;primary_key" json:"payment_id"`
	SaleID        uuid.UUID `gorm:"type:uuid" json:"sale_id"`
	EmployeeID    uuid.UUID `gorm:"type:uuid" json:"employee_id"`
	PaymentDate   time.Time `gorm:"type:timestamp;not null" json:"payment_date"`
	Amount        float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
	PaymentMethod uuid.UUID `gorm:"type:uuid" json:"payment_method"`
	RoleID        uuid.UUID `gorm:"type:uuid;not null" json:"role_id"`
	CompanyID     uuid.UUID `gorm:"type:uuid;not null" json:"company_id"`
	CreatedAt     time.Time `gorm:"type:timestamp" json:"created_at"`
	CreatedBy     uuid.UUID `gorm:"type:uuid" json:"created_by"`
	UpdatedAt     time.Time `gorm:"type:timestamp" json:"updated_at"`
	UpdatedBy     uuid.UUID `gorm:"type:uuid" json:"updated_by"`
}
