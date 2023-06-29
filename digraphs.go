// Package digraphs provides functions for looking up ViM-style digraphs
package digraphs

import (
	_ "embed"
	"strconv"
	"strings"
)

var (
	digraphMap     map[string]rune
	descriptionMap map[string]string
)

//go:embed digraphs.txt
var digraphs string

func init() {
	// Used to avoid memory allocations while filling the maps
	const digraphsInDigraphFile = 1305

	digraphMap = make(map[string]rune, digraphsInDigraphFile)
	descriptionMap = make(map[string]string, digraphsInDigraphFile)

	lines := strings.Split(digraphs, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}
		fields := strings.Fields(line)
		digraph := fields[0]
		//hexvalue := fields[1]
		decvalue := fields[2]

		num, err := strconv.Atoi(decvalue)
		if err != nil {
			continue
		}

		digraphMap[digraph] = rune(num)

		description := strings.Join(fields[3:], " ")
		descriptionMap[digraph] = description
	}
}

// MustLookup tries to look up a digraph and return a rune,
// but does not return any bool or error if the digraph string is not found.
func MustLookup(digraph string) rune {
	return digraphMap[digraph]
}

// Lookup a digraph and return a rune. Returns false if the digraph could not be found.
func Lookup(digraph string) (rune, bool) {
	r, ok := digraphMap[digraph]
	return r, ok
}

// MustLookupDescription tries to look up a digraph and return a description,
// but it edoes not return any bool or error if the digraph string is not found.
func MustLookupDescription(digraph string) string {
	return descriptionMap[digraph]
}

// LookupDescription tries to look up a digraph and return a description.
// Returns false if the digraph could not be found.
func LookupDescription(digraph string) (string, bool) {
	description, ok := descriptionMap[digraph]
	return description, ok
}
