package parser

import (
	"pdf_balance_parser/pkg/helper"
	"strconv"
)

type Line2 struct {
	Irrf           float64
	IrrfDayTrade   float64
	OperationalTax float64
	TaxRegisterBmf float64
	TaxBmf         float64
	TaxBmfOp       string
}

func (l *Line2) ParseLines(line *string) {
	var columns []string
	columns = helper.Explode(" ", *line)
	l.Irrf, _ = strconv.ParseFloat(columns[0], 64)
	l.IrrfDayTrade, _ = strconv.ParseFloat(columns[1], 64)
	l.OperationalTax, _ = strconv.ParseFloat(columns[2], 64)
	l.TaxRegisterBmf, _ = strconv.ParseFloat(columns[3], 64)
	l.TaxBmf, _ = strconv.ParseFloat(columns[4], 64)
	l.TaxBmfOp = columns[5]
}
