package auction

type AuctionDayLine2 struct {
	Value string
}

func (al *AuctionDayLine2) ParseLines(line *string) {
	al.Value = *line
}
