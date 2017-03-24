package filters

import "testing"
import "github.com/stretchr/testify/assert"
import "github.com/appPlant/alpinepass/src/data"

func TestFilterLocationTrue(t *testing.T) {
	filter := "location=Earth"
	config := data.Config{Location: "Earth"}
	filteredConfig := filterProperty(filter, config)

	assert.Equal(t, true, filteredConfig.IsValid,
		"The Config should match the filter.")
}

func TestFilterLocationFalse(t *testing.T) {
	filter := "location=Earth"
	config := data.Config{Location: "Moon"}
	filteredConfig := filterProperty(filter, config)

	assert.Equal(t, false, filteredConfig.IsValid,
		"The Config should not match the filter.")
}
