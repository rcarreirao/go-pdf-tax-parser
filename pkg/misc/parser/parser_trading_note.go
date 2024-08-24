package parser

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/rcarreirao/pdf_balance_parser/pkg/model/auction"
	"github.com/rcarreirao/pdf_balance_parser/pkg/model/document"

	"github.com/rcarreirao/pdf_balance_parser/pkg/model/parser"
	"github.com/rcarreirao/pdf_balance_parser/pkg/model/trading_note"
	trading_note_repository "github.com/rcarreirao/pdf_balance_parser/pkg/repository/trading_note"

	"github.com/djimenez/iconv-go"
)

type pFunc func(*string)

func LoadRegexMapTradingNote() {
	regexes = make(map[int]map[int]string)
	regexes[0] = map[int]string{0: `(?i)Venda dis`, 1: `[0-9a-zA-ZÀ-ÿ ]*`, 2: `ócios`, 3: `(?i)ócios[0-9, |a-zA-Z]* [C|D]+ `}
	regexes[1] = map[int]string{0: `(?i)IRRF`, 1: `[0-9, |a-z A-ZÀ-ÿ&+().]*`, 2: `gar\)`, 3: `(?i)gar\)[0-9, |a-zA-Z]* [C|D]+ `}
	regexes[2] = map[int]string{0: `(?i)\+Ou`, 1: `[0-9, |a-z A-ZÀ-ÿ&+().]*`, 2: `ais`, 3: `(?i)ais[0-9, |a-zA-Z]* [C|D]+ `}
	regexes[3] = map[int]string{0: `(?i) Outros`, 1: `[0-9, |a-z A-ZÀ-ÿ&+().#]*`, 2: `nota`, 3: `(?i)nota[0-9, |a-zA-Z]* [C|D]+ `}
	for key := range regexes {
		regexesKeys = append(regexesKeys, key)
	}
	sort.Ints(regexesKeys)
}

func adjustCurrencyFormat(line *string) {
	*line = strings.Replace(*line, ".", "", 10)
	*line = strings.Replace(*line, ",", ".", 10)
}

func ParseDocumentTradingNote(auctionDay auction.AuctionDays, d document.Document) {
	output, _ := iconv.ConvertString(d.Content, "iso-8859-1", "iso-8859-1")
	fmt.Println(output)
	LoadRegexMapTradingNote()

	var headerRegex *regexp.Regexp
	var regexError error
	var bodyRegex *regexp.Regexp
	summary := new(parser.Summary)
	var rowsFunc = []pFunc{summary.Line1.ParseLines, summary.Line2.ParseLines, summary.Line3.ParseLines, summary.Line4.ParseLines}

	for _, key := range regexesKeys {
		headerRegex, regexError = regexp.Compile(regexes[key][0] + regexes[key][1] + regexes[key][2])
		if headerRegex != nil {
			for index, match := range headerRegex.FindStringSubmatch(d.Content) {
				fmt.Printf("[%d] %s\n", index, match)
			}
			bodyRegex, _ = regexp.Compile(regexes[key][3])
			for index, match := range bodyRegex.FindStringSubmatch(d.Content) {
				replaceBoundariesRowValues(regexes[key][2], &match)
				adjustCurrencyFormat(&match)
				fmt.Printf("[%d] %s\n", index, strings.Split(match, " "))
				rowsFunc[key](&match)
			}
		} else {
			fmt.Printf(fmt.Sprintf("\033[0;31m%s\033[0m", regexError))
		}
	}
	var tradingNote = new(trading_note.TradingNotes)
	var tradingNoteSummary = new(trading_note.TradingNoteSummaries)
	tradingNoteSummary.AuctionDay = auctionDay
	composeTradingNote(summary, tradingNote)
	composeTradingNoteSummary(summary, tradingNoteSummary)
	store(tradingNoteSummary)
}

func composeTradingNote(summary *parser.Summary, trading_note *trading_note.TradingNotes) {
}
func composeTradingNoteSummary(summary *parser.Summary, trading_note_summary *trading_note.TradingNoteSummaries) {
	// Line 1
	trading_note_summary.SellAvailable = summary.Line1.SellAvailable
	trading_note_summary.BuyAvailable = summary.Line1.BuyAvailable
	trading_note_summary.SellOptions = summary.Line1.SellOptions
	trading_note_summary.BuyOptions = summary.Line1.BuyOptions
	trading_note_summary.BusinessValue = summary.Line1.BusinessValue
	trading_note_summary.BusinessValueOp = summary.Line1.BusinessValueOp
	// Line 2
	trading_note_summary.Irrf = summary.Line2.Irrf
	trading_note_summary.IrrfDayTrade = summary.Line2.IrrfDayTrade
	trading_note_summary.OperationalTax = summary.Line2.OperationalTax
	trading_note_summary.TaxRegisterBmf = summary.Line2.TaxRegisterBmf
	trading_note_summary.TaxBmf = summary.Line2.TaxBmf
	trading_note_summary.TaxBmfOp = summary.Line2.TaxBmfOp

	// Line 3
	trading_note_summary.OtherCosts = summary.Line3.OtherCosts
	trading_note_summary.Taxes = summary.Line3.Taxes
	trading_note_summary.PositionAdjustment = summary.Line3.PositionAdjustment
	trading_note_summary.DayTradeAdjustment = summary.Line3.DayTradeAdjustment
	trading_note_summary.DayTradeAdjustmentOp = summary.Line3.DayTradeAdjustmentOp
	trading_note_summary.TotalOperationalCosts = summary.Line3.TotalOperationalCosts
	trading_note_summary.TotalOperationalCostsOp = summary.Line3.TotalOperationalCostsOp

	// Line 4
	trading_note_summary.Others = summary.Line4.Others
	trading_note_summary.IrrfOperational = summary.Line4.IrrfOperational
	trading_note_summary.TotalInvestmentAccount = summary.Line4.TotalInvestmentAccount
	trading_note_summary.TotalNormalAccount = summary.Line4.TotalNormalAccount
	trading_note_summary.TotalNet = summary.Line4.TotalNet
	trading_note_summary.TotalNetOp = summary.Line4.TotalNetOp
	trading_note_summary.TotalNetInvoice = summary.Line4.TotalNetInvoice
	trading_note_summary.TotalNetInvoiceOp = summary.Line4.TotalNetInvoiceOp
}

func store(trading_note_summary *trading_note.TradingNoteSummaries) {
	tradingRepository := new(trading_note_repository.TradingNoteRepository)
	tradingRepository.New()
	tradingRepository.StoreOrUpdate(map[string]interface{}{"auction_day_id": trading_note_summary.AuctionDay.ID}, trading_note_summary)
}
