# i18n-gen
a go-i18n command line tools

this project was fork from https://github.com/mylukin/easy-i18n/commit/b39b840e1e769a3a80f7e6acabc456a09d0b1ad5.

---

i18n-gen is a Go package and a command that helps you translate Go programs into multiple languages.

* Supports pluralized strings with =x or >x expression.
* Supports strings with similar to fmt.Sprintf format syntax.
* Supports session with i18n printer.
* Supports message files of any format (e.g. JSON, TOML, YAML).

# Package go-i18n

The go-i18n package provides support for looking up messages according to a set of locale preferences.

```go
package main

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

	printer := i18n.RegistPrinter(sessionId, language.SimplifiedChinese)
	defer i18n.DeletePrinter(sessionId)

	i18n.Session(sessionId).Printf(`hello world!`)
	fmt.Println()

	i18n.Printf(sessionId, `no session example`)
	fmt.Println()

	printer.Printf("printer example")
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

```

# Command i18n-gen

The i18n-gen command manages message files used by the i18n package.

```
go get -u github.com/Rain-31/i18n-gen
i18n-gen -h

  update, u    merge translations and generate catalog
  extract, e   extracts strings to be translated from code
  generate, g  generates code to insert translated messages
```

### Extracting messages

Use `i18n-gen extract . ./locales/en.json` to extract all i18n.Sprintf function literals in Go source files to a message file for translation.

`./locales/en.json`
```json
{
  "%s has %d cat.": "%s has %d cat.",
  "%s has %d cats.": "%s has %d cats.",
  "%s have %d apples.": "%s have %d apples.",
  "%s have an apple.": "%s have an apple.",
  "%s have two apples.": "%s have two apples.",
  "hello %s!": "hello %s!",
  "hello world!": "hello world!"
}
```

### Translating a new language

1. Create an empty message file for the language that you want to add (e.g. `zh-Hans.json`).
2. Run `i18n-gen update ./locales/en.json ./locales/zh-Hans.json` to populate `zh-Hans.json` with the mesages to be translated.

	`./locales/zh-Hans.json`
	```json
	{
	  "%s has %d cat.": "%s有%d只猫。",
	  "%s has %d cats.": "%s有%d只猫。",
	  "%s have %d apples.": "%s有%d个苹果。",
	  "%s have an apple.": "%s有一个苹果。",
	  "%s have two apples.": "%s有两个苹果。",
	  "hello %s!": "你好%s！",
	  "hello world!": "你好世界！"
	}
	```
3. After `zh-Hans.json` has been translated, run `i18n-gen generate --pkg=catalog ./locales ./catalog/catalog.go`.

4. Import `catalog` package in main.go, example: `import _ "github.com/Rain-31/i18n-gen/catalog"` 

### Translating new messages

If you have added new messages to your program:

1. Run `i18n-gen extract` to update `./locales/en.json` with the new messages.
2. Run `i18n-gen update ./locales/en.json` to generate updated `./locales/new-language.json` files.
3. Translate all the messages in the `./locales/new-language.json` files.
4. Run `i18n-gen generate --pkg=catalog ./locales ./catalog/catalog.go` to merge the translated messages into the go files.

## License

i18n-gen is available under the MIT license. See the [LICENSE](LICENSE) file for more info.