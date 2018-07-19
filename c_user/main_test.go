package main

import (
	"testing"

	"github.com/mwf/vgo-modules/c"
)

func TestC(t *testing.T) {
	expected := "CA"
	if c.CA != expected {
		t.Fatalf("c.CA: got %s, expected: %s", c.CA, expected)
	}
}
