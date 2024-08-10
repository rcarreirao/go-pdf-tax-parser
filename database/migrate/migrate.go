package migrate

import (
	"fmt"
	"pdf_balance_parser/pkg/model/trading_note"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func MigrateExec() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	fmt.Println("Migrating Trading Notes")
	db.AutoMigrate(&trading_note.TradingNotes{})
	fmt.Println("Migrating Trading Note Summary")
	db.AutoMigrate(&trading_note.TradingNoteSummaries{})
}
