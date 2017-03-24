package filters

import "testing"
import "github.com/stretchr/testify/assert"
import "github.com/appPlant/alpinepass/src/data"

func TestFilterLocationTrue(t *testing.T) {
	config := data.Config{Location: "Earth"}
	filter := "location=Earth"
	assert.Equal(t, true, filterProperty(filter, config).IsValid,
		"The Config should match the filter.")
}

func TestFilterLocationFalse(t *testing.T) {
	config := data.Config{Location: "Moon"}
	filter := "location=Earth"
	assert.Equal(t, false, filterProperty(filter, config).IsValid,
		"The Config should not match the filter.")
}
