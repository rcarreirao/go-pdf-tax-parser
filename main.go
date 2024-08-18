package main

import (
	"fmt"
	"log"
	"os"
	"pdf_balance_parser/database/migrate"
	"pdf_balance_parser/misc/parser"
	"pdf_balance_parser/misc/pdf"
	"pdf_balance_parser/pkg/model/document"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := &cli.App{
		Name:  "Pdf Tax Parser",
		Usage: "An application to calculate monthly taxes for day trade users",
		Commands: []*cli.Command{
			{
				Name:    "migrate",
				Aliases: []string{"m"},
				Usage:   "migrate models to database",
				Action: func(cCtx *cli.Context) error {
					migrate.MigrateExec()
					return nil
				},
			},
			{
				Name:    "parse",
				Aliases: []string{"p"},
				Usage:   "parse a file",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Initializing")
					/**
					Trying first to open pdf file
					*/
					if cCtx.Args().Len() < 1 {
						return cli.Exit("File parameter not sent", 86)
					}
					content, err := pdf.ReadPdf(cCtx.Args().Get(0))
					document := document.Document{
						Content: content,
					}
					parser.ParseDocument(document)
					if err != nil {
						panic(err)
					}
					return nil
				},
			},
		},
		Action: func(cCtx *cli.Context) error {
			// default path without parameter
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	return

}
