package main

import (
	"testing"
)

func Test_pekofyTextNoEndPunct(t *testing.T) {
	example := "Hello there"

	want := "Hello there peko."

	res, err := pekofyText(example)

	if !(res == want) || (err != nil) {
		t.Fatalf(`Expected: %q but got: %q`, want, res)
	}
}

func Test_pekofyTextNormal(t *testing.T) {
	example := "Gura is very cute."

	want := "Gura is very cute peko."

	res, err := pekofyText(example)

	if !(res == want) || (err != nil) {
		t.Fatalf(`ExpectedL %q but got: %q`, want, res)
	}
}
