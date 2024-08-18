package parser

import (
	"pdf_balance_parser/pkg/helper"
	"strconv"
)

type Line1 struct {
	SellAvailable   float64
	BuyAvailable    float64
	SellOptions     float64
	BuyOptions      float64
	BusinessValue   float64
	BusinessValueOp string
}

func (l *Line1) ParseLines(line *string) {
	var columns []string
	columns = helper.Explode(" ", *line)
	l.SellAvailable, _ = strconv.ParseFloat(columns[0], 64)
	l.BuyAvailable, _ = strconv.ParseFloat(columns[1], 64)
	l.SellOptions, _ = strconv.ParseFloat(columns[2], 64)
	l.BusinessValue, _ = strconv.ParseFloat(columns[3], 64)
	l.BusinessValueOp = columns[4]

}
