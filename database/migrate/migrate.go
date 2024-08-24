package migrate

import (
	"fmt"
	"os"

	"github.com/rcarreirao/pdf_balance_parser/pkg/helper"
	"github.com/rcarreirao/pdf_balance_parser/pkg/model/auction"
	"github.com/rcarreirao/pdf_balance_parser/pkg/model/trading_note"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func MigrateExec() {
	db := checkDatabaseFile()

	// Migrate the schema
	fmt.Println("Migrating Auction Days")
	db.AutoMigrate(&auction.AuctionDays{})
	fmt.Println("Migrating Trading Notes")
	db.AutoMigrate(&trading_note.TradingNotes{})
	fmt.Println("Migrating Trading Note Summary")
	db.AutoMigrate(&trading_note.TradingNoteSummaries{})
}

func checkDatabaseFile() *gorm.DB {
	pathInfo := helper.Pathinfo(os.Getenv("DB_DATABASE"))
	if !helper.Exists(pathInfo["dirname"]) {
		os.MkdirAll(pathInfo["dirname"], 0755)
	}
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_DATABASE")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
