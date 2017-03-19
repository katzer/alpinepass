package util

import "testing"

func TestCleanStringFiller(t *testing.T) {
	if Filler != "_" {
		t.Errorf("The Filler string might not match the other tests in this file.")
	}
}

func TestCleanString(t *testing.T) {
	expected := "This_is_some_input."
	actual := CleanString("This is some input.")
	if expected != actual {
		t.Errorf("The cleaned strings do not match.")
	}
}
