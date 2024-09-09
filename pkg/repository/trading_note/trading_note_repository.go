package trading_note_repository

import (
	"os"

	"github.com/rcarreirao/pdf_balance_parser/pkg/model/trading_note"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TradingNoteRepository struct {
	connection gorm.DB
}

func (r *TradingNoteRepository) New() TradingNoteRepository {
	var err error
	conn, err := gorm.Open(sqlite.Open(os.Getenv("DB_DATABASE")), &gorm.Config{})
	r.connection = *conn
	if err != nil {
		panic("failed to connect on database")
	}
	return *r
}

func (repository TradingNoteRepository) Store(trading_note_summary *trading_note.TradingNoteSummaries) {
	repository.connection.Create(trading_note_summary)
}

func (repository TradingNoteRepository) StoreOrUpdate(conditional map[string]interface{}, trading_note_summary *trading_note.TradingNoteSummaries) {
	repository.connection.FirstOrCreate(trading_note_summary, conditional)
	repository.connection.Model(trading_note_summary).Updates(trading_note_summary)
}

func (repository TradingNoteRepository) List() []trading_note.TradingNoteSummaries {
	var tradingNoteSummaries []trading_note.TradingNoteSummaries
	repository.connection.Find(&tradingNoteSummaries)
	return tradingNoteSummaries
}

func (repository TradingNoteRepository) Search(tradingNoteSummary trading_note.TradingNoteSummaries) []trading_note.TradingNoteSummaries {
	var tradingNoteSummaries []trading_note.TradingNoteSummaries
	repository.connection.Find(&tradingNoteSummaries, tradingNoteSummary)
	return tradingNoteSummaries
}
