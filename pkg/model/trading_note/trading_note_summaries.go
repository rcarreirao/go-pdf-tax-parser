package trading_note

import (
	"github.com/rcarreirao/pdf_balance_parser/pkg/model/auction"

	"gorm.io/gorm"
)

type TradingNoteSummaries struct {
	gorm.Model
	AuctionDayID    int
	AuctionDay      auction.AuctionDays
	SellAvailable   float64
	BuyAvailable    float64
	SellOptions     float64
	BuyOptions      float64
	BusinessValue   float64
	BusinessValueOp string

	Irrf           float64
	IrrfDayTrade   float64
	OperationalTax float64
	TaxRegisterBmf float64
	TaxBmf         float64
	TaxBmfOp       string

	OtherCosts              float64
	Taxes                   float64
	PositionAdjustment      float64
	DayTradeAdjustment      float64
	DayTradeAdjustmentOp    string
	TotalOperationalCosts   float64
	TotalOperationalCostsOp string

	Others                 float64
	IrrfOperational        float64
	TotalInvestmentAccount float64
	TotalNormalAccount     float64
	TotalNet               float64
	TotalNetOp             string
	TotalNetInvoice        float64
	TotalNetInvoiceOp      string
}

func (tradingNoteSummary TradingNoteSummaries) GetBusinessValue() float64 {
	var factor float64 = 1.0
	if tradingNoteSummary.BusinessValueOp == "D" {
		factor = -1.0
	}
	return tradingNoteSummary.BusinessValue * factor
}

func (tradingNoteSummary TradingNoteSummaries) GetTaxBmfValue() float64 {
	var factor float64 = 1.0
	if tradingNoteSummary.TaxBmfOp == "D" {
		factor = -1.0
	}
	return tradingNoteSummary.TaxBmf * factor
}

func (tradingNoteSummary TradingNoteSummaries) GetDayTradeAdjustmentValue() float64 {
	var factor float64 = 1.0
	if tradingNoteSummary.DayTradeAdjustmentOp == "D" {
		factor = -1.0
	}
	return tradingNoteSummary.DayTradeAdjustment * factor
}

func (tradingNoteSummary TradingNoteSummaries) GetTotalOperationalCostsValue() float64 {
	var factor float64 = 1.0
	if tradingNoteSummary.TotalOperationalCostsOp == "D" {
		factor = -1.0
	}
	return tradingNoteSummary.TotalOperationalCosts * factor
}

func (tradingNoteSummary TradingNoteSummaries) GetTotalNetValue() float64 {
	var factor float64 = 1.0
	if tradingNoteSummary.TotalNetOp == "D" {
		factor = -1.0
	}
	return tradingNoteSummary.TotalNet * factor
}

func (tradingNoteSummary TradingNoteSummaries) GetTotalNetInvoiceValue() float64 {
	var factor float64 = 1.0
	if tradingNoteSummary.TotalNetInvoiceOp == "D" {
		factor = -1.0
	}
	return tradingNoteSummary.TotalNetInvoice * factor
}
