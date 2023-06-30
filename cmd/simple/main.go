package main

import (
	"fmt"

	"github.com/xyproto/digraph"
)

func main() {
	fmt.Printf("The symbol for My is %c\n", digraph.MustLookup("My"))
}
