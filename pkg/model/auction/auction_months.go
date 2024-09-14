package auction

import (
	"time"

	"gorm.io/gorm"
)

type AuctionMonths struct {
	ID        uint `gorm:"primarykey"`
	Month     uint
	Year      uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
