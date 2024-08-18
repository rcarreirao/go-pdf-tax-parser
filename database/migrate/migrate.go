package migrate

import (
	"fmt"
	"os"
	"pdf_balance_parser/pkg/model/auction"
	"pdf_balance_parser/pkg/model/trading_note"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func MigrateExec() {
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_DATABASE")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	fmt.Println("Migrating Auction Days")
	db.AutoMigrate(&auction.AuctionDays{})
	fmt.Println("Migrating Trading Notes")
	db.AutoMigrate(&trading_note.TradingNotes{})
	fmt.Println("Migrating Trading Note Summary")
	db.AutoMigrate(&trading_note.TradingNoteSummaries{})
}
