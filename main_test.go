package main

import (
	"testing"
)

func Test_pekofyTextNoEndPunct(t *testing.T) {
	example := "Hello there"

	want := "Hello there peko."

	err := pekofyText(&example)

	if !(example == want) || (err != nil) {
		t.Fatalf(`Expected: %q but got: %q`, want, example)
	}
}

func Test_pekofyTextNormal(t *testing.T) {
	example := "Gura is very cute."

	want := "Gura is very cute peko."

	err := pekofyText(&example)

	if !(example == want) || (err != nil) {
		t.Fatalf(`ExpectedL %q but got: %q`, want, example)
	}
}

func Test_pekofyTextWithCos(t *testing.T) {
	example := "I like going to concerts"

	want := "I like going to pekoncerts peko."

	err := pekofyText(&example)

	if !(example == want) || (err != nil) {
		t.Fatalf(`ExpectedL %q but got: %q`, want, example)
	}
}
