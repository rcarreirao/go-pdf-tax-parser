package auction

import (
	"strconv"
)

type AuctionDayLine3 struct {
	Value float64
}

func (al *AuctionDayLine3) ParseLines(line *string) {
	al.Value, _ = strconv.ParseFloat(*line, 64)
}
