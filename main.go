package main

//go:generate go run main.go extract . ./locales/en.json
//go:generate go run main.go update ./locales/en.json ./locales/zh-Hans.json
//go:generate go run main.go update ./locales/en.json ./locales/zh-Hant.json
//go:generate go run main.go generate --pkg=catalog ./locales ./catalog/catalog.go

import (
	"fmt"
	"log"
	"os"

	"github.com/Rain-31/go-i18n/v1/i18n"
	_ "github.com/Rain-31/i18n-gen/catalog"
	tools "github.com/Rain-31/i18n-gen/tools"
	"github.com/Xuanwo/go-locale"
	uuid "github.com/satori/go.uuid"
	"github.com/urfave/cli/v2"
)

func main() {
	// Detect OS language
	tag, _ := locale.Detect()
	var sessionId string
	if uuid, err := uuid.NewV4(); err != nil {
		log.Fatal(err)
	} else {
		sessionId = uuid.String()
	}

	// Set Language
	i18n.RegistPrinter(sessionId, tag)
	defer i18n.DeletePrinter(sessionId)

	appName := "i18n-gen"

	app := &cli.App{
		HelpName: appName,
		Name:     appName,
		Usage:    i18n.Sprintf(sessionId, `a tool for managing message translations.`),
		Action: func(c *cli.Context) error {
			cli.ShowAppHelp(c)
			return nil
		},

		Commands: []*cli.Command{
			{
				Name:      "update",
				Aliases:   []string{"u"},
				Usage:     i18n.Sprintf(sessionId, `merge translations and generate catalog`),
				UsageText: i18n.Sprintf(sessionId, `%s update srcfile destfile`, appName),
				Action: func(c *cli.Context) error {
					srcFile := c.Args().Get(0)
					if len(srcFile) == 0 {
						return fmt.Errorf(i18n.Sprintf(sessionId, `srcfile cannot be empty`))
					}

					destFile := c.Args().Get(1)
					if len(destFile) == 0 {
						return fmt.Errorf(i18n.Sprintf(sessionId, `destfile cannot be empty`))
					}

					err := tools.Update(sessionId, srcFile, destFile)

					return err
				},
			},
			{
				Name:      "extract",
				Aliases:   []string{"e"},
				Usage:     i18n.Sprintf(sessionId, `extracts strings to be translated from code`),
				UsageText: i18n.Sprintf(sessionId, `%s extract [path] [outfile]`, appName),
				Action: func(c *cli.Context) error {
					path := c.Args().Get(0)
					if len(path) == 0 {
						path = "."
					}
					outFile := c.Args().Get(1)
					if len(outFile) == 0 {
						outFile = "./locales/en.json"
					}
					err := tools.Extract([]string{
						path,
					}, outFile)
					return err
				},
			},
			{
				Name:      "generate",
				Aliases:   []string{"g"},
				Usage:     i18n.Sprintf(sessionId, `generates code to insert translated messages`),
				UsageText: i18n.Sprintf(sessionId, `%s generate [path] [outfile]`, appName),
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "pkg",
						Value: "main",
						Usage: i18n.Sprintf(sessionId, `generated go file package name`),
					},
				},
				Action: func(c *cli.Context) error {
					path := c.Args().Get(0)
					if len(path) == 0 {
						path = "./locales"
					}
					outFile := c.Args().Get(1)
					if len(outFile) == 0 {
						outFile = "./catalog.go"
					}
					pkgName := c.String("pkg")
					err := tools.Generate(
						sessionId,
						pkgName,
						[]string{
							path,
						}, outFile)
					return err
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
