package parser

import (
	"strconv"

	"github.com/rcarreirao/pdf_balance_parser/pkg/helper"
)

type Line3 struct {
	OtherCosts              float64
	Taxes                   float64
	PositionAdjustment      float64
	DayTradeAdjustment      float64
	DayTradeAdjustmentOp    string
	TotalOperationalCosts   float64
	TotalOperationalCostsOp string
}

func (l *Line3) ParseLines(line *string) {
	var columns []string
	columns = helper.Explode(" ", *line)
	l.OtherCosts, _ = strconv.ParseFloat(columns[0], 64)
	l.Taxes, _ = strconv.ParseFloat(columns[1], 64)
	l.PositionAdjustment, _ = strconv.ParseFloat(columns[2], 64)
	l.DayTradeAdjustment, _ = strconv.ParseFloat(columns[3], 64)
	l.DayTradeAdjustmentOp = columns[4]
	l.TotalOperationalCosts, _ = strconv.ParseFloat(columns[5], 64)
	l.TotalOperationalCostsOp = columns[6]
}
