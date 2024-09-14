package auction_repository

import (
	"os"

	"github.com/rcarreirao/pdf_balance_parser/pkg/model/auction"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type AuctionMonthRepository struct {
	connection gorm.DB
}

func (r *AuctionMonthRepository) New() AuctionMonthRepository {
	var err error
	conn, err := gorm.Open(sqlite.Open(os.Getenv("DB_DATABASE")), &gorm.Config{})
	r.connection = *conn
	if err != nil {
		panic("failed to connect on database")
	}
	return *r
}

func (repository AuctionMonthRepository) Store(auctionMonth *auction.AuctionMonths) {
	repository.connection.Create(auctionMonth)
}

func (repository AuctionMonthRepository) StoreOrUpdate(conditional *auction.AuctionMonths, AuctionMonth *auction.AuctionMonths) *auction.AuctionMonths {
	repository.connection.FirstOrCreate(AuctionMonth, conditional)
	repository.connection.Model(AuctionMonth).Updates(AuctionMonth)
	return AuctionMonth
}

func (repository AuctionMonthRepository) Find(id uint) auction.AuctionMonths {
	var AuctionMonth = auction.AuctionMonths{ID: id}
	repository.connection.First(&AuctionMonth)
	return AuctionMonth
}

func (repository AuctionMonthRepository) List() []auction.AuctionMonths {
	var auctionMonths []auction.AuctionMonths
	repository.connection.Find(&auctionMonths)
	return auctionMonths
}
