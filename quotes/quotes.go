package quotes

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"strings"
)

var (
	Quotes      []string
	goPath      = os.Getenv("GOPATH")
	projectRoot string
	quoteFile   string
)

func init() {
	rand.Seed(64)

	if goPath == "" {
		projectRoot = "."
	} else {
		projectRoot = path.Join(goPath, "src", "github.com", "darthlukan", "bct", "quotes")
	}

	quoteFile = fmt.Sprintf("%s/quotes.txt", projectRoot)
	content, err := ioutil.ReadFile(quoteFile)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	Quotes = strings.Split(string(content), "\n\n")
	if len(Quotes) > 0 {
		fmt.Printf("Quotes populated!\n")
	}
}

func RandomQuote() string {
	return Quotes[rand.Intn(len(Quotes))]
}
