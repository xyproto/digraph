package digraphs

import (
	"fmt"
	"testing"
)

func TestLookup(t *testing.T) {
	if MustLookup("My") != 'µ' {
		t.Fail()
	}
}

func TestLookupDescription(t *testing.T) {
	if desc := MustLookupDescription("My"); desc != "MICRO SIGN" {
		t.Fail()
	}
}
