package main

import (
	"fmt"

	"github.com/xyproto/ask"
	"github.com/xyproto/digraph"
)

func main() {
	twoLetters := ask.Ask("Lookup a digraph: ")
	symbol, found := digraph.Lookup(twoLetters)
	if !found {
		fmt.Printf("Could not find a digraph for %s\n", twoLetters)
		return
	}
	description := digraph.MustLookupDescription(twoLetters)
	fmt.Printf("Found %c [%v]: %s\n", symbol, symbol, description)
}
