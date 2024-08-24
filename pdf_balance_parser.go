package pdf_balance_parser

import (
	"fmt"
	"log"
	"os"

	"github.com/rcarreirao/pdf_balance_parser/database/migrate"
	"github.com/rcarreirao/pdf_balance_parser/pkg/misc/parser"

	"github.com/joho/godotenv"
	pdf_eof_fix "github.com/rcarreirao/golang_pdf_eof_fix"
	"github.com/urfave/cli/v2"
)

func pdf_balance_parser() {
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
				Name:    "parse-directory",
				Aliases: []string{"pd"},
				Usage:   "parse multiple files from directory",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("Initializing parser directory")
					if cCtx.Args().Len() < 1 {
						return cli.Exit("File parameter not sent", 86)
					}
					parser.ParseDirectory(cCtx.Args().Get(0))
					fmt.Println("Parser directory finished")
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
					pdf_eof_fix.FixEofOnPdfFile(cCtx.Args().Get(0))

					document := parser.ParseFile(cCtx.Args().Get(0))
					parser.ParseDocument(document)
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

}
