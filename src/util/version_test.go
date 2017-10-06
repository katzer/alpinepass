package util

import "testing"
import "github.com/stretchr/testify/assert"

func TestVersionIsCorrect(t *testing.T) {
	assert.Equal(t, "1.4.3", Version,
		"The version is not correct. Did you forget it to adjust somewhere?")
}
