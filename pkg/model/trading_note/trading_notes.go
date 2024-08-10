package trading_note

import "gorm.io/gorm"

type TradingNotes struct {
	gorm.Model
	Operation      string
	Commodity      string
	Due            string
	Quantity       string
	Price          uint
	BusinessType   string
	OperationPrice string
	OperationType  string
	OperationalTax float32
}
