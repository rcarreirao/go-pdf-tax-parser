package auction

import (
	"time"
)

type AuctionDayLine1 struct {
	Date       string
	AuctionDay time.Time
}

func (al *AuctionDayLine1) ParseLines(line *string) {
	layout := "02/01/2006"
	date, _ := time.Parse(layout, *line)
	al.Date = *line
	al.AuctionDay = date

}
