package util

import "testing"
import "github.com/stretchr/testify/assert"

func TestCleanStringFiller(t *testing.T) {
	assert.Equal(t, "_", Filler, "The Filler string might not match the other tests in this file.")
}

func TestCleanString(t *testing.T) {
	expected := "This_is_some_input."
	actual := CleanString("This is some input.")
	assert.Equal(t, expected, actual, "The cleaned strings should match.")
}
