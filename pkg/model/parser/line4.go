package parser

import (
	"strconv"

	"github.com/rcarreirao/pdf_balance_parser/pkg/helper"
)

type Line4 struct {
	Others                   float64
	IrrfOperational          float64
	TotalInvestmentAccount   float64
	TotalInvestmentAccountOp string
	TotalNormalAccount       float64
	TotalNet                 float64
	TotalNetOp               string
	TotalNetInvoice          float64
	TotalNetInvoiceOp        string
}

func (l *Line4) ParseLines(line *string) {
	var columns []string
	columns = helper.Explode(" ", *line)
	l.Others, _ = strconv.ParseFloat(columns[0], 64)
	l.IrrfOperational, _ = strconv.ParseFloat(columns[1], 64)
	l.TotalInvestmentAccount, _ = strconv.ParseFloat(columns[2], 64)
	l.TotalNormalAccount, _ = strconv.ParseFloat(columns[3], 64)
	l.TotalInvestmentAccountOp = columns[4]
	l.TotalNet, _ = strconv.ParseFloat(columns[5], 64)
	l.TotalNetOp = columns[6]
	l.TotalNetInvoice, _ = strconv.ParseFloat(columns[7], 64)
	l.TotalNetInvoiceOp = columns[8]
}
