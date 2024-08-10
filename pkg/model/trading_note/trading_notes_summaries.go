package trading_note

import "gorm.io/gorm"

type TradingNoteSummaries struct {
	gorm.Model
	SellAvailable           float64
	BuyAvailable            float64
	SellOptions             float64
	BuyOptions              float64
	BusinessValue           float64
	BusinessValueOp         string
	BusinessType            float64
	OperationPrice          float64
	OperationType           float64
	OperationalTax          float64
	Irrf                    float64
	IrrfDayTrade            float64
	OperationTax            float64
	TaxRegisterBmf          float64
	TaxBmf                  float64
	TaxBmfOp                string
	OtherCosts              float64
	Taxes                   float64
	PositionAdjustment      float64
	DayTradeAdjustment      float64
	DayTradeAdjustmentOp    string
	TotalOperationalCosts   float64
	TotalOperationalCostsOp string
}
