package digraphs

import (
	"testing"
)

func TestMustLookup(t *testing.T) {
	if MustLookup("My") != 'µ' {
		t.Fail()
	}
}

func TestLookup(t *testing.T) {
	if r, ok := Lookup("My"); !ok || r != 'µ' {
		t.Fail()
	}
}

func TestMustLookupDescription(t *testing.T) {
	if desc := MustLookupDescription("My"); desc != "MICRO SIGN" {
		t.Fail()
	}
}

func TestLookupDescription(t *testing.T) {
	if desc, ok := LookupDescription("My"); !ok || desc != "MICRO SIGN" {
		t.Fail()
	}
}
