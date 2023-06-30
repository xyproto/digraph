package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"unicode"

	"github.com/xyproto/digraph"
)

func printableSymbol(symbol rune) string {
	if unicode.IsPrint(symbol) {
		return fmt.Sprintf("%c", symbol)
	}
	return "N/A"
}

func main() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer w.Flush()

	for _, twoLetters := range digraph.All() {
		symbol := digraph.MustLookup(twoLetters)
		description := digraph.MustLookupDescription(twoLetters)

		symbolStr := printableSymbol(symbol)

		// print hex code, two letters, description and symbol
		fmt.Fprintf(w, "%04X\t%s\t%s\t%s\n", symbol, twoLetters, description, symbolStr)
	}
}
