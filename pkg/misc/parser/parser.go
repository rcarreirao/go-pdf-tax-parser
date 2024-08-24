package parser

import (
	"log"
	"os"
	"pdf_balance_parser/pkg/misc/pdf"
	"pdf_balance_parser/pkg/model/document"
	"strings"
)

var regexes map[int]map[int]string
var regexesKeys []int

func ParseFile(file string) document.Document {
	content, err := pdf.ReadPdf(file)
	if err != nil {
		panic(err)
	}
	document := document.Document{
		Content: content,
	}
	return document
}

func ParseDirectory(directory string) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		ParseDocument(ParseFile(e.Name()))
	}
}

func ParseDocument(d document.Document) {
	auctionDay := ParseDocumentAuction(d)
	ParseDocumentTradingNote(auctionDay, d)
}

func replaceBoundariesRowValues(regex string, line *string) {
	regex = strings.Replace(regex, "\\)", ")", 10)
	*line = strings.Replace(strings.Replace(*line, " |", "", 3), regex, "", 1)
	*line = strings.Replace(*line, "  ", "", 3)
	*line = strings.Replace(*line, "|", "", 3)
	*line = strings.Trim(*line, " ")
}
