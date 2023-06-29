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
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		fields := strings.Fields(line)
		digraph := fields[0]
		//hexvalue := fields[1]
		decvalue := fields[2]
		description := strings.Join(fields[3:], " ")

		num, err := strconv.Atoi(decvalue)
		if err != nil {
			continue
		}

		digraphMap[digraph] = rune(num)
		descriptionMap[digraph] = description
	}
}

func Lookup(digraph string) (rune, bool) {
	r, ok := digraphMap[digraph]
	return r, ok
}

func LookupDescription(digraph string) (string, bool) {
	description, ok := descriptionMap[digraph]
	return description, ok
}

func MustLookup(digraph string) rune {
	return digraphMap[digraph]
}

func MustLookupDescription(digraph string) string {
	return descriptionMap[digraph]
}
