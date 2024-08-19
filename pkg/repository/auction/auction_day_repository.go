package auction_repository

import (
	"os"
	"pdf_balance_parser/pkg/model/auction"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type AuctionDayRepository struct {
	connection gorm.DB
}

func (r *AuctionDayRepository) New() AuctionDayRepository {
	var err error
	conn, err := gorm.Open(sqlite.Open(os.Getenv("DB_DATABASE")), &gorm.Config{})
	r.connection = *conn
	if err != nil {
		panic("failed to connect on database")
	}
	return *r
}

func (repository AuctionDayRepository) Store(auctionDay *auction.AuctionDays) {
	repository.connection.Create(auctionDay)
}

func (repository AuctionDayRepository) StoreOrUpdate(auctionDay *auction.AuctionDays) auction.AuctionDays {
	var tempAuctionDay = repository.Find(auctionDay.ID)
	return tempAuctionDay
}

func (repository AuctionDayRepository) Find(id uint) auction.AuctionDays {
	var AuctionDay = auction.AuctionDays{ID: id}
	repository.connection.First(&AuctionDay)
	return AuctionDay
}
