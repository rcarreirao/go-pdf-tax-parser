package parser

import (
	"fmt"
	"pdf_balance_parser/pkg/model/auction"
	"pdf_balance_parser/pkg/model/document"
	auction_repository "pdf_balance_parser/pkg/repository/auction"
	"regexp"
	"sort"
	"strings"

	"github.com/djimenez/iconv-go"
)

type auctionDayPFunc func(*string)

func LoadRegexMapAuction() {
	regexes = make(map[int]map[int]string)
	regexes[0] = map[int]string{0: `(?i)Data pregão`, 1: `[0-9\/]*`, 2: `(?i)[0-9\/]*`, 3: `\d[0-9\/]*`}
	regexes[1] = map[int]string{0: `(?i)Codigo do Cliente`, 1: `[0-9]*`, 2: ``, 3: `\d[0-9]*`}
	regexes[2] = map[int]string{0: `(?i)Nr. nota`, 1: `[0-9.]*`, 2: ``, 3: `\d[0-9.]*`}
	for key := range regexes {
		regexesKeys = append(regexesKeys, key)
	}
	sort.Ints(regexesKeys)
}

func ParseDocumentAuction(d document.Document) auction.AuctionDays {
	LoadRegexMapAuction()

	output, _ := iconv.ConvertString(d.Content, "iso-8859-1", "iso-8859-1")
	fmt.Println(output)

	var headerRegex *regexp.Regexp
	var regexError error
	var bodyRegex *regexp.Regexp
	var auctionDay = new(auction.AuctionDays)
	var auctionDaySummary = new(auction.AuctionDaySummary)
	var auctionDayRowsFunc = []auctionDayPFunc{auctionDaySummary.AuctionDayLine1.ParseLines, auctionDaySummary.AuctionDayLine2.ParseLines, auctionDaySummary.AuctionDayLine3.ParseLines}

	for _, key := range regexesKeys {
		headerRegex, regexError = regexp.Compile(regexes[key][0] + regexes[key][1])
		if headerRegex != nil {
			var match string
			var index int
			for index, match = range headerRegex.FindStringSubmatch(d.Content) {
				fmt.Printf("[%d] %s\n", index, match)
			}
			bodyRegex, _ = regexp.Compile(regexes[key][3])
			for index, matches := range bodyRegex.FindStringSubmatch(match) {
				replaceBoundariesRowValues(regexes[key][2], &matches)
				fmt.Printf("[%d] %s\n", index, strings.Split(matches, " "))
				auctionDayRowsFunc[key](&matches)
			}
		} else {
			fmt.Printf(fmt.Sprintf("\033[0;31m%s\033[0m", regexError))
		}
	}
	composeAuctionDay(auctionDaySummary, auctionDay)
	storeOrUpdateAuctionDay(auctionDay)
	return *auctionDay
}

func composeAuctionDay(auctionDaySummary *auction.AuctionDaySummary, auctionDay *auction.AuctionDays) {
	auctionDay.AuctionDay = auctionDaySummary.AuctionDayLine1.AuctionDay
	auctionDay.CustomerCode = auctionDaySummary.AuctionDayLine2.Value
	auctionDay.InvoiceId = auctionDaySummary.AuctionDayLine3.Value
}

func storeAuctionDay(auctionDay *auction.AuctionDays) {
	tradingRepository := new(auction_repository.AuctionDayRepository)
	tradingRepository.New().Store(auctionDay)
}
func storeOrUpdateAuctionDay(auctionDay *auction.AuctionDays) {
	tradingRepository := new(auction_repository.AuctionDayRepository)
	tradingRepository.New().StoreOrUpdate(&auction.AuctionDays{AuctionDay: auctionDay.AuctionDay}, auctionDay)
}

func listAuctionDays() {
	tradingRepository := new(auction_repository.AuctionDayRepository)
	tradingRepository.List()

}
