package auction

import (
	"time"

	"gorm.io/gorm"
)

type AuctionDays struct {
	ID           uint `gorm:"primarykey"`
	AuctionDay   time.Time
	CustomerCode string
	InvoiceId    float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
