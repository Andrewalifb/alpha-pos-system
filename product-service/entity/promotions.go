package entity

type PosPromotion struct {
	PromotionID  uuid.UUID `gorm:"type:uuid;primary_key" json:"promotion_id"`
	ProductID    uuid.UUID `gorm:"type:uuid;not null" json:"product_id"`
	StartDate    time.Time `gorm:"type:date;not null" json:"start_date"`
	EndDate      time.Time `gorm:"type:date;not null" json:"end_date"`
	DiscountRate float64   `gorm:"type:decimal(3,2);not null" json:"discount_rate"`
	CompanyID    uuid.UUID `gorm:"type:uuid;not null" json:"company_id"`
	CreatedAt    time.Time `gorm:"type:timestamp" json:"created_at"`
	CreatedBy    uuid.UUID `gorm:"type:uuid" json:"created_by"`
	UpdatedAt    time.Time `gorm:"type:timestamp" json:"updated_at"`
	UpdatedBy    uuid.UUID `gorm:"type:uuid" json:"updated_by"`
}
