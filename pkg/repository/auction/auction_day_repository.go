package auction_repository

import (
	"os"

	"github.com/rcarreirao/pdf_balance_parser/pkg/model/auction"

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

func (repository AuctionDayRepository) StoreOrUpdate(conditional *auction.AuctionDays, AuctionDay *auction.AuctionDays) *auction.AuctionDays {
	repository.connection.FirstOrCreate(AuctionDay, conditional)
	repository.connection.Model(AuctionDay).Updates(AuctionDay)
	return AuctionDay
}

func (repository AuctionDayRepository) Find(id uint) auction.AuctionDays {
	var AuctionDay = auction.AuctionDays{ID: id}
	repository.connection.First(&AuctionDay)
	return AuctionDay
}

func (repository AuctionDayRepository) List() []auction.AuctionDays {
	var auctionDays []auction.AuctionDays
	repository.connection.Find(&auctionDays)
	return auctionDays
}

func (repository AuctionDayRepository) ListMonthlyAuctions() []auction.MonthlyAuction {
	var monthlyAuction []auction.MonthlyAuction
	repository.connection.Table("auction_days").Distinct("strftime('%m-%Y', auction_day) as AuctionMonth").Find(&monthlyAuction)
	return monthlyAuction
}
