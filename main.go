package main

import (
	"fmt"
	"os"
	"pdf_balance_parser/misc/parser"
	"pdf_balance_parser/misc/pdf"
	"pdf_balance_parser/pkg/model/document"
)

func main() {
	/**
	Trying first to open pdf file
	*/
	if len(os.Args) <= 1 {
		fmt.Println("File parameter not sent")
		os.Exit(1)
	}
	parser.LoadRegexMap()
	content, err := pdf.ReadPdf(os.Args[1])
	document := document.Document{
		Content: content,
	}
	parser.ParseDocument(document)
	if err != nil {
		panic(err)
	}
	return

}
