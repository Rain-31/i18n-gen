package main

//go:generate i18n-gen extract . ./locales/en.json
//go:generate i18n-gen update ./locales/en.json ./locales/zh-Hans.json
//go:generate i18n-gen generate --pkg=catalog ./locales ./catalog/catalog.go
//go:generate go build -o example

import (
	"fmt"
	"log"
	"os"

	"github.com/Rain-31/go-i18n/v1/i18n"
	_ "github.com/Rain-31/i18n-gen/example/catalog"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/text/language"
)

func main() {
	var sessionId string
	if uuid, err := uuid.NewV4(); err != nil {
		log.Fatal(err)
	} else {
		sessionId = uuid.String()
	}

	i18n.RegistPrinter(sessionId, language.SimplifiedChinese)
	defer i18n.DeletePrinter(sessionId)

	i18n.Session(sessionId).Printf(`hello world!`)
	fmt.Println()

	name := `Lukin`

	i18n.Session(sessionId).Printf(`hello %s!`, name)
	fmt.Println()

	i18n.Session(sessionId).Printf(`%s has %d cat.`, name, 1)
	fmt.Println()

	i18n.Session(sessionId).Printf(`%s has %d cat.`, name, 2, i18n.Plural(
		`%[2]d=1`, `%s has %d cat.`,
		`%[2]d>1`, `%s has %d cats.`,
	))
	fmt.Println()

	i18n.Session(sessionId).Fprintf(os.Stderr, `%s have %d apple.`, name, 2, i18n.Plural(
		`%[2]d=1`, `%s have an apple.`,
		`%[2]d=2`, `%s have two apples.`,
		`%[2]d>2`, `%s have %d apples.`,
	))
	fmt.Println()
}
